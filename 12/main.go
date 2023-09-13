package main

import "fmt"

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"} //объявляем слайс строк со значениями
	many(array)
}

// функиция множества слов(строк) с слайса
func many(array []string) {
	set := make(map[string]struct{}) // создаем мапу

	for _, value := range array { // перебираем слайс
		set[value] = struct{}{}
	}

	for value := range set {
		fmt.Println(value) // выводим результат
	}
}
