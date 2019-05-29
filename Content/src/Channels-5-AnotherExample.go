package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	c2 := make(chan int)
	go func() {
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(v int) {
				for j := 0; j < 10; j++ {
					c <- (v + j)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()
	go recieve(c, c2)
	for value := range c2 {
		fmt.Println(value)
	}
}

func recieve(c, c2 chan int) {
	for i := range c {
		c2 <- i
	}
	close(c2)
}
