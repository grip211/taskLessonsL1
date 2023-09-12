package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan struct{}) // создаем канал

	go func(ch chan struct{}) {
		ticker := time.NewTicker(time.Second * 2)
		defer ticker.Stop() // закрываем счетчик

		for {
			select {
			case <-channel:
				fmt.Println("Пока :(")
				return
			case <-ticker.C:
				// что-то делает горутина каждые 5 секунд
				fmt.Println("я делаю что-то")
			}
		}

	}(channel)

	// где-то во вне
	<-time.After(time.Second * 10)
	close(channel) // закрываем канал и в горутину летит zero value для канала тем самым принимает "сигнал"

	<-time.After(time.Second)
}
