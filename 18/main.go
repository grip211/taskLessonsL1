package main

import (
	"fmt"
	"log"
	"sync"
)

type Counter struct {
	num int
	sync.Mutex
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()

	c.num += 1
}

func (c *Counter) Value() int {
	return c.num
}

func main() {
	cnt := &Counter{
		num: 0,
	}

	finish := make(chan struct{})

	go Do(cnt, finish)

	select {
	case <-finish:
		log.Printf("Work done with count: %d", cnt.Value())
	}
}

func Do(cnt *Counter, finish chan struct{}) {
	wg := sync.WaitGroup{}

	// Hard work in goroutine and finish work
	for i := 0; i < 20; i++ {
		wg.Add(1)

		// Тяжелая работа
		go func(num int, cnt *Counter, wg *sync.WaitGroup) {
			defer wg.Done()

			fmt.Printf("Worker %d starting\n", num)
			cnt.Inc()
			fmt.Printf("Worker %d done\n", num)
		}(i, cnt, &wg)
	}

	wg.Wait()
	finish <- struct{}{}
	close(finish)
}
