package main

import (
	"fmt"
	"sync"
)

func main() {

	c := 0
	n := 200
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			m.Lock()
			c++
			m.Unlock()
			wg.Done()

		}(i)
	}
	wg.Wait()

	fmt.Println(c)
}
