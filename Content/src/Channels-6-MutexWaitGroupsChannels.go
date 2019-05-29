package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)
	var mutex sync.Mutex
	go func() {
		var wg sync.WaitGroup
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func(v int) {
				for j := 0; j < 5; j++ {
					c <- fmt.Sprintf("%d-TEST", v)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
		mutex.Unlock()
	}()
	fmt.Println("Lock has been acquired, pushing values onto the channel...")
	mutex.Lock()
	fmt.Println("Lock has been released, printing values:")
	for val := range c {
		fmt.Println(val)
	}
}
