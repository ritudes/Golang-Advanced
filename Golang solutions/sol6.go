package main

import "fmt"

func main() {
	var a float64 = 67.88
	fmt.Println("After conversion: ", int(a))

	var num1 int = 2
	var num2 float64 = 9.32
	var sum float64
	sum = num2 + float64(num1)
	fmt.Println("Sum =  ", sum)
}
