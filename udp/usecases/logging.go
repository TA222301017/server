package usecases

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"server/models"
	gsetup "server/setup"
	"server/udp/setup"
	"server/udp/template"
	"server/udp/utils"
	"strings"
	"time"

	"gorm.io/gorm"
)

func KeyExchange(p template.BasePacket, addr *net.UDPAddr) (*template.BasePacket, error) {
	db := gsetup.DB

	lockID := strings.ToUpper(hex.EncodeToString(p.Data[:16]))
	pLen := len(p.Data)

	if pLen < 16 {
		return nil, errors.New("packet too short")
	}

	fmt.Println(hex.EncodeToString(p.Data))
	fmt.Println(hex.EncodeToString(p.Signature))

	pubKey, err := utils.ParseECDSAPublickKey(p.Data[16:pLen])
	if err != nil {
		return nil, err
	}

	if err := utils.VerifyPacket(p.Bytes(), pubKey); err != nil {
		return nil, err
	}

	if err := utils.SaveECDSAPublicKey(pubKey, lockID); err != nil {
		return nil, err
	}

	serverPubKeyBytes := utils.MarshalECDSAPublicKey(setup.PublicKey)

	var lock models.Lock
	if err := db.First(&lock, "lock_id = ?", lockID).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if lock.ID == 0 {
		lock = models.Lock{
			LockID:    lockID,
			IpAddress: addr.IP.String(),
			Label:     fmt.Sprintf("Lock on %s", addr.IP.String()),
			PublicKey: hex.EncodeToString(p.Data[16:pLen]),
			Status:    true,
		}

		if err := db.Create(&lock).Error; err != nil {
			return nil, err
		}
	} else {
		lock.IpAddress = addr.IP.String()
		lock.PublicKey = hex.EncodeToString(p.Data[16:pLen])
		lock.Status = true

		if err := db.Save(&lock).Error; err != nil {
			return nil, err
		}
	}

	t := time.Now().Unix()
	timestamp := make([]byte, 8)
	binary.BigEndian.PutUint64(timestamp, uint64(t))

	return utils.MakePacket(
		template.KeyExchange,
		append(timestamp, append(p.Data[:16], serverPubKeyBytes...)...),
		setup.PrivateKey,
	)
}

func LogAccessEvent(p template.BasePacket) (*template.BasePacket, error) {
	db := gsetup.DB

	pLen := len(p.Data)
	if pLen < 32 {
		return nil, errors.New("packet too short")
	}

	lockID := strings.ToUpper(hex.EncodeToString(p.Data[:16]))
	keyID := strings.ToUpper(hex.EncodeToString(p.Data[16:32]))

	var lock models.Lock
	if err := db.First(&lock, "lock_id = ?", lockID).Error; err != nil {
		return nil, err
	}

	pubKeyBytes, _ := hex.DecodeString(lock.PublicKey)
	pubKey, _ := utils.ParseECDSAPublickKey(pubKeyBytes)
	if err := utils.VerifyPacket(p.Bytes(), pubKey); err != nil {
		p, _ := utils.MakePacket(template.LogAccessEvent, p.Data, setup.PrivateKey)
		return p, errors.New("invalid packet signature")
	}

	var key models.Key
	if err := db.First(&key, "key_id = ?", keyID).Error; err != nil {
		return nil, err
	}

	var personel models.Personel
	if err := db.First(&personel, "key_id = ?", key.ID).Error; err != nil {
		return nil, err
	}

	accessLog := models.AccessLog{
		LockID:           lock.ID,
		Lock:             lock,
		KeyID:            key.ID,
		Key:              key,
		PersonelName:     personel.Name,
		PersonelIDNumber: personel.IDNumber,
		Timestamp:        time.Now(),
		Location:         lock.Location,
	}

	if err := db.Create(&accessLog).Error; err != nil {
		return nil, err
	}

	gsetup.Channel.AccessMessage <- &accessLog

	return nil, nil
}

func LogRSSIEvent(p template.BasePacket) (*template.BasePacket, error) {
	db := gsetup.DB

	pLen := len(p.Data)
	if pLen < 33 {
		return nil, errors.New("packet too short")
	}

	lockID := strings.ToUpper(hex.EncodeToString(p.Data[:16]))
	keyID := strings.ToUpper(hex.EncodeToString(p.Data[16:32]))
	rssi := int(int8(p.Data[32]))

	var lock models.Lock
	if err := db.First(&lock, "lock_id = ?", lockID).Error; err != nil {
		return nil, errors.New("lock not found")
	}

	pubKeyBytes, _ := hex.DecodeString(lock.PublicKey)
	pubKey, _ := utils.ParseECDSAPublickKey(pubKeyBytes)
	if err := utils.VerifyPacket(p.Bytes(), pubKey); err != nil {
		p, _ := utils.MakePacket(template.LogRSSIEvent, p.Data, setup.PrivateKey)
		return p, err
	}

	var key models.Key
	if err := db.First(&key, "key_id = ?", keyID).Error; err != nil {
		return nil, errors.New("key not found")
	}

	var personel models.Personel
	if err := db.First(&personel, "key_id = ?", key.ID).Error; err != nil {
		return nil, errors.New("personel not found")
	}

	rssiLog := models.RSSILog{
		RSSI:       int(rssi),
		PersonelID: personel.ID,
		Personel:   personel,
		LockID:     lock.ID,
		Lock:       lock,
		KeyID:      key.ID,
		Key:        key,
		Timestamp:  time.Now(),
	}

	if err := db.Create(&rssiLog).Error; err != nil {
		return nil, err
	}

	gsetup.Channel.RSSIMessage <- &rssiLog

	return nil, nil
}

func RequestHealthcheck(lock *models.Lock) (*template.BasePacket, error) {
	var data []byte

	temp, err := hex.DecodeString(lock.PublicKey)
	if err != nil {
		return nil, err
	}

	pubKey, err := utils.ParseECDSAPublickKey(temp)
	if err != nil {
		return nil, err
	}

	lockID, err := hex.DecodeString(lock.LockID)
	if err != nil {
		return nil, err
	}

	data = append(lockID, temp...)
	packet, err := utils.MakePacket(template.LogHealthcheckEvent, data, setup.PrivateKey)
	if err != nil {
		return nil, err
	}

	res, err := utils.SendUDPPacket(packet, lock.IpAddress)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(res.Data[:16], lockID) {
		return res, errors.New("lock_id doesn't match")
	}

	if !bytes.Equal(res.Data[16:], utils.MarshalECDSAPublicKey(pubKey)) {
		return res, errors.New("public_key doesn't match")
	}

	return res, nil
}
