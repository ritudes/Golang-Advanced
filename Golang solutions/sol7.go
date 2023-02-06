package main

import "fmt"

func main() {

	stumarks := map[string]int{
		"Riya": 40,
		"Nathan": 26,
		"Neil": 34,
		"Chris": 37,
		"McLusky": 50,
	}

	for key,value := range stumarks{
		if value == 34{
			stumarks[key] = value+1
		}

	fmt.Println("Student name: ",key)
	fmt.Println("Marks scored: ",stumarks[key])
	}
}
