package main

import (
	"log"
	"time"
)

// функиция  sleep
func sleep(d time.Duration) {
	<-time.After(d) // ждем в канале
}

func main() {
	log.Println("Сделаем что нибудь?") // вывод
	sleep(5 * time.Second)
	log.Println("продолжим? нет? пока!")
}
