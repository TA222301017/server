package setup

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"io"
	"log"
	"os"
	"time"
)

var (
	PrivateKey *ecdsa.PrivateKey = nil
	PublicKey  *ecdsa.PublicKey  = nil
)

func generateKeys() error {
	PrivateKey, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		return err
	}
	PublicKey = &PrivateKey.PublicKey

	x509EncodedPrivateKey, _ := x509.MarshalECPrivateKey(PrivateKey)
	x509EncodedPublicKey, _ := x509.MarshalPKIXPublicKey(&PrivateKey.PublicKey)

	pubFile, err := os.OpenFile("keys/server_pub.pem", os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	privFile, err := os.OpenFile("keys/server_priv.pem", os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer pubFile.Close()
	defer privFile.Close()

	if err := pem.Encode(privFile, &pem.Block{
		Type: "SERVER ECDSA PRIVATE KEY",
		Headers: map[string]string{
			"CREATED AT": time.Now().String(),
		},
		Bytes: x509EncodedPrivateKey,
	}); err != nil {
		log.Printf("Failed to save secret key : %v\n", err)
		return err
	}

	if err := pem.Encode(pubFile, &pem.Block{
		Type: "SERVER ECDSA PUBLIC KEY",
		Headers: map[string]string{
			"CREATED AT": time.Now().String(),
		},
		Bytes: x509EncodedPublicKey,
	}); err != nil {
		log.Printf("Failed to save public key : %v\n", err)
		return err
	}
	log.Println("Generated server ECDSA keys")

	return nil
}

func loadKeys() error {
	log.Println("Opening server secret key and public key files...")
	pubFile, err := os.Open("keys/server_pub.pem")
	if err != nil {
		log.Printf("Failed to open keys/server_pub.pem : %v", err)
		return err
	}

	privFile, err := os.Open("keys/server_priv.pem")
	if err != nil {
		log.Printf("Failed to open keys/server_priv.pem : %v", err)
		return err
	}

	pubBytes, err := io.ReadAll(pubFile)
	if err != nil {
		log.Printf("Failed to read keys/server_pub.pem : %v", err)
		return err
	}

	privBytes, err := io.ReadAll(privFile)
	if err != nil {
		log.Printf("Failed to read keys/server_priv.pem : %v", err)
		return err
	}

	log.Println("Decoding key files...")
	pubBlock, _ := pem.Decode(pubBytes)
	privBlock, _ := pem.Decode(privBytes)

	// Loading keys to variables
	PrivateKey, err = x509.ParseECPrivateKey(privBlock.Bytes)
	if err != nil {
		log.Printf("Failed to decode server private key : %v", err)
		return err
	}

	publicKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		log.Printf("Failed to decode server private key : %v", err)
		return err
	}
	PublicKey = publicKey.(*ecdsa.PublicKey)

	return nil
}

func Keys() {
	if err := loadKeys(); err != nil {
		log.Printf("Failed to load keys : %v\n", err)
		if err := generateKeys(); err != nil {
			log.Panicf("Failed to load keys : %v\n", err)
		}
	}
}
