package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int, 10)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for job := range jobs {
			fmt.Fprintln(os.Stdout, job)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	timeout := time.After(2 * time.Second)

loop:
	for {
		select {
		case jobs <- rand.Intn(100):
			time.Sleep(200 * time.Millisecond)
		case <-timeout:
			fmt.Println("\nCompletion signal received, closing channel...")
			close(jobs)
			break loop
		}
	}

	wg.Wait()
	fmt.Println("All workers have completed their tasks. The programme has been stopped.")
}
