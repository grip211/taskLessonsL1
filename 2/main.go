package main

import (
	"fmt"
	"sync"
)

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
	arr := [5]int{2, 4, 6, 8, 10}
	c := make(chan int, 5)

	var wg sync.WaitGroup

	for _, v := range arr {
		wg.Add(1)
		go square(c, &wg, v)
	}
	wg.Wait()
	close(c)

	for i := 0; i < cap(c); i++ {
		fmt.Println(<-c)
	}
}

func square(c chan int, wg *sync.WaitGroup, v int) {
	c <- v * v
	wg.Done()
}
