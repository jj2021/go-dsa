package dsa

import (
	"crypto/rand"
	"fmt"
	"godsa/pkg/sha"
	"math/big"
)

type KeyPair struct {
	Params  Parameters
	Public  publicKey
	Private privateKey
}

type Parameters struct {
	P, Q, G *big.Int
}

type privateKey struct {
	*big.Int
}

type publicKey struct {
	*big.Int
}

type Signature struct {
	r *big.Int
	s *big.Int
}

func GenerateKeyPair() KeyPair {
	var pair KeyPair
	params := generateGlobalParameters()
	pair.Params = params
	x, y := generateKeys(params)
	pair.Private = x
	pair.Public = y
	return pair
}

func generateKeys(params Parameters) (privateKey, publicKey) {
	priv := new(big.Int)
	pub := new(big.Int)
	/*
		qbytes := params.Q.Bytes()
		fmt.Printf("q len: %v\n", len(qbytes))
		pbytes := params.P.Bytes()
		fmt.Printf("p len: %v\n", len(pbytes))
	*/

	// generate rand number of N bits
	c := new(big.Int)
	valid := false

	for !valid {
		n := 20
		b := make([]byte, n)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Printf("Could not generate random bits: %v\n", err.Error())
			return privateKey{priv}, publicKey{pub}
		}

		c.SetBytes(b)

		qSubTwo := new(big.Int)
		two := big.NewInt(2)
		qSubTwo.Sub(params.Q, two)
		if c.Cmp(qSubTwo) != 1 {
			valid = true
		}
	}

	// calc priv
	one := big.NewInt(1)
	priv.Add(c, one)

	// calc pub
	pub.Exp(params.G, priv, params.P)

	return privateKey{priv}, publicKey{pub}
}

func generateGlobalParameters() Parameters {
	// Will always use the smallest possible size
	// as this implementation will only support
	// the SHA1 hash function
	rounds := 40
	L := 1024
	N := 160
	//fmt.Printf("L: %v, N: %v\n", L, N)

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

	//fmt.Printf("q: %v\n", params.Q)

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

func Sign(content []byte, privKey privateKey, params Parameters) Signature {
	var signature Signature
	r := big.NewInt(0)
	s := big.NewInt(0)
	z := new(big.Int)

	for r.Cmp(big.NewInt(0)) == 0 || s.Cmp(big.NewInt(0)) == 0 {
		k, kInv, err := generateMessageSecret(params)
		if err != nil {
			fmt.Printf("Could not generate message secret: %s\n", err.Error())
			return signature
		}

		digest := sha.Digest(content)

		r.Exp(params.G, k, params.P)
		r.Mod(r, params.Q)

		z.SetBytes(digest)

		xr := new(big.Int)
		sumZxr := new(big.Int)
		kInvZxr := new(big.Int)

		xr.Mul(privKey.Int, r)
		sumZxr.Add(z, xr)
		kInvZxr.Mul(kInv, sumZxr)
		s.Mod(kInvZxr, params.Q)
	}

	signature = Signature{r: r, s: s}
	return signature
}

func generateMessageSecret(params Parameters) (*big.Int, *big.Int, error) {
	k := new(big.Int)
	kInverse := new(big.Int)

	c := new(big.Int)
	valid := false

	for !valid {
		n := 20
		b := make([]byte, n)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Printf("Could not generate random bits: %v\n", err.Error())
			return k, kInverse, err
		}

		c.SetBytes(b)

		qSubTwo := new(big.Int)
		two := big.NewInt(2)
		qSubTwo.Sub(params.Q, two)
		if c.Cmp(qSubTwo) != 1 {
			valid = true
		}
	}
	fmt.Printf("c: %v\n", c)

	// calc k
	one := big.NewInt(1)
	k.Add(c, one)

	// calc modular inverse of k
	val := kInverse.ModInverse(k, params.Q)
	if val == nil {
		return k, kInverse, fmt.Errorf("Error: k is not relatively prime to q\n")
	}
	return k, kInverse, nil
}

func Verify(sig Signature, content []byte, pubKey publicKey, params Parameters) (bool, error) {
	// basic signature validity check
	if sig.r.Cmp(big.NewInt(0)) == 0 || sig.s.Cmp(big.NewInt(0)) == 0 {
		return false, fmt.Errorf("Error: Invalid signature fields")
	}

	sInv := new(big.Int)
	w := new(big.Int)
	z := new(big.Int)
	zw := new(big.Int)
	rw := new(big.Int)
	u1 := new(big.Int)
	u2 := new(big.Int)
	v := new(big.Int)

	sInv.ModInverse(sig.s, params.Q)
	w.Mod(sInv, params.Q)

	digest := sha.Digest(content)
	z.SetBytes(digest)

	zw.Mul(z, w)
	u1.Mod(zw, params.Q)

	rw.Mul(sig.r, w)
	u2.Mod(rw, params.Q)

	v.Exp(params.G, u1, params.P)
	u2.Exp(pubKey.Int, u2, params.P)

	v.Mul(v, u2)
	v.Mod(v, params.P)
	v.Mod(v, params.Q)

	if v.Cmp(sig.r) == 0 {
		return true, nil
	}

	return false, fmt.Errorf("Error: Signature does not match")

}
