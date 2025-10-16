package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	jobs := make(chan int, 10)
	wg := sync.WaitGroup{}

	workerCount := flag.Int("w", 0, "worker count")
	flag.Parse()

	if *workerCount == 0 {
		panic("Please enter -w and number worker count")
	}

	for i := 0; i < *workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for job := range jobs {
				fmt.Fprintln(os.Stdout, job)
				time.Sleep(300 * time.Millisecond)
			}
		}()
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

loop:
	for {
		select {
		case jobs <- rand.Intn(100):
			time.Sleep(200 * time.Millisecond)
		case <-stop:
			fmt.Println("\nCompletion signal received, closing channel...")
			close(jobs)
			break loop
		}
	}

	wg.Wait()
	fmt.Println("All workers have completed their tasks. The programme has been stopped.")
}
