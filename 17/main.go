package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{1, 4, 50, 123, 124, 129, 130, 200} // отсортированный массив слайса чисел

	// выбираем нужное число которое ищем
	x := 129

	// с помощью общей универсальной функцией бинарного поиска sort.Search ищем нужное число
	//наименьший индекс i, при котором f(i) истинно (равно true)
	//или n, если такого индекса нет.

	i := sort.Search(len(a), func(i int) bool { return x <= a[i] })
	if i < len(a) && a[i] == x {
		fmt.Printf("Найдено %d по индексу %d в %v.\n", x, i, a)
	} else {
		fmt.Printf("Не найдено %d в %v.\n", x, a)
	}
}
