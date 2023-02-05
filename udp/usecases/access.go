package usecases

import (
	"encoding/binary"
	"encoding/hex"
	"server/models"
	"server/udp/setup"
	"server/udp/template"
	"server/udp/utils"
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

	return utils.SendUDPPacket(packet, accessRule.Lock.IpAddress)
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

	return utils.SendUDPPacket(packet, accessRule.Lock.IpAddress)
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
