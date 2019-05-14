package main

import (
	"fmt"
	"runtime"
)

var xbool bool
var xint int
var xfloat float64
var xstring string

const a = 10
const (
	g         = iota
	h         = 10 // Untyped constant
	c float64 = 68.876
	d string  = "kh" // Typed constants
	e         = iota
	f         = iota
)

const (
	b  = 1 << (iota * 10)
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Println("Before assignment")
	fmt.Println(xbool)
	fmt.Println(xint)
	fmt.Println(xfloat)
	fmt.Println(xstring)
	xbool = true
	xint = 10
	xfloat = 45.7656565
	xstring = "Something"
	xstring = `Raw string`
	fmt.Println("\nAfter assignment")
	fmt.Println(xbool)
	fmt.Println(xint)
	fmt.Println(xfloat)
	fmt.Println(xstring)
	fmt.Println([]byte(xstring))
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println("\nSizes")
	fmt.Printf("%d\t\t%b\n", b, b)
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
	fmt.Println("\nSystem Info")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
