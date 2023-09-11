package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	numbers := [5]int{2, 4, 6, 8, 10} //создаем массив из 5 элементов int

	ch := make(chan int, len(numbers)) // создаем буферизированный канал с длинной массива

	var wg sync.WaitGroup

	wg.Add(1)
	go squareOfNumbers(ch, numbers, &wg) // запускаем горутину, считаем квадраты чисел и получаем в канале

	wg.Add(1)
	go sumOfNumbers(ch, &wg) // считаем сумму полученных чисел и получать их в канале

	wg.Wait()
}

// вычисление квадрата чисел массива
func squareOfNumbers(ch chan int, numbers [5]int, wg *sync.WaitGroup) {
	localWg := sync.WaitGroup{} // локально обявляем экземпляр

	for _, v := range numbers { // перебираем элементы массива
		localWg.Add(1)
		go func(v int) {
			s := v * v                      // квадрат числа обявленного в массиве
			fmt.Printf("%d^2 = %d\n", v, s) // выводим результат квадрата каждого числа с массива
			ch <- s                         // отправляем полученные результаты в канал
			localWg.Done()
		}(v) // передаем занчение явно
	}
	localWg.Wait()
	close(ch) // закрываем канал
	wg.Done()
}

// вычисление суммы полученных при квадрате чисел массива
func sumOfNumbers(ch chan int, wg *sync.WaitGroup) {
	result := 0 // объявление переменной для хранения чисел

	for v := range ch { // ждем числа из канала
		result += v // прибавляем к чисалм полученным с канала
	}
	fmt.Printf("Sum of squares of numbers is: %d\n", result) // выводим сумму квадратов чисел массива
	wg.Done()
}
