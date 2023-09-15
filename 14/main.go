package main

import "fmt"

func main() {

	var x interface{} = 10.2
	switch v := x.(type) {
	case int:
		fmt.Println("int :", v)
	case float64:
		fmt.Println("Float64 :", v)
	case string:
		fmt.Println("string :", v)
	case bool:
		fmt.Println("bool :", v)
	default:
		fmt.Println("unknown")

	}
}
