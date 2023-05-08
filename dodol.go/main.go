package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"
	"errors"
	"fmt"
)

func ParseECDSAPublickKey(pubKey []byte) (*ecdsa.PublicKey, error) {
	x, y := elliptic.Unmarshal(elliptic.P256(), pubKey)
	if x == nil {
		return nil, errors.New("invalid public key")
	}

	return &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}, nil
}

func reverseBytes(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	sigString := "022004DB10FC4D0A58A845E6A933708ED13FFB1354F2AF7C25DCEF283AA6B3054BC00221008B7899940BAACD99ECD79653C7FC204DD96C2C77D9DC4CF867AFBA69E197141F"
	pKeyString := "043844646C9FD36602CA18DF5ABF2BC34ECA730FCB39EC77FBD44F2BE50360ABA51FD7533593672333A2952C499A99A7452EFE11B59E075DE8C13F9C74C5FDF07B"
	hashString := "ABC2A55FDE12D50246E9BFC12187D75A7CDA340588AA0281B18B9F83A4708A2B"
	// hashString := "2B8A70A4839F8BB18102AA880534DA7C5AD78721C1BFE94602D512DE5FA5C2AB"

	// msg := 'SAC_LOCK'
	msgBytes := []byte{83, 65, 67, 95, 76, 79, 67, 75, 0}
	// msgBytes := []byte(msg)
	// reverseBytes(msgBytes)
	t := sha256.Sum256(msgBytes)
	d := t[:]

	pKeyBytes, err := hex.DecodeString(pKeyString)
	if err != nil {
		fmt.Println("gagal parsing pkey sbg hex")
		return
	}

	sig, err := hex.DecodeString(sigString)
	if err != nil {
		fmt.Println("gagal parsing sig")
		return
	}

	hash, err := hex.DecodeString(hashString)
	if err != nil {
		fmt.Println("gagal parsing hash")
		return
	}

	pKey, err := ParseECDSAPublickKey(pKeyBytes)
	if err != nil {
		fmt.Println("gagal parsing pkey")
		return
	}

	// reverseBytes(sig)
	// reverseBytes(d)
	reverseBytes(hash)

	var r asn1.RawValue
	var s asn1.RawValue
	rest, err := asn1.Unmarshal(sig, &r)
	if err != nil {
		panic(err)
	} else {
		// fmt.Println(r)
	}

	if _, err := asn1.Unmarshal(rest, &s); err != nil {
		panic(err)
	} else {
		// fmt.Println(s)
	}
	// sig = append([]byte{0x00}, sig...)
	sig = append([]byte{byte(len(sig))}, sig...)
	sig = append([]byte{0x30}, sig...)

	fmt.Println(sig)
	fmt.Println(hash)
	fmt.Println(d)

	if ecdsa.VerifyASN1(pKey, d, sig) {
		fmt.Println("ok")
	} else {
		fmt.Println("fucek")
	}
}
