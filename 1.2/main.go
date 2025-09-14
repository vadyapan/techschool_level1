package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var nums [5]int = [5]int{2, 4, 6, 8, 10}

	for _, num := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			squareOfNumbers(num)
		}()
	}

	wg.Wait()
}

func squareOfNumbers(num int) {
	res := num * num
	fmt.Println(res)
}

/*
// Go 1.25
func main() {
	var wg sync.WaitGroup
	var nums [5]int = [5]int{2, 4, 6, 8, 10}

	for _, num := range nums {
		wg.Go(func() {
			squareOfNumbers(num)
		})
	}

	wg.Wait()
}
*/
