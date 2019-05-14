package main

import "fmt"

/*
append() is defined in builtin
*/
func main() {
	var x [5]int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y []int
	y = append(y, 10)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	z := []int{2, 4, 5, 10, 67} // Composite literal
	z = append(z, 10)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(z[1:])
	z = append(y, z...)
	fmt.Println(z)

	//Deleting
	z = append(z[:2], z[5:]...)
	fmt.Println(z)
}
