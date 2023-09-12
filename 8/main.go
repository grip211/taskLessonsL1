package main

import (
	"fmt"
)

func main() {
	var i int64 = 10
	fmt.Println("Before change:", i)

	// установим 24ый бит в 1
	i |= 1 << 2
	fmt.Println("After first change:", i)

	// установим 16ый бит в 0
	i &= ^(1 << 2)

	fmt.Println("After second change:", i)
}

/*
Поразрядные операции примеры
func main() {
	var a int = 5 | 2  // 101 | 010 = 111  - 7
	var b int = 6 & 2  // 110 & 010 = 10  - 2
	var c int = 5 ^ 2  // 101 ^ 010 = 111 - 7
	var d int = 5 &^ 6 // 101 &^ 110 = 001 - 1
}

*/
