package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type person struct {
	Firstname string `json:first_name`
	Lastname  string `json:last_name`
	Age       int    `json:age`
}

func closeOnMarshalFailure() {
	r := recover()
	log.Println(r)
	if r != nil {
		log.Fatalln(r)
	}
}
func main() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	f, err := os.Open("names.txt")

	if err != nil {
		log.Println(err)
	} else {
		defer f.Close()
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(bs))
		}
	}
	p := person{"Gaurav", "Saini", 22}

	value, err := json.Marshal(p)
	log.Println(value, err)
	defer closeOnMarshalFailure()
	if err != nil {
		panic(err)
		return
	} else {
		log.Println(string(value))
	}
}
