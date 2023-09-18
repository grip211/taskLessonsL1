package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "asdf"     // true
	str2 := "asdfF"    // false
	str3 := "asdfeNgK" // false
	str4 := "aAdf"     // false

	fmt.Printf("str1: %s, res: %v\n", str1, uniqueStrings(str1))
	fmt.Printf("str2: %s, res: %v\n", str2, uniqueStrings(str2))
	fmt.Printf("str3: %s, res: %v\n", str3, uniqueStrings(str3))
	fmt.Printf("str4: %s, res: %v\n", str4, uniqueStrings(str4))
}

func uniqueStrings(str string) bool {
	set := make(map[rune]struct{}) // храним уникальные символы в мапе

	lowString := strings.ToLower(str) // приводим строку к меньшому регистру

	for _, value := range lowString {
		if _, exists := set[value]; exists {
			return false
		} else {
			set[value] = struct{}{}
		}
	}
	return true
}
