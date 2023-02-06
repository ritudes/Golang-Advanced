package main

import (
	"fmt"
	"sync"
)

type employee struct {
	name string
	age  int
}

func get(s employee, wg *sync.WaitGroup) {
	fmt.Println(s.name, " ", s.age)
	wg.Done()
}

func process(s employee, ch chan employee, wg *sync.WaitGroup) {
	ch <- s
	wg.Done()
}
func main() {
	name := "Ritu"
	ch := make(chan employee, 5)
	var wg sync.WaitGroup
	for i := 24; i <= 33; i = i + 1 {
		s := employee{
			name: name,
			age:  i,
		}
		wg.Add(1)
		go process(s, ch, &wg)

		wg.Add(1)
		go get(<-ch, &wg)
	}

	wg.Wait()


}
