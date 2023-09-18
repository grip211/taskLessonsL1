package main

import "fmt"

// Код копирует один элемент и выполняется за постоянное время.
func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Удалить элемент по индексу i из a.

	// Выполняем сдвиг a[i+1:] влево на один индекс.
	copy(a[i:], a[i+1:])

	// Удаляем последний элемент (записать нулевое значение).
	a[len(a)-1] = ""

	// Отсекаем срез.
	a = a[:len(a)-1]

	// вывод
	fmt.Println(a) // [A B D E]
}