package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"server/udp/utils"
)

// NOTES:
// 1. 3 byte pertama di public key bikin key nya jd invalid, gatau kenapa
// 2. 2 byte data harus diprepend ke signature, byte pertama nilainya 03, byte kedua panjang signature yg sebelum dimodif dalam hex
// 3. hasil hashing di psoc sama hasil hashing di server urutan bytenya kebalik, kudu deal2an dulu mau pake urutan byte yg mana
// 4. setelah signature dibuat dan ditambahin data dr poin 2, kasih padding di awal signature biar panjangnya 72 byte
type R struct {
	Value asn1.RawValue
}

type S struct {
	Value asn1.RawValue
}

func main() {
	signature, _ := hex.DecodeString("3044022069270FED671BACE8D91A4F0B199F67F368C0BCBE68ED1A9BCB533DEEEC4FB33602205BEE04D556678415B481ED84DC6F1C34959E6F91C27C22707B44C9859E4D4397")
	// sb, _ := hex.DecodeString("3022")
	pKeyBytes, _ := hex.DecodeString("047E8BBC0D9F7B22749661AAA140C2D892784B8080D2A0ABE544EDA97843482BF312D4EB355AFEA393AAAB5DB62ED7804BB47676F05194A299DEF9375C82D4B771")

	// kp, _ := utils.GenerateECDSAKeys()
	// s := hex.EncodeToString(elliptic.Marshal(elliptic.P256(), kp.PublicKey.X, kp.PublicKey.Y))
	// fmt.Println(s, len(s)/2)

	pKey, err := utils.ParseECDSAPublickKey(pKeyBytes)
	if err != nil {
		panic(err)
	}

	message := append([]byte("SAC_LOCK"), 0)
	t := sha256.Sum256(message)
	digest := t[:]

	for i := 31; i >= 0; i-- {
		digest = append(digest, t[i])
		fmt.Println(digest)
	}

	// var r asn1.RawValue
	// var s asn1.RawValue
	// if _, err := asn1.Unmarshal(rb, &r); err != nil {
	// 	panic(err)
	// }

	// if _, err := asn1.Unmarshal(sb, &s); err != nil {
	// 	panic(err)
	// }

	// if !ecdsa.Verify(pKey, message, r.Value, s.V) {
	// 	fmt.Println("ASDF")
	// }

	// err = utils.VerifyPacket(append(message[:], signature...), pKey)
	if !ecdsa.VerifyASN1(pKey, digest[:], signature) {
		fmt.Println("error: invalid signature")
	} else {
		fmt.Println("ok")
	}
}
