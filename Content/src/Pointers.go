package main

import (
	"fmt"
)

func foo(a *int) {
	*a++
}

func main() {
	a := 42
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*&a)
}
