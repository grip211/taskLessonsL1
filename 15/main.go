package main

import (
	"fmt"
	"strings"
)

func someFunc() string {
	return createHugeString(1 << 10)[:100] // в данном случае если символы будут занимать вместо 1 байта - 2 (к примеру русские буквы), будем получать половину строки (50 символов вместо 100, так как 1 символ теперь будет занимать два байта)

}

func main() {

	fmt.Println(someFunc()) // 50 символов

	fmt.Println(someFuncFix()) // 100 cимволов как и нужно
}

func createHugeString(n int) string {
	str := strings.Builder{}

	for i := 0; i < n; i++ {
		str.WriteString("а")
	}

	return str.String()

}

func someFuncFix() string {
	v := []rune(createHugeString(1 << 10))
	return string(v[:100])
}
