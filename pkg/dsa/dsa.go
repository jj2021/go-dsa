package dsa

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Printf("dsa implementation\n")
}

type KeyPair struct {
	Params  Parameters
	Public  publicKey
	Private privateKey
}

type Parameters struct {
	P, Q, G *big.Int
}

type privateKey *big.Int
type publicKey *big.Int

func GenerateKeyPair() KeyPair {
	params := generateGlobalParameters()
	return KeyPair{Params: params}
}

func generateGlobalParameters() Parameters {
	// Will always use the smallest possible size
	// as this implementation will only support
	// the SHA1 hash function
	rounds := 40
	L := 1024
	N := 160
	fmt.Printf("L: %v, N: %v\n", L, N)

	// initialize variables
	params := Parameters{}
	g := new(big.Int)

	for {
		// generate q
		q, err := rand.Prime(rand.Reader, N)
		if err != nil {
			fmt.Printf("Error Generating q: %v", err.Error())
			return params
		}

		// Ensure q is sufficiently large by setting the first
		// bit as 1
		qBytes := q.Bytes()
		qBytes[0] = 0x80
		q.SetBytes(qBytes)

		if !q.ProbablyPrime(rounds) {
			continue
		}
		params.Q = q
		break
	}

	fmt.Printf("q: %v\n", params.Q)

	for i := 0; i < 4*L; i++ {
		// generate p
		p, err := rand.Prime(rand.Reader, L)
		if err != nil {
			fmt.Printf("Error Generating p: %v", err.Error())
			return params
		}

		// Ensure p is sufficiently large by setting the first
		// bit as 1
		pBytes := p.Bytes()
		pBytes[0] = 0x80
		p.SetBytes(pBytes)

		// create p such that q | p - 1
		remainder := new(big.Int)
		diff := new(big.Int)
		one := big.NewInt(1)
		two := big.NewInt(2)

		doubleQ := new(big.Int)
		doubleQ.Mul(two, params.Q)
		remainder.Mod(p, doubleQ)

		diff.Sub(remainder, one)
		p.Sub(p, diff)

		if !p.ProbablyPrime(rounds) {
			continue
		}

		params.P = p
		break
	}

	params.G = g

	return params
}

func Sign() {

}

func Verify() {

}
