package usecases

import (
	"encoding/binary"
	"server/models"
	"server/udp/template"
	"server/udp/utils"
)

func AddAccessRule(accessRule models.AccessRule) (*template.BasePacket, error) {
	var data []byte = make([]byte, 8)

	binary.BigEndian.PutUint64(data, accessRule.ID)
	data = binary.BigEndian.AppendUint64(data, accessRule.LockID)
	data = binary.BigEndian.AppendUint64(data, accessRule.KeyID)
	data = binary.BigEndian.AppendUint64(data, uint64(accessRule.StartsAt.Unix()))
	data = binary.BigEndian.AppendUint64(data, uint64(accessRule.EndsAt.Unix()))

	packet, err := utils.MakePacket(template.AddAccessRule, data)
	if err != nil {
		return nil, err
	}

	return utils.SendUDPPacket(packet, accessRule.Lock.IpAddress)
}

func EditAccessRule(accessRule models.AccessRule) (*template.BasePacket, error) {
	var data []byte = make([]byte, 8)

	binary.BigEndian.PutUint64(data, accessRule.ID)
	data = binary.BigEndian.AppendUint64(data, accessRule.LockID)
	data = binary.BigEndian.AppendUint64(data, accessRule.KeyID)
	data = binary.BigEndian.AppendUint64(data, uint64(accessRule.StartsAt.Unix()))
	data = binary.BigEndian.AppendUint64(data, uint64(accessRule.EndsAt.Unix()))

	packet, err := utils.MakePacket(template.EditAccessRule, data)
	if err != nil {
		return nil, err
	}

	return utils.SendUDPPacket(packet, accessRule.Lock.IpAddress)
}

func DeleteAccessRule(accessRuleID uint64, ipAddress string) (*template.BasePacket, error) {
	var data []byte = make([]byte, 8)

	binary.BigEndian.PutUint64(data, accessRuleID)

	packet, err := utils.MakePacket(template.EditAccessRule, data)
	if err != nil {
		return nil, err
	}

	return utils.SendUDPPacket(packet, ipAddress)
}
