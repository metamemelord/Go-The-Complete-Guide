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

	// Using make with slices
	w := make([]int, 10, 100)
	fmt.Println(w)
	fmt.Println(len(w)) // Current length
	fmt.Println(cap(w)) // Length of the underlying array

	w = append(w, 323)

	// Multiple dimensional slices
	v1 := []string{"hehe", "hehe", "has", "dele"}
	fmt.Println(v1)
	v2 := []string{"rinkiya", "ke", "papa"}
	fmt.Println(v2)
	multV := [][]string{v1, v2}
	fmt.Println(multV)

	//Deleting
	z = append(z[:2], z[5:]...)
	fmt.Println(z)

	/*
	 * MAPS
	 */
	m := map[string]int{
		"Papa":    42,
		"Rinkiya": 12,
		"Mommy":   42,
	}
	m["Anotherone"] = 20
	fmt.Println(m)
	fmt.Println(m["Papa"])
	if v, exists := m["SomeKey?"]; exists {
		fmt.Println("MAP CONTAINS THIS VALUE ", v)
	} else {
		fmt.Printf("MAP DOES NOT CONTAIN KEY '%s'\n", "SomeKey?")
	}
	if v, exists := m["Papa"]; exists {
		fmt.Printf("MAP CONTAINS VALUE %d for KEY '%s'\n", v, "Papa")
	}

	fmt.Println("\nRanging over a Map")
	for k, v := range m {
		fmt.Printf("Key: %s => Value: %d\n", k, v)
	}

	delete(m, "RinkiyaKePapa")               // Works even when key doesn't exist
	if v, ok := m["Somekeydoesnexist"]; ok { // Check before deleting
		fmt.Printf("Deleting...")
		delete(m, "Somekeydoesnexist")
	}
}
