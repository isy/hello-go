package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	a := slice[:]
	d := slice[1:2]
	fmt.Println(slice)
	fmt.Println(a)
	fmt.Println(d)
}
