package main

import (
	"fmt"
	"sync"
)

// создаем каналы для чисел и второй — результат операции x*2
// 3 горутины функиций для чисел, квадрата чисел и вывода полученного результата
func main() {
	numsCh := make(chan int)
	generationCh := make(chan int)

	wg := new(sync.WaitGroup)

	wg.Add(3)
	go gettingNumbers(wg, numsCh)
	go generationNumbers(wg, numsCh, generationCh)
	go outputsNumbers(wg, generationCh)

	wg.Wait()

}

// функия для чисел. где обяъвляем и отправляем в канал
func gettingNumbers(wg *sync.WaitGroup, ch chan int) {
	numbers := []int{10, 20, 30, 40, 50, 55}

	for _, n := range numbers {
		ch <- n
	}
	close(ch)
	wg.Done()
}

// функия для возведения чисел в квадрат, и отправки
func generationNumbers(wg *sync.WaitGroup, numsCh <-chan int, squareCh chan int) {
	for n := range numsCh {
		sum := n * 2
		squareCh <- sum
	}
	close(squareCh)
	wg.Done()
}

// функция вывода результата с канала
func outputsNumbers(wg *sync.WaitGroup, ch <-chan int) {
	for n := range ch {
		fmt.Println(n)
	}
	wg.Done()
}
