package main

import "fmt"

func main() {
	a := [3]int{1,2,3}
	b := [3]int{4,5,6}

	a = b
	fmt.Println(a)
	a[0] = 1
	fmt.Println(b)
}
