package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// var (
	// 	p *big.Int // big prime
	// 	q *big.Int // big prime
	// )
	// p := new(big.Int)
	// p.SetString("58258341155718285019", 10)

	// q := new(big.Int)
	// q.SetString("59867257146064236763", 10)
	// fmt.Println(p)
	// fmt.Println(q)

	p, _ := rand.Prime(rand.Reader, 100)
	q, _ := rand.Prime(rand.Reader, 50)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("q: %v\n", q)

	pp := new(big.Int).Sub(p, big.NewInt(1)) // p-1
	fmt.Printf("pp: %v\n", pp)

	h, _ := rand.Int(rand.Reader, pp) // random number in [0, p-1). TODO (1, p-1)
	fmt.Printf("h: %v\n", h)

	ppq := new(big.Int).Div(pp, q) // (p-1)/q
	fmt.Printf("ppq: %v\n", ppq)

	g := new(big.Int).Exp(h, ppq, p)

	fmt.Printf("g: %v\n", g)

	xa, _ := rand.Int(rand.Reader, big.NewInt(10))
	xb, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Printf("a: %v\n", xa)
	fmt.Printf("b: %v\n", xb)

	ya := new(big.Int).Exp(g, xa, p)
	yb := new(big.Int).Exp(g, xb, p)
	fmt.Printf("ya: %v\n", ya)
	fmt.Printf("yb: %v\n", yb)

	zz1 := new(big.Int).Exp(yb, xa, p)
	zz2 := new(big.Int).Exp(ya, xb, p)
	fmt.Printf("zz1: %v\n", zz1)
	fmt.Printf("zz2: %v\n", zz2)

	if zz1.Cmp(zz2) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}
