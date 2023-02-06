package main

import "fmt"

func main() {

	arr := [5]int{45,46,67,88,99}
	add := arr[2]+arr[4]
	fmt.Println("The sum is: ", add)

	i := 0
	j := len(arr) - 1
	slice := arr[i+3:j]
	fmt.Println("Slice: ", slice)
}
