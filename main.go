package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/wire"
)

const ProtocolVersion = 70002

// pre 1.3 // var GuldencoinTestnet wire.BitcoinNet = 0xdcb7c1fc // 0xfc, 0xc1, 0xb7, 0xdc
var GuldencoinTestnet wire.BitcoinNet = 0x00f7fefc //0xfc, 0xfe, 0xf7, 0x00

// pre 1.3 // var GuldencoinMainnet wire.BitcoinNet = 0xdbb6c0fb // 0xfb, 0xc0, 0xb6, 0xdb
var GuldencoinMainnet wire.BitcoinNet = 0xe0f7fefc // 0xfc, 0xfe, 0xf7, 0xe0

var SelectedGuldencoinNet = GuldencoinMainnet

var PrivKeyBase64Encoded = ``

func main() {

	// create an alert
	timeLocation, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		log.Fatalf("Could not parse time location: %v", err)
	}
	relayUntil, err := time.ParseInLocation("2006-01-02 15:04:05", "2015-12-30 10:00:00", timeLocation)
	if err != nil {
		log.Fatalf("Could not parse relayUntil time: %v", err)
	}
	expiration, err := time.ParseInLocation("2006-01-02 15:04:05", "2015-12-31 14:00:00", timeLocation)
	if err != nil {
		log.Fatalf("Could not parse expiration time: %v", err)
	}
	// alert := &wire.Alert{
	// 	Version:    1,
	// 	RelayUntil: relayUntil.Unix(),
	// 	Expiration: expiration.Unix(),
	// 	ID:         1,
	// 	Cancel:     0,
	// 	SetCancel:  nil,
	// 	MinVer:     0,
	// 	MaxVer:     ProtocolVersion + 1,
	// 	SetSubVer:  nil,
	// 	Priority:   10,
	// Comment:    "IMPORTANT: A new version of the Guldencoin wallet will be released 25 september. Updating is mandatory! Please read guldencoin.com/dgw3",
	// StatusBar: "IMPORTANT: Upcoming mandatory update 25 sept. Please read guldencoin.com/dgw3\n" +
	// 	"BELANGRIJK: Verplichte wallet update op 25 september. Lees a.u.b. guldencoin.com/nl/dgw3",
	// 	Reserved: "",
	// }
	// alert := &wire.Alert{
	// 	Version:    1,
	// 	RelayUntil: relayUntil.Unix(), // 2014-09-25 18:58:00
	// 	Expiration: expiration.Unix(), // 2014-09-25 20:00:00
	// 	ID:         2,
	// 	Cancel:     0,
	// 	SetCancel:  nil,
	// 	MinVer:     0,
	// 	MaxVer:     ProtocolVersion + 1,
	// 	SetSubVer:  nil,
	// 	Priority:   20,
	// 	Comment:    "IMPORTANT: A new version of the Guldencoin wallet will be released 25 september. Updating is mandatory! Please read guldencoin.com/dgw3",
	// 	StatusBar: "IMPORTANT: Shut down this wallet and download the new version 1.3 in about an hour! Please read guldencoin.com/dgw3\n" +
	// 		"BELANGRIJK: Sluit deze wallet af en download de nieuwe versie 1.3 over ongeveer een uur. Lees a.u.b. guldencoin.com/nl/dgw3",
	// 	Reserved: "",
	// }
	// alert := &wire.Alert{
	// 	Version:    1,
	// 	RelayUntil: relayUntil.Unix(),
	// 	Expiration: expiration.Unix(),
	// 	ID:         3,
	// 	Cancel:     0,
	// 	SetCancel:  nil,
	// 	MinVer:     0,
	// 	MaxVer:     ProtocolVersion + 1,
	// 	SetSubVer:  []string{"/Guldencoin:1.3.0/"}, // nil
	// 	Priority:   20,
	// 	Comment:    "IMPORTANT: A new version of the Guldencoin is released. Updating is mandatory!",
	// 	StatusBar: "IMPORTANT: Shut down this wallet and download the new version 1.3.1! Very important and mandatory update!!\n" +
	// 		"BELANGRIJK: Sluit deze wallet af en download de nieuwe versie 1.3.1! Zeer belangrijke update!",
	// 	Reserved: "",
	// }
	// alert := &wire.Alert{
	// 	Version:    1,
	// 	RelayUntil: relayUntil.Unix(),
	// 	Expiration: expiration.Unix(),
	// 	ID:         3,
	// 	Cancel:     0,
	// 	SetCancel:  nil,
	// 	MinVer:     0,
	// 	MaxVer:     ProtocolVersion + 1,
	// 	SetSubVer:  []string{"/Guldencoin:1.3.1/"}, // nil
	// 	Priority:   20,
	// 	Comment:    "IMPORTANT: A new version of the Guldencoin is released. Updating is mandatory!",
	// 	StatusBar: "IMPORTANT: Shut down this wallet and download the new version 1.4.0! Very important and mandatory update!!\n" +
	// 		"BELANGRIJK: Sluit deze wallet af en download de nieuwe versie 1.4.0! Zeer belangrijke update!",
	// 	Reserved: "",
	// }
	// alert := &wire.Alert{
	// 	Version:    1,
	// 	RelayUntil: relayUntil.Unix(),
	// 	Expiration: expiration.Unix(),
	// 	ID:         3,
	// 	Cancel:     0,
	// 	SetCancel:  nil,
	// 	MinVer:     0,
	// 	MaxVer:     ProtocolVersion + 1,
	// 	SetSubVer:  []string{"/Guldencoin:1.4.0/"}, // nil
	// 	Priority:   20,
	// 	Comment:    "IMPORTANT: A new version of the Guldencoin wallet is released. Updating is mandatory!",
	// 	StatusBar: "IMPORTANT: Shut down this wallet and download the new version 1.5.0! Very important and mandatory update!!\n" +
	// 		"BELANGRIJK: Sluit deze wallet direct af en download de nieuwe versie 1.5.0! Zeer belangrijke update!",
	// 	Reserved: "",
	// }
	alert := &wire.Alert{
		Version:    1,
		RelayUntil: relayUntil.Unix(),
		Expiration: expiration.Unix(),
		ID:         4,
		Cancel:     0,
		SetCancel:  nil,
		MinVer:     0,
		MaxVer:     ProtocolVersion + 1,
		SetSubVer:  []string{"/Guldencoin:1.5.0/"}, // nil
		Priority:   20,
		Comment:    "IMPORTANT: This version of the Gulden wallet is outdated. Please update.",
		StatusBar: "IMPORTANT: Please shut down this wallet and download the new version at https://gulden.com\n" +
			"BELANGRIJK: Sluit a.u.b. af en download de nieuwe versie op https://gulden.com",
		Reserved: "",
	}
	alertBuf := &bytes.Buffer{}
	err = alert.Serialize(alertBuf, ProtocolVersion)
	if err != nil {
		log.Fatalf("error serializing wire.Alert: %v", err)
	}
	alertBufBytes := alertBuf.Bytes()

	// decode privKey from string
	privKeyBytes, err := base64.StdEncoding.DecodeString(PrivKeyBase64Encoded)
	if err != nil {
		log.Fatalf("error base64 decoding PrivKey: %v", err)
	}

	// get priv and pub key from privKeyBytes
	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)
	pubKeySerializedHybrid := pubKey.SerializeHybrid()
	log.Printf("pubKey hybrid: `%s`", hex.EncodeToString(pubKeySerializedHybrid))

	// create signature
	r, s, err := ecdsa.Sign(rand.Reader, privKey.ToECDSA(), wire.DoubleSha256(alertBufBytes))
	if err != nil {
		log.Fatalf("error signing alert: %v", err)
	}
	signature := &btcec.Signature{
		R: r,
		S: s,
	}
	signatureBytes := signature.Serialize()

	// connect to localhost node
	c, err := net.Dial("tcp", "seed-000.gulden.com:9231") // 141.138.139.53:9923 seed-000.guldencoin.net:9231
	if err != nil {
		log.Fatalf("unable to connect(dial) to external guldencoin process: %v", err)
	}

	// create and send version message to start talking with node
	msgVersion, err := wire.NewMsgVersionFromConn(c, 1234, 1)
	if err != nil {
		log.Fatalf("error creating msg version from conn: %v", err)
	}
	msgVersion.ProtocolVersion = ProtocolVersion
	err = wire.WriteMessage(c, msgVersion, ProtocolVersion, SelectedGuldencoinNet)
	if err != nil {
		log.Fatalf("error writing version message: %v", err)
	}

	// create alert message
	msgAlert := wire.NewMsgAlert(alertBufBytes, signatureBytes)

	// write message hexdump to stdout
	err = msgAlert.BtcEncode(hex.Dumper(os.Stdout), ProtocolVersion)
	fmt.Println()
	if err != nil {
		log.Fatalf("error btcEndocing msgAlert: %v", err)
	}

	// parse pubkey
	checkPubKey, err := btcec.ParsePubKey(pubKeySerializedHybrid, btcec.S256())
	if err != nil {
		log.Fatalf("error parsing checkPubKey: %v", err)
	}

	// dump and parse signature
	log.Printf("signature: \n%s", hex.Dump(msgAlert.Signature))
	parsedSignature, err := btcec.ParseSignature(msgAlert.Signature, btcec.S256())
	if err != nil {
		log.Fatalf("error parsing signature to check: %v", err)
	}

	// verify signature
	checkOk := ecdsa.Verify(checkPubKey.ToECDSA(), wire.DoubleSha256(msgAlert.SerializedPayload), parsedSignature.R, parsedSignature.S)
	if !checkOk {
		log.Fatalf("signature could not be verified locally")
	}

	// write message to node
	err = wire.WriteMessage(c, msgAlert, ProtocolVersion, SelectedGuldencoinNet)
	if err != nil {
		log.Fatalf("error writing alert message: %v", err)
	}

	log.Println("all done")
	select {}
}

