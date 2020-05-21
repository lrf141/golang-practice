package main

import "fmt"

func main() {
	s := make([]int, 10, 100)
	fmt.Println(s)
	fmt.Printf("len(%d)\n", len(s))
	fmt.Printf("cap(%d)\n", cap(s))
}
