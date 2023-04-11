package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"server/udp/template"
)

const SIGNATURE_LEN = 64

func PadSignature(sig []byte) []byte {
	t := []byte{}
	for i := len(sig); i < SIGNATURE_LEN; i++ {
		t = append(t, 0x00)
	}
	return append(t, sig...)
}

func TrimSignature(sig []byte) []byte {
	for i := 0; i < len(sig); i++ {
		if sig[i] != 0x00 {
			return sig[i:]
		}
	}
	return sig
}

func MakePacket(opCode byte, data []byte, privKey *ecdsa.PrivateKey) (*template.BasePacket, error) {
	packet := template.BasePacket{
		OpCode: opCode,
		Data:   data,
	}

	temp := []byte{}
	temp = append(temp, opCode)
	temp = append(temp, data...)

	hash := sha256.Sum256(temp)
	signature, err := ecdsa.SignASN1(rand.Reader, privKey, hash[:])
	if err != nil {
		return nil, err
	}

	packet.Signature = PadSignature(signature)

	return &packet, nil
}

func ParsePacket(pBytes []byte, pLen int) (*template.BasePacket, error) {
	if pLen < 65 {
		return nil, errors.New("packet too short")
	}

	pBytes = pBytes[:pLen]

	return &template.BasePacket{
		OpCode:    pBytes[0],
		Data:      pBytes[1 : pLen-SIGNATURE_LEN],
		Signature: pBytes[pLen-SIGNATURE_LEN:],
	}, nil
}

func VerifyPacket(packet []byte, pk *ecdsa.PublicKey) error {
	packetLen := len(packet)

	if packetLen < SIGNATURE_LEN {
		return errors.New("packet too short")
	}

	signature := packet[packetLen-SIGNATURE_LEN:]
	temp := packet[:packetLen-SIGNATURE_LEN]

	hash := sha256.Sum256(temp)
	if ecdsa.VerifyASN1(pk, hash[:], TrimSignature(signature)) {
		return nil
	}

	return errors.New("invalid signature")
}
