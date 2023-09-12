package main

import (
	"context"
	"fmt"
	"time"
)

//-завязать все горутины на переданный в них context.

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go someFunc(ctx)

	// где-то во вне
	<-time.After(time.Second * 10)
	cancel() // отменяем контекст и закрываем ctx.Done

	<-time.After(time.Second)
}

func someFunc(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop() // закрываем счетчик

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Пока :(")
			return
		case <-ticker.C:
			// что-то делает горутина каждые 5 секунд
			fmt.Println("я делаю что-то")
		}
	}

}
