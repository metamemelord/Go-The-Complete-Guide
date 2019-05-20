package main

import "fmt"

type lmfao int

func main() {
	fmt.Println("Wew I am working fine, fren")
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(n, "lol")
	v := []int{}
	for i := 0; i < n; i++ {
		var val int
		fmt.Scan(&val)
		v = append(v, val)
	}
	for i := 0; i < len(v); i++ {
		fmt.Println(v[i])
	}
	fmt.Println(v)
	var b lmfao = 10
	fmt.Printf("%T %T", b, int(b))
}
