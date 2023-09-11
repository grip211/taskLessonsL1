package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает
//значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

/*
func main() {
	arr := [5]int{2, 4, 6, 8, 10} // создаем массив из 5 элементов int
	var wg sync.WaitGroup

	for _, v := range arr {
		go square(v, &wg)
		wg.Add(1) // в группе дна горутина
	}
	wg.Wait() // ожидаем завершение всех горутин wg

}

func square(v int, wg *sync.WaitGroup) {
	fmt.Println(v * v)
	wg.Done() //уменьшает внутренний счетчик активных элементов на единицу (вызываем горутину)
}
*/

func main() {
	arr := [5]int{2, 4, 6, 8, 10} // создаем массив из 5 элементов int
	c := make(chan int, 5)        // создаем буферизированный канал с емкостью 5

	var wg sync.WaitGroup

	for _, v := range arr { // перебираем элементы массива
		wg.Add(1)
		go square(c, &wg, v)
	}
	wg.Wait()
	close(c) // закрываем канал

	for i := 0; i < cap(c); i++ {
		fmt.Println(<-c) // получаем данные с канала
	}
}

func square(c chan int, wg *sync.WaitGroup, v int) {
	c <- v * v // отправляем данные в канал
	wg.Done()
}
