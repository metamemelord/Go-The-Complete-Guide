package main

import (
	"fmt"
	"time"
)

func main() {
	// Type 1
	// for init; condition; update {
	// 	statement
	// }
	for i := 0; i <= 100; i++ {
		i += 15
		fmt.Println(i)
	}

	// Type 2 - Similar while loop
	// for condition {
	// 	statement
	// }
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Type 3 - Infinite Loop
	// for {
	// 	statement
	// }
	for {
		fmt.Println(10)
		break
	}

	// Type 4 - Range Loop
	// for idx, val := range collection {
	// 	statement
	// }
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Type 5 - Range Loop with map
	// for key, val := range map {
	// 	statement
	// }
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Type 6 - Range Loop with map
	// for key := range map {
	// 	statement
	// }
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Type 7 - Range Loop with string
	// for idx, UTF := range map {
	// 	statement
	// }
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	// if - else if - else
	// if condition {

	// } else if condition-2 {

	// } else {

	// }

	// switch statements
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//**********
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//**********
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//**********
	whatAmI := func(i interface{}) {
		switch t := i.(type) { // Assign value and perform switch
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
