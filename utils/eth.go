package utils

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func GetAddressFromPrivateKeyStr(privateKeyStr string) (address string) {
	address = ""
	privkeyECDSA, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Println("error generating private key")
	}

	publicKey := privkeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return
}
