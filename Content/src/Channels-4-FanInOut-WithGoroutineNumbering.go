package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fin := make(chan string)
	fout := make(chan string)
	m := make(map[string]int)
	go func() {
		for i := 0; i < 500; i++ {
			fin <- "TEST"
		}
		close(fin)
	}()
	go fanInOut(fin, fout)
	for value := range fout {
		m[value]++
		//fmt.Println(value)
	}
	fmt.Println(m)
}

func fanInOut(fin <-chan string, fout chan<- string) {
	var wg sync.WaitGroup
	const goRoutines = 5
	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func(l int) {
			for v := range fin {
				fout <- timeConsuming(v, l)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(fout)
}

func timeConsuming(value string, i int) string {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(1000)))
	return fmt.Sprintf("%s-%d", value, i)
}
