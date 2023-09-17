package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "snow dog sun" // присваиваем строку переменной

	p := strings.Split(s, " ") // делим строку на слова

	// меняем местами слова путем перебора слов
	for i, j := 0, len(p)-1; i < len(p)/2; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
	// объединяем слова в строку
	s = strings.Join(p, " ")
	// распечатываем
	fmt.Println(s)
}
