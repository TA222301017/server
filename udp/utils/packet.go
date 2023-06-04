package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"server/udp/template"
)

const SIGNATURE_LEN = 72

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
	if pLen < SIGNATURE_LEN {
		return nil, errors.New("packet too short")
	}

	pBytes = pBytes[:pLen]

	return &template.BasePacket{
		OpCode:    pBytes[0],
		Data:      pBytes[1 : pLen-SIGNATURE_LEN],
		Signature: pBytes[pLen-SIGNATURE_LEN:],
	}, nil
}

func reverseBytes(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func VerifyPacket(packet []byte, pk *ecdsa.PublicKey) error {
	packetLen := len(packet)

	if packetLen < SIGNATURE_LEN {
		return errors.New("packet too short")
	}

	if pk == nil {
		return errors.New("invalid public key")
	}

	signature := packet[packetLen-SIGNATURE_LEN:]
	temp := packet[:packetLen-SIGNATURE_LEN]

	h := sha256.Sum256(temp)
	hash := h[:]

	if ecdsa.VerifyASN1(pk, hash, TrimSignature(signature)) {
		return nil
	}

	return errors.New("invalid signature")
}

func TrimLeftID(ID []byte) []byte {
	for i := 0; i < len(ID); i++ {
		if ID[i] != 0x00 {
			return ID[i:]
		}
	}

	return ID
}

func PadLeftID(ID []byte) []byte {
	padding := make([]byte, 16-len(ID))
	return append(padding, ID...)
}

func TrimRightID(ID []byte) []byte {
	for i := len(ID) - 1; i >= 0; i-- {
		if ID[i] != 0x00 {
			return ID[:i+1]
		}
	}

	return ID
}

func PadRightID(ID []byte) []byte {
	padding := make([]byte, 16-len(ID))
	return append(ID, padding...)
}
