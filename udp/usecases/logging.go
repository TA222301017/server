package usecases

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/binary"
	"encoding/hex"
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

	x, y := elliptic.Unmarshal(elliptic.P224(), p.Data[16:pLen])
	pubKey := ecdsa.PublicKey{
		Curve: elliptic.P224(),
		X:     x,
		Y:     y,
	}

	if err := utils.SaveECDSAPublicKey(pubKey, lockID); err != nil {
		return nil, err
	}

	serverPubKey := setup.PublicKey
	serverPubKeyBytes := elliptic.Marshal(elliptic.P224(), serverPubKey.X, serverPubKey.Y)

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
	)
}

func LogAccessEvent(p template.BasePacket) (*template.BasePacket, error) {
	db := gsetup.DB

	lockID := binary.BigEndian.Uint64(p.Data[:8])
	keyID := binary.BigEndian.Uint64(p.Data[8:16])

	var personel models.Personel
	if err := db.First(&personel, "key_id = ?", keyID).Error; err != nil {
		return nil, err
	}

	accessLog := models.AccessLog{
		LockID:           lockID,
		KeyID:            keyID,
		PersonelName:     personel.Name,
		PersonelIDNumber: personel.IDNumber,
		Timestamp:        time.Now(),
	}

	if err := db.Create(&accessLog).Error; err != nil {
		return nil, err
	}

	return utils.MakePacket(template.LogAccessEvent, p.Data)
}

func LogRSSIEvent(p template.BasePacket) (*template.BasePacket, error) {
	db := gsetup.DB

	lockID := binary.BigEndian.Uint64(p.Data[:8])
	keyID := binary.BigEndian.Uint64(p.Data[8:16])
	rssi := int(p.Data[16])

	rssiLog := models.RSSILog{
		RSSI:      rssi,
		LockID:    lockID,
		KeyID:     keyID,
		Timestamp: time.Now(),
	}

	if err := db.Create(&rssiLog).Error; err != nil {
		return nil, err
	}

	return utils.MakePacket(template.LogRSSIEvent, p.Data)
}

func RequestHealthcheck(lock *models.Lock) (*template.BasePacket, error) {
	var data []byte = make([]byte, 8)

	pubKey, err := hex.DecodeString(lock.PublicKey)
	if err != nil {
		return nil, err
	}

	binary.BigEndian.PutUint64(data, lock.ID)
	data = append(data, pubKey...)

	packet, err := utils.MakePacket(template.LogHealthcheckEvent, data)
	if err != nil {
		return nil, err
	}

	res, err := utils.SendUDPPacket(packet, lock.IpAddress)
	if err != nil {
		return nil, err
	}

	return res, utils.VerifyPacket(res.Bytes(), setup.PublicKey)
}
