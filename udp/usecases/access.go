package usecases

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
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

func AddAccessRule(accessRule models.AccessRule, lock models.Lock, key models.Key) (*template.BasePacket, error) {
	var data []byte = make([]byte, 8)

	lockIDHex, err := hex.DecodeString(lock.LockID)
	keyIDHex, err := hex.DecodeString(key.KeyID)

	binary.BigEndian.PutUint64(data, accessRule.ID)
	data = append(data, lockIDHex...)
	data = append(data, keyIDHex...)
	data = binary.BigEndian.AppendUint64(data, uint64(accessRule.StartsAt.Unix()))
	data = binary.BigEndian.AppendUint64(data, uint64(accessRule.EndsAt.Unix()))

	packet, err := utils.MakePacket(template.AddAccessRule, data, setup.PrivateKey)
	if err != nil {
		return nil, err
	}

	return utils.SendUDPPacket(packet, lock.IpAddress)
}

func EditAccessRule(accessRule models.AccessRule, lock models.Lock, key models.Key) (*template.BasePacket, error) {
	var data []byte = make([]byte, 8)

	lockIDHex, err := hex.DecodeString(lock.LockID)
	keyIDHex, err := hex.DecodeString(key.KeyID)

	binary.BigEndian.PutUint64(data, accessRule.ID)
	data = append(data, lockIDHex...)
	data = append(data, keyIDHex...)
	data = binary.BigEndian.AppendUint64(data, uint64(accessRule.StartsAt.Unix()))
	data = binary.BigEndian.AppendUint64(data, uint64(accessRule.EndsAt.Unix()))

	packet, err := utils.MakePacket(template.EditAccessRule, data, setup.PrivateKey)
	if err != nil {
		return nil, err
	}

	return utils.SendUDPPacket(packet, lock.IpAddress)
}

func DeleteAccessRule(accessRuleID uint64, ipAddress string) (*template.BasePacket, error) {
	var data []byte = make([]byte, 8)

	binary.BigEndian.PutUint64(data, accessRuleID)

	packet, err := utils.MakePacket(template.DeleteAccessRule, data, setup.PrivateKey)
	if err != nil {
		return nil, err
	}

	return utils.SendUDPPacket(packet, ipAddress)
}

func SyncAccessRules(p template.BasePacket, addr *net.UDPAddr) (*template.BasePacket, error) {
	db := gsetup.DB

	lockID := strings.ToUpper(hex.EncodeToString(p.Data[:16]))

	var lock models.Lock
	if err := db.First(&lock, "lock_id = ?", lockID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no such lock")
		}

		return nil, err
	}

	var i byte
	lockAccessRuleIds := make([]uint64, 0)
	lockAccessRuleMap := make(map[uint64]bool, p.Data[16])
	for i = 0; i < p.Data[16]; i++ {
		id := binary.BigEndian.Uint64(p.Data[i*8+17 : i*8+25])
		lockAccessRuleIds = append(lockAccessRuleIds, id)
		lockAccessRuleMap[id] = false
	}

	serverAccessRules := []models.AccessRule{}
	serverAccessRulesMap := make(map[uint64]bool, 0)
	db.Find(&serverAccessRules).Where("lock_id = ? AND ends_at <= ?", lockID, time.Now())
	for i := 0; i < (len(serverAccessRules)); i++ {
		serverAccessRulesMap[serverAccessRules[i].ID] = false
	}

	for i = 0; i < p.Data[16]; i++ {
		if _, exists := serverAccessRulesMap[lockAccessRuleIds[i]]; !exists {
			DeleteAccessRule(lockAccessRuleIds[i], lock.IpAddress)
		}
	}

	for i := 0; i < len(serverAccessRules); i++ {
		if _, exists := lockAccessRuleMap[serverAccessRules[i].ID]; !exists {
			var key models.Key
			db.First(&key, serverAccessRules[i].KeyID)
			AddAccessRule(serverAccessRules[i], lock, key)
		}
	}

	var responseData []byte
	responseData = append(responseData, byte(len(serverAccessRules)))
	for _, accessRule := range serverAccessRules {
		responseData = binary.BigEndian.AppendUint64(responseData, accessRule.ID)
	}

	return utils.MakePacket(
		template.KeyExchange,
		append([]byte{byte(len(serverAccessRules))}, responseData...),
		setup.PrivateKey,
	)
}
