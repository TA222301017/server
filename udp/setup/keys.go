package setup

import (
	"crypto/ecdsa"
	"log"
	"os"
	"path/filepath"
	"server/udp/utils"
)

var (
	PrivateKey *ecdsa.PrivateKey = nil
	PublicKey  *ecdsa.PublicKey  = nil
)

func generateKeys() error {
	PrivateKey, err := utils.GenerateECDSAKeys()
	if err != nil {
		return err
	}
	PublicKey = &PrivateKey.PublicKey

	if err := utils.SaveECDSAPrivateKey(PrivateKey, "server"); err != nil {
		log.Println(err.Error)
		return err
	}

	if err := utils.SaveECDSAPublicKey(PublicKey, "server"); err != nil {
		log.Println(err.Error)
		return err
	}

	log.Println("Generated server ECDSA keys")

	return nil
}

func loadKeys() error {
	privKey, err := utils.LoadECDSAPrivateKey("server")
	if err != nil {
		log.Println(err.Error)
		return err
	}

	pubKey, err := utils.LoadECDSAPublicKey("server")
	if err != nil {
		log.Println(err.Error)
		return err
	}

	PrivateKey = privKey
	PublicKey = pubKey

	return nil
}

func Keys() {
	path := filepath.Join(".", "keys")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	err := loadKeys()
	if err != nil {
		log.Printf("Failed to load keys : %v\n", err)
		if err := generateKeys(); err != nil {
			log.Panicf("Failed to generate keys : %v\n", err)
		}
	}
}
