package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("%s <key-filename>\n", os.Args[0])
		return
	}
	
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Ed25519 Private Key: %s", hex.EncodeToString(pri))
	log.Printf("Ed25519 Public Key: %s", hex.EncodeToString(pub))

	priFileName := os.Args[1]
	pubFileName := fmt.Sprintf("%s.pub", priFileName)

	pkcsBytes, err := x509.MarshalPKCS8PrivateKey(pri)
	if err != nil {
		log.Fatal(err)
	}

	pemBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pkcsBytes,
	}

	priFile, err := os.Create(priFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer priFile.Close()

	if err := pem.Encode(priFile, pemBlock); err != nil {
		log.Fatal(err)
	}

	log.Printf("The Ed25519 Private key was saved in PKCS#8 specification at %s", priFile.Name())

	// Write Public Key

	pkixBytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		log.Fatal(err)
	}

	pemBlock = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pkixBytes,
	}

	pubFile, err := os.Create(pubFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer pubFile.Close()

	if err := pem.Encode(pubFile, pemBlock); err != nil {
		log.Fatal(err)
	}

	log.Printf("The Ed25519 Public key was saved in PKIX/X.509 - SubjectPublicKeyInfo specification at %s", pubFile.Name())
}
