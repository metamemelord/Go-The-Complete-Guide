package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	o := make(chan int)
	e := make(chan int)
	fin := make(chan int)

	go send(e, o)

	go receive(e, o, fin)

	for v := range fin {
		fmt.Println(v)
	}

	// Fanout
	fout := make(chan int)
	c := make(chan int)
	go populate(c)
	go fanOut(c, fout)
	for v := range fout {
		fmt.Println(v)
	}
	fmt.Println("Exiting...")
}

func send(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			o <- i
		} else {
			e <- i
		}
	}
	close(e)
	close(o)
}

func receive(e, o <-chan int, fin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fin)
}

func populate(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}

func fanOut(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(val int) {
			c2 <- timeConsuming(val)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsuming(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
