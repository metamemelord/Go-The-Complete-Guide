package main

import (
	"fmt"
)

// Reference for composition https://www.ardanlabs.com/blog/2015/09/composition-with-go.html

// func (receiver) identifier(parameters) (returns) { code }

// Everything in Go is PASS BY VALUE

func incrementInt(val *int) {
	*val++
}

// Variadic parameter
func foo(x ...int) { // Can accept 0 parameters
	fmt.Println(x)
	fmt.Printf("%T\n\n", x)
}

func bar(str string, x ...int) { // Can accept 0 parameters
	fmt.Println(str, x)
}

// For defer
func baz() {
	fmt.Println("Baz")
}

func foobar() {
	fmt.Println("Foobar")
}

// Methods
type person struct {
	first string
	last  string
}

type professional struct {
	person
	designation string
}

func (p person) work() {
	fmt.Println(p.first, "is working!")
}

func (p professional) work() {
	fmt.Println(p.first, "is working!")
}

type employee interface {
	work()
}

func lmao(e employee) {
	switch e.(type) {
	case person:
		// Assertion
		fmt.Println("Well, person also works! -", e.(person).first)
	case professional:
		fmt.Println("Professional works, we know! -", e.(professional).first)
	}
	e.work()
}

// Global function type
var k func() int = func() int {
	fmt.Println("hehe")
	return 10
}

// Functions returning functions
func barReturnsFunction() func() int {
	fmt.Println("\nI am inside bar")
	return func() int {
		return 10
	}
}

// Callbacks
func callbackFoo(t func(int) int) {
	fmt.Println("I am inside of callbackFoo")
	fmt.Println(t(100))
}

// Closure
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	v := 10
	fmt.Println(v)
	incrementInt(&v)
	fmt.Println(v)

	foo(1, 2, 3)
	foo()

	x := []int{1, 2, 3}
	foo(x...)
	bar("hehe")
	bar("hehe", 1, 2, 3)
	fmt.Println()
	fmt.Println()

	// Defer
	defer baz() // This like executes after main() is exited
	foobar()

	// Methods and interfaces
	var p employee = professional{
		person{"Gaurav", "Saini"},
		"Software Engineer",
	}
	p2 := person{"Gaurav", "Saini"}
	lmao(p)
	lmao(p2)

	// Anonymous function
	func() {
		fmt.Println("Anon")
	}()

	func(i int) {
		fmt.Println(i)
	}(10)

	// Function type
	f := func() {
		fmt.Println("Hehe has dele")
	}
	f()
	g := func(n int) int {
		return n * n
	}
	fmt.Println(g(10))
	fmt.Println(k())

	// Functions returning functions
	fmt.Println(barReturnsFunction()())

	// Callback
	callbackFoo(func(n int) int {
		fmt.Println("I am inside of callback function passed nigguh!")
		return n * n
	})

	secondCallback := func(n int) int {
		fmt.Println("I am inside the second callback function")
		return n - 10
	}
	callbackFoo(secondCallback)
	fmt.Println("\nClosure:")
	// Closure
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())

	fmt.Println("\nDefer was used for below function:")
}
