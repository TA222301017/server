package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func GenerateECDSAKeys() (*ecdsa.PrivateKey, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

func LoadECDSAPrivateKey(name string) (*ecdsa.PrivateKey, error) {
	var (
		privateKey   *ecdsa.PrivateKey
		privFilePath string = fmt.Sprintf("keys/%s_priv.pem", name)
	)

	privFile, err := os.Open(privFilePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open %s : %s", privFilePath, err)
	}
	defer privFile.Close()

	privBytes, err := io.ReadAll(privFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to read %s : %s", privFilePath, err)
	}

	privBlock, _ := pem.Decode(privBytes)

	privateKey, err = x509.ParseECPrivateKey(privBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode private key : %s", err)
	}

	return privateKey, nil
}

func LoadECDSAPublicKey(name string) (*ecdsa.PublicKey, error) {
	var (
		publicKey   *ecdsa.PublicKey
		pubFilePath string = fmt.Sprintf("keys/%s_pub.pem", name)
	)

	pubFile, err := os.Open(pubFilePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open %s : %s", pubFilePath, err)
	}
	defer pubFile.Close()

	pubBytes, err := io.ReadAll(pubFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to read %s : %s", pubFilePath, err)
	}

	pubBlock, _ := pem.Decode(pubBytes)

	temp, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode public key : %s", err)
	}
	publicKey = temp.(*ecdsa.PublicKey)

	return publicKey, nil
}

func SaveECDSAPublicKey(pub *ecdsa.PublicKey, name string) error {
	filename := fmt.Sprintf("keys/%s_pub.pem", name)
	pubFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer pubFile.Close()

	x509EncodedPublicKey, _ := x509.MarshalPKIXPublicKey(pub)

	if err := pem.Encode(pubFile, &pem.Block{
		Type: "EC PUBLIC KEY",
		Headers: map[string]string{
			"CREATED AT": time.Now().String(),
			"FOR":        name,
		},
		Bytes: x509EncodedPublicKey,
	}); err != nil {
		return fmt.Errorf("Failed to save public key : %s\n", err)
	}

	return nil
}

func SaveECDSAPrivateKey(priv *ecdsa.PrivateKey, name string) error {
	filename := fmt.Sprintf("keys/%s_priv.pem", name)
	privFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer privFile.Close()

	x509EncodedPrivateKey, _ := x509.MarshalECPrivateKey(priv)

	if err := pem.Encode(privFile, &pem.Block{
		Type: "EC PRIVATE KEY",
		Headers: map[string]string{
			"CREATED AT": time.Now().String(),
			"FOR":        name,
		},
		Bytes: x509EncodedPrivateKey,
	}); err != nil {
		return fmt.Errorf("Failed to save private key : %s\n", err)
	}

	return nil
}

func MarshalECDSAPublicKey(pubKey *ecdsa.PublicKey) []byte {
	return elliptic.Marshal(elliptic.P256(), pubKey.X, pubKey.Y)
}

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
