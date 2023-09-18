package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(231231231574546434)
	b := big.NewInt(231231231574542221)

	t := new(big.Int)
	t.Add(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, t)
	t.Sub(a, b)
	fmt.Printf("%v - %v = %v\n", a, b, t)
	t.Div(a, b)
	fmt.Printf("%v / %v = %v\n", a, b, t)
	t.Mul(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, t)
}
