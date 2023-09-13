package main

import "fmt"

func main() {

	slice1 := []int{6, 2, 8, 4, 5} //создаем слайс1 с числами
	slice2 := []int{6, 8, 1, 3, 2} // создаем слайс 2 с числами

	fmt.Println(intersecting(slice1, slice2)) // вывводим результат функиции  пересечение в  двух слайсах чисел
}

// функия перебора слайсов с числами где пересекаются числа
func intersecting(slice1 []int, slice2 []int) []int {
	var slice []int // обявляем пустой слайс

	for _, i := range slice1 { // перебираем первый слайс
		for _, i2 := range slice2 { // перебираем второй слайс
			if i == i2 { // если имеются совпадения
				slice = append(slice, i) // то добаляем в слайс
			}
		}
	}
	return slice // вовзращаем слайс
}
