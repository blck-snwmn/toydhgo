package main

import (
	"fmt"

	"github.com/blck-snwmn/toydhgo"
)

func main() {
	p, _, g := toydhgo.GenerateParam()
	alicesKey := toydhgo.New(p, g)
	bobsKey := toydhgo.New(p, g)

	alicesSKey := alicesKey.GenerateSharedKey(bobsKey.PublicKey())
	bobsSKey := bobsKey.GenerateSharedKey(alicesKey.PublicKey())

	fmt.Println("alice's shared key:", alicesSKey)
	fmt.Println("bob's shared key:", bobsSKey)

	if alicesSKey.Cmp(bobsSKey) != 0 {
		panic("shared key is not equal")
	}
	fmt.Println("shared key is equal")
}
