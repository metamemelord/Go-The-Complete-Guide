package main

import "fmt"

type person struct {
	first string
	last  string
}

type occupation struct {
	person // No need to expand this, unqualified type name acts as field
	// person person // can also be used
	first string // same as the value in person struct
	title string
}

/********** Test Code **********/
type list []int

func (s *list) append(value int) *list {
	*s = append(*s, value)
	return s
}

func testAppend() {
	x := list{10, 20}
	x.append(20).append(30)
	fmt.Println(x)
}

/*******************************/

func main() {
	testAppend()
	p1 := person{ // Value of type Person
		first: "Rinkiya ke",
		last:  "Papa",
	}

	o1 := occupation{p1, "lel", "Foftware Bngineer"}

	fmt.Println(p1, o1)
	// Elements of child are bound to the parent automatically (Also called Embedded struct)
	fmt.Println(o1.person, o1.title, o1.first, o1.person.first, o1.person.last, o1.last) // person.first person.last are promoted fields

	// Anynomous struct
	p2 := struct {
		first string
		last  string
	}{
		first: "Hehe",
		last:  "Has dele",
	}

	fmt.Println(p2)
}
