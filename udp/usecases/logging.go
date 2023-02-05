package usecases

import (
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
)

func KeyExchange(p template.BasePacket, addr *net.UDPAddr) (*template.BasePacket, error) {
	db := gsetup.DB

	lockID := strings.ToUpper(hex.EncodeToString(p.Data[:16]))
	pLen := len(p.Data)

	if pLen < 16 {
		return nil, errors.New("packet too short")
	}

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

	lock := models.Lock{
		LockID:    lockID,
		IpAddress: addr.IP.String(),
		Label:     fmt.Sprintf("Lock on %s", addr.IP.String()),
		PublicKey: hex.EncodeToString(p.Data[16:pLen]),
	}

	if err := db.Create(&lock).Error; err != nil {
		return nil, err
	}

	return utils.MakePacket(
		template.KeyExchange,
		append(p.Data[:16], serverPubKeyBytes...),
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
		return nil, errors.New("invalid packet signature")
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
		KeyID:            key.ID,
		PersonelName:     personel.Name,
		PersonelIDNumber: personel.IDNumber,
		Timestamp:        time.Now(),
		Location:         lock.Location,
	}

	if err := db.Create(&accessLog).Error; err != nil {
		return nil, err
	}

	return utils.MakePacket(template.LogAccessEvent, p.Data, setup.PrivateKey)
}

func LogRSSIEvent(p template.BasePacket) (*template.BasePacket, error) {
	db := gsetup.DB

	pLen := len(p.Data)
	if pLen < 33 {
		return nil, errors.New("packet too short")
	}

	lockID := strings.ToUpper(hex.EncodeToString(p.Data[:16]))
	keyID := strings.ToUpper(hex.EncodeToString(p.Data[16:32]))
	rssi := int(p.Data[32])

	var lock models.Lock
	if err := db.First(&lock, "lock_id = ?", lockID).Error; err != nil {
		return nil, errors.New("lock not found")
	}

	pubKeyBytes, _ := hex.DecodeString(lock.PublicKey)
	pubKey, _ := utils.ParseECDSAPublickKey(pubKeyBytes)
	if err := utils.VerifyPacket(p.Bytes(), pubKey); err != nil {
		return nil, err
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
		RSSI:       rssi,
		PersonelID: personel.ID,
		LockID:     lock.ID,
		KeyID:      key.ID,
		Timestamp:  time.Now(),
	}

	if err := db.Create(&rssiLog).Error; err != nil {
		return nil, err
	}

	return utils.MakePacket(template.LogRSSIEvent, p.Data, setup.PrivateKey)
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

	lockID, _ := hex.DecodeString(lock.LockID)
	data = append(lockID, temp...)

	packet, err := utils.MakePacket(template.LogHealthcheckEvent, data, setup.PrivateKey)
	if err != nil {
		return nil, err
	}

	res, err := utils.SendUDPPacket(packet, lock.IpAddress)
	if err != nil {
		return nil, err
	}

	return res, utils.VerifyPacket(res.Bytes(), pubKey)
}
