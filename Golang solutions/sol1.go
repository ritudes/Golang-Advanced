package main

import "fmt"

func main() {
	var f float64 
	f = 4.5

	fmt.Println("Initial value of f: ", f)

	fmt.Println("Address of f is: ", &f)

	a := &f
	fmt.Println("Value of f by using address: ", *a)

	*a = 7.669
	fmt.Println("Updated value of f: ", f)
}
