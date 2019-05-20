package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type test struct {
	T1 int    `json:"t1"`
	T2 string `json:"t2"`
}

type person struct {
	name string
	age  int
}

type ByName []person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].name < bn[j].name }

func main() {
	t := test{10, "hehe"}
	fmt.Println(t)
	str, _ := json.MarshalIndent(t, "", "\t")
	fmt.Printf(string(str))
	var unmarshalled test
	err := json.Unmarshal(str, &unmarshalled)
	if err != nil {
		fmt.Println("Sed")
	}
	fmt.Println(unmarshalled)

	fmt.Println("Hello, World!")
	fmt.Fprintln(os.Stdout, "Hello, World!")
	io.WriteString(os.Stdout, "Hello, World!\n")

	xi := []int{2, 5, 7, 3, 2, 5, 3, 1, 323, 43, 234, 454, 656}
	xs := []string{"ajkshd", "hf", "lmfsa", "oh", "kek", "lolp", "lolaxmaoz"}
	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)

	// Custom sorting
	var people []person
	for idx := range xs {
		people = append(people, person{xs[len(xs)-idx-1], xi[idx]})
	}
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

	// Bcrypt
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `password1234`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}
	fmt.Println("You're logged in")
}
