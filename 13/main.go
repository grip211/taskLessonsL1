package main

import "fmt"

func main() {

	a := 4
	b := 5

	fmt.Println("До замены :", a, b) // здесь до замены
	a, b = b, a
	fmt.Println("После замены местами :", a, b) // после замены местами
}
