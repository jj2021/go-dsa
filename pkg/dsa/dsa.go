package dsa

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Printf("dsa implementation\n")
}

type KeyPair struct {
	P       big.Int
	Q       big.Int
	G       big.Int
	Public  publicKey
	Private privateKey
}

type privateKey big.Int
type publicKey big.Int

func GenerateKeyPair() KeyPair {
	return KeyPair{}
}

func Sign() {

}

func Verify() {

}
