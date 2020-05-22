package main

import "fmt"

func main() {
	a := [10]int{1,2,3,4,5,6,7,8,9,10}
	s1 := a[:]
	s2 := a[2:]
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)

	s1 = s2
	s2[0] = 1
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
}
