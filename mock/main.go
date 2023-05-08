package main

import (
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"server/udp/template"
	"server/udp/utils"
	"strings"
	"time"
)

type accessRule struct {
	ID       uint64
	LockID   string
	KeyID    string
	StartsAt string
	EndsAt   string
}

var accessRules map[uint64]accessRule = make(map[uint64]accessRule)

func randomBytes(l int) []byte {
	var ret []byte = make([]byte, l)
	rand.Read(ret)
	return ret
}

func printAccessRules() {
	b, err := json.MarshalIndent(accessRules, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}

func main() {
	args := os.Args
	if len(args) == 1 {
		os.Exit(1)
	}
	privKey, err := utils.LoadECDSAPrivateKey("lock")
	pubKey, err := utils.LoadECDSAPublicKey("lock")
	lockID := randomBytes(16)
	keyID, _ := hex.DecodeString("DEADBEEFDEADBEEFBEEFDEADDEADBEEF")

	if err != nil {
		fmt.Println(err)
		privKey, _ = utils.GenerateECDSAKeys()
		pubKey = &privKey.PublicKey

		utils.SaveECDSAPrivateKey(privKey, "lock")
		utils.SaveECDSAPublicKey(pubKey, "lock")
	}

	lockIDFile, err := os.Open("lock_id.txt")
	defer lockIDFile.Close()
	if err == nil {
		id, err := io.ReadAll(lockIDFile)
		if err == nil {
			hex.Decode(lockID, id)
		}
	} else {
		file, _ := os.Create("lock_id.txt")
		defer file.Close()
		file.Write([]byte(hex.EncodeToString(lockID)))
	}

	switch args[1] {
	case "keyex":
		pubKeyBytes := elliptic.Marshal(elliptic.P256(), pubKey.X, pubKey.Y)
		packet, err := utils.MakePacket(template.KeyExchange, append(lockID, pubKeyBytes...), privKey)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(hex.EncodeToString(pubKeyBytes))

		fmt.Println("Packet:", strings.ToUpper(hex.EncodeToString(packet.Bytes())))
		res, err := utils.SendUDPPacket(packet, "127.0.0.1", "8888")
		if err != nil {
			fmt.Println(err)
			return
		}

		serverPublicKey, err := utils.ParseECDSAPublickKey(res.Data[24:])
		if err != nil {
			fmt.Println(err)
			return
		}

		err = utils.VerifyPacket(res.Bytes(), serverPublicKey)
		fmt.Println("Verified: ", err == nil)

		if err := utils.SaveECDSAPublicKey(serverPublicKey, "server"); err != nil {
			fmt.Println(err)
			return
		}

	case "access":
		packet, err := utils.MakePacket(template.LogAccessEvent, append(lockID, keyID...), privKey)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Packet:", strings.ToUpper(hex.EncodeToString(packet.Bytes())))
		res, err := utils.SendUDPPacket(packet, "127.0.0.1", "8888")
		if err != nil {
			fmt.Println(err)
			return
		}

		serverPublicKey, err := utils.LoadECDSAPublicKey("server")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = utils.VerifyPacket(res.Bytes(), serverPublicKey)
		fmt.Println("Verified: ", err == nil)

	case "rssi":
		data := append(lockID, keyID...)
		data = append(data, randomBytes(1)...)
		packet, err := utils.MakePacket(template.LogRSSIEvent, data, privKey)
		fmt.Println(hex.EncodeToString(packet.Bytes()))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Packet:", strings.ToUpper(hex.EncodeToString(packet.Bytes())))
		res, err := utils.SendUDPPacket(packet, "127.0.0.1", "8888")
		if err != nil {
			fmt.Println(err)
			return
		}

		serverPublicKey, err := utils.LoadECDSAPublicKey("server")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = utils.VerifyPacket(res.Bytes(), serverPublicKey)
		fmt.Println("Verified: ", err == nil)

	case "listen":
		serverPublicKey, err := utils.LoadECDSAPublicKey("server")
		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 2048)
		conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8000})
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		for {
			pLen, remoteAddr, err := conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println(err)
				continue
			}

			packet, err := utils.ParsePacket(buffer, pLen)
			if err != nil {
				fmt.Println(err)
				continue
			}

			switch packet.OpCode {
			case template.AddAccessRule:
				err = utils.VerifyPacket(packet.Bytes(), serverPublicKey)
				fmt.Println("Verified: ", err == nil)
				if err != nil {
					return
				}

				accessRuleID := binary.BigEndian.Uint64(packet.Data[:8])
				lockID := packet.Data[8:24]
				keyID := packet.Data[40:56]
				startsAt := time.Unix(int64(binary.BigEndian.Uint64(packet.Data[56:64])), 0)
				endsAt := time.Unix(int64(binary.BigEndian.Uint64(packet.Data[64:72])), 0)

				accessRules[accessRuleID] = accessRule{
					ID:       accessRuleID,
					LockID:   hex.EncodeToString(lockID),
					KeyID:    hex.EncodeToString(keyID),
					StartsAt: startsAt.String(),
					EndsAt:   endsAt.String(),
				}
				printAccessRules()

				res, err := utils.MakePacket(
					template.AddAccessRule,
					packet.Data,
					privKey,
				)
				if err != nil {
					fmt.Println(err)
					return
				}
				conn.WriteToUDP(res.Bytes(), remoteAddr)

			case template.EditAccessRule:
				err = utils.VerifyPacket(packet.Bytes(), serverPublicKey)
				fmt.Println("Verified: ", err == nil)
				if err != nil {
					return
				}

				accessRuleID := binary.BigEndian.Uint64(packet.Data[:8])
				lockID := packet.Data[8:24]
				keyID := packet.Data[40:56]
				startsAt := time.Unix(int64(binary.BigEndian.Uint64(packet.Data[56:64])), 0)
				endsAt := time.Unix(int64(binary.BigEndian.Uint64(packet.Data[64:72])), 0)

				accessRules[accessRuleID] = accessRule{
					ID:       accessRuleID,
					LockID:   hex.EncodeToString(lockID),
					KeyID:    hex.EncodeToString(keyID),
					StartsAt: startsAt.String(),
					EndsAt:   endsAt.String(),
				}
				printAccessRules()

				res, err := utils.MakePacket(
					template.AddAccessRule,
					packet.Data,
					privKey,
				)
				if err != nil {
					fmt.Println(err)
					return
				}
				conn.WriteToUDP(res.Bytes(), remoteAddr)

			case template.DeleteAccessRule:
				err = utils.VerifyPacket(packet.Bytes(), serverPublicKey)
				fmt.Println("Verified: ", err == nil)
				if err != nil {
					return
				}

				accessRuleID := binary.BigEndian.Uint64(packet.Data[:8])

				delete(accessRules, accessRuleID)
				printAccessRules()

				res, err := utils.MakePacket(
					template.AddAccessRule,
					packet.Data,
					privKey,
				)
				if err != nil {
					fmt.Println(err)
					return
				}
				conn.WriteToUDP(res.Bytes(), remoteAddr)

			case template.LogHealthcheckEvent:
				err = utils.VerifyPacket(packet.Bytes(), serverPublicKey)
				fmt.Println("Verified: ", err == nil)
				if err != nil {
					fmt.Println(err)
					return
				}

				temp := packet.Data[:16]
				fmt.Println(temp)
				fmt.Println("Lock ID match: ", 0 == bytes.Compare(temp, lockID))
				temp = packet.Data[16:]
				fmt.Println("Public Key match: ", 0 == bytes.Compare(temp, utils.MarshalECDSAPublicKey(pubKey)))

				res, err := utils.MakePacket(
					template.LogHealthcheckEvent,
					packet.Data,
					privKey,
				)
				if err != nil {
					fmt.Println(err)
					return
				}
				conn.WriteToUDP(res.Bytes(), remoteAddr)
			}
		}

	default:
		fmt.Printf("\nLock Mock - Program untuk mensimulasikan subsistem lock\n")
		fmt.Println("Daftar Perintah: ")
		fmt.Printf(" keyex  : Kirim Paket KeyExchange\n")
		fmt.Printf(" access : Kirim Paket LogAccessEvent\n")
		fmt.Printf(" rssi   : Kirim Paket LogRSSIEvent\n")
		fmt.Printf(" listen : Listen\n\n")
		fmt.Println("Penggunaan: ")
		fmt.Println("  go run main.go <perintah>\n")
		os.Exit(1)
	}

	fmt.Print("\n")
}
