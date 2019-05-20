package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Printf("%T", s) // This should be a pointer to execute s.area() as area() has a pointer receiver
	fmt.Println("Area:", s.area())
}
func foo(i string) {
	fmt.Println(i)
	wg.Done()
}

func main() {
	// WaitGroups
	wg.Add(2)
	fmt.Println(runtime.NumCPU())
	go foo("Hey man!")
	// wg.Add(1)
	go foo("Oh ooy, does this work?")
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()

	//Method sets
	c := circle{3.343}
	fmt.Println("\n----------Method sets----------")
	fmt.Println(c.area())
	// info(c) // ERROR: This type does not implement the shape interface
	info(&c)

	// Incrementer program with Race Condition
	counter := 0
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			val := counter
			runtime.Gosched()
			val++
			counter = val
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("\n----------Race Condition----------")
	fmt.Println(counter)

	// Incrementer program with Mutex to fix Race Condition
	counter = 0
	var mutex sync.Mutex
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			mutex.Lock()
			val := counter
			runtime.Gosched()
			val++
			counter = val
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("\n----------Race Condition----------")
	fmt.Println(counter)

	var incrementer int64
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			val := incrementer
			runtime.Gosched()
			val++
			incrementer = val
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("\n----------Atomic----------")
	fmt.Println(counter)

	fmt.Println("\n----------Arch and OS----------")
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
}
