package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"time"
)

func SaveECDSAPublicKey(pub ecdsa.PublicKey, name string) error {
	filename := fmt.Sprintf("keys/%s_pub.pem", name)
	pubFile, err := os.OpenFile(filename, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer pubFile.Close()

	x509EncodedPublicKey, _ := x509.MarshalPKIXPublicKey(&pub)

	if err := pem.Encode(pubFile, &pem.Block{
		Type: "ECDSA PUBLIC KEY",
		Headers: map[string]string{
			"CREATED AT": time.Now().String(),
			"FOR":        name,
		},
		Bytes: x509EncodedPublicKey,
	}); err != nil {
		log.Printf("Failed to save public key : %v\n", err)
		return err
	}

	return nil
}
