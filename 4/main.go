package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	n := numberWorkers() //объявляем количество воркеров

	dataCh := make(chan interface{}, n) // создаем канал данных
	breakCh := breakCommands()          // функция остановки команды

	for i := n; i > 0; i-- { // запускаем горутину печати данных
		go printData(dataCh)
	}

PutData: // создаем цикл
	for i := 0; ; i++ {
		select {
		case <-breakCh: // проверяем канал на вход данных
			close(dataCh)  // если есть то закрываем канал
			close(breakCh) // и канал остановки команды
			break PutData  // выходим из цикла
		default: // если ничего не пришло, вводим число
			dataCh <- i // отправляем в канал
		}
	}
}

// функиция количества воркеров
func numberWorkers() int {
	fmt.Println("Number of workers: ") // выводим содержимое
	n := 0                             // объявляем переменную
	_, err := fmt.Scanln(&n)           //	если ведем не число или ничего то паника
	if err != nil {
		panic(err)
	}

	return n // возвращаем число
}

// функиция для завершения
func breakCommands() chan os.Signal {
	ch := make(chan os.Signal, 1) // объявляем канал который завершается по нажатию Ctrl+C
	signal.Notify(ch, os.Interrupt)

	return ch // возвращаем канал
}

// функия принятия и вывода данных
func printData(ch <-chan interface{}) {
	for d := range ch { // принимаем данные через канал
		fmt.Println(d) // выводим данные
	}
}
