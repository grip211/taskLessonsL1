package main

import "fmt"

type Human struct {
	Action
	name    string
	surname string
	age     int
}

type Action struct {
	job string
}

func main() {
	var tom = Human{
		name:    "Tom",
		surname: "Sivers",
		age:     24,
		Action: Action{
			job: "Doctor",
		},
	}
	fmt.Println(tom.name)
	fmt.Println(tom.age)
	fmt.Println(tom.job)
}
