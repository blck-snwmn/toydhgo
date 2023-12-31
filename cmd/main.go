package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	p, _, g := gen()

	alice := newDH(p, g)
	bob := newDH(p, g)
	alicesKey := alice.getSharedKey(bob.publicKey)
	bobsKey := bob.getSharedKey(alice.publicKey)

	fmt.Printf("alice's key: %v\n", alicesKey)
	fmt.Printf("bob's key: %v\n", bobsKey)

	if alicesKey.Cmp(bobsKey) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}

func gen() (p *big.Int, q *big.Int, g *big.Int) {
	p, _ = rand.Prime(rand.Reader, 100)
	q, _ = rand.Prime(rand.Reader, 50)

	pp := new(big.Int).Sub(p, big.NewInt(1)) // p-1

	h, _ := rand.Int(rand.Reader, pp) // random number in [0, p-1). TODO (1, p-1)

	ppq := new(big.Int).Div(pp, q) // (p-1)/q

	g = new(big.Int).Exp(h, ppq, p)
	return
}

func newDH(p, g *big.Int) *dh {
	x, _ := rand.Int(rand.Reader, big.NewInt(10))
	y := new(big.Int).Exp(g, x, p)
	return &dh{
		prime:      p,
		privateKey: x,
		publicKey:  y,
	}
}

type dh struct {
	prime      *big.Int
	privateKey *big.Int
	publicKey  *big.Int
}

func (d *dh) getSharedKey(otherPubicKey *big.Int) *big.Int {
	return new(big.Int).Exp(otherPubicKey, d.privateKey, d.prime)
}
