package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

func main() {
	p := new(Person)
	fmt.Printf("%+v\n", p)

	num := new(int)
	fmt.Println(*num)
	
}
