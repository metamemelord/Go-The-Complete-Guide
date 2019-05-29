package main

import (
	"fmt"
	"time"
)

var c chan int

func preCalc() {
	c = make(chan int)
}

func main() {
	// Channels block, hence the below statement doesn't work
	preCalc()
	go func() {
		fmt.Println("Inner goroutine has started, but blocked by channel!")
		fmt.Println("Unblocking this goroutine in..")
		for i := 0; i < 3; i++ {
			fmt.Printf(" %d", 3-i)
			time.Sleep(time.Second)
		}
		c <- 42
		fmt.Println("\nInner goroutine done!")
	}()
	fmt.Println("\n", <-c)

	// Buffered channels
	ch_buf := make(chan int, 1)
	// Line below works irrespective if there's a receiver or now.
	ch_buf <- 10
	// ch_buf <- 10 // This line won't work cus there's space for 1

	// Directional Channels
	cr := make(<-chan int)
	cs := make(chan<- int)

	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
	fmt.Printf("c\t%T\n", (<-chan int)(c)) // General to specific conversion

	go foo(c)

	bar(c)

	// Range on channels
	go func() {
		for i := 0; i < 3; i++ {
			c <- i + 1
		}
		close(c)
	}()
	fmt.Println("These values are coming from range: ")
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Range done!")

	// Select statement for channels (Waiting on multiple channels)
	o := make(chan int)
	e := make(chan int)
	q := make(chan int)

	go send(e, o, q)

	receive(e, o, q)

	// Comma-ok Idiom operates return Zero values IF THE CHANNEL IS CLOSED
	c = make(chan int)
	close(c)
	v, ok := <-c
	fmt.Println(v, ok)

	fmt.Println("Quiting the program now...")
}

// Putting something on a channel inside goroutine pushes that goroutine on sleep and hence
// BLOCKS the execution of that goroutine (this causing a deadlock if all goroutines are BLOCKED), UNLESS there is
// a buffer defined on the channel. If the buffer is full, the goroutine blocks.

// send
func foo(c chan<- int) {
	c <- 2341
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case i, ok := <-e:
			fmt.Println("Even value", i, ok)
		case i, ok := <-o:
			fmt.Println("Odd value", i, ok)
		case i, ok := <-q:
			fmt.Println("The odd-even channel has ended", i, ok)
			return
		default:
			fmt.Println("Sed, nothing to capture")
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			o <- i
		} else {
			e <- i
		}
	}
	close(e)
	q <- -1
}
