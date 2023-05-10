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
	msgString := "00000000006459F52100000000000000000000000000000000042E4D8810B4CF5678BC228E7DE7B69FA8481477709BA3A0523861B86898BA42DB41E844647CD2655FB5B739D35B31C7ABEE2B67ADC34C33082E7ED389583EABD1"
	sigString := "022100C3C9EEF848B09A9B05150E88BEDC331BB18C5931C2A67749D753FD41B0057F4D022100BF6DB62132A88C01811F5807B24676E1FAE50335A882D5C1ED74C91A0E72DA25"
	pKeyString := "042E4D8810B4CF5678BC228E7DE7B69FA8481477709BA3A0523861B86898BA42DB41E844647CD2655FB5B739D35B31C7ABEE2B67ADC34C33082E7ED389583EABD1"
	hashString := "405EF917FD716985BD4009ADBB4308E44BFC1B2494CCD6264E35141E8169A798"
	// hashString := "2B8A70A4839F8BB18102AA880534DA7C5AD78721C1BFE94602D512DE5FA5C2AB"

	// msg := 'SAC_LOCK'
	msgBytes, err := hex.DecodeString(msgString)
	if err != nil {
		fmt.Println("gagal parsing message")
		return
	}
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
	fmt.Println("HASH", hash)
	fmt.Println("HASH FROM MSG", d)

	if ecdsa.VerifyASN1(pKey, hash, sig) {
		fmt.Println("ok")
	} else {
		fmt.Println("fucek")
	}
}
