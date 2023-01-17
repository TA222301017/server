package usecases

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"server/udp/setup"
	"server/udp/template"
	"server/udp/utils"
	"strings"
)

func KeyExchange(p template.BasePacket) (*template.BasePacket, error) {
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

	return utils.MakePacket(
		template.KeyExchange,
		append(p.Data[:16], serverPubKeyBytes...),
	)
}
