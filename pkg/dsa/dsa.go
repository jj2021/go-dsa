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

	// generate p
	for i := 0; i < 4*L; i++ {
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

	// generate value g
	g := new(big.Int)
	e := new(big.Int)
	pSubOne := new(big.Int)
	hMax := new(big.Int)
	one := big.NewInt(1)
	h := new(big.Int)

	pSubOne.Sub(params.P, one)
	e.Div(pSubOne, params.Q)

	hMax.Sub(pSubOne, one)
	h, err := rand.Int(rand.Reader, hMax)
	if err != nil {
		fmt.Printf("Could not generate random value h\n")
		return params
	}

	// ensure h is not zero
	h.Add(h, one)

	g.Exp(h, e, params.P)

	params.G = g

	return params
}

func Sign() {

}

func Verify() {

}
