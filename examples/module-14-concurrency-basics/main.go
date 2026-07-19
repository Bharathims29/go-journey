package main

import (
	"fmt"
	"sync"
	"time"
)

func printTask(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("finished:", name)
}

func sendGreeting(messages chan<- string) {
	messages <- "hello from a channel"
}

func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		results <- fmt.Sprintf("worker %d completed %s", id, job)
	}
}

func slowResult() <-chan string {
	result := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		result <- "result arrived"
	}()

	return result
}

func main() {
	fmt.Println("Module 14: Concurrency Basics")
	fmt.Println()

	fmt.Println("WaitGroup with goroutines")
	var wg sync.WaitGroup
	for _, task := range []string{"load data", "validate input", "save report"} {
		wg.Add(1)
		go printTask(task, &wg)
	}
	wg.Wait()
	fmt.Println()

	fmt.Println("Unbuffered channel")
	messages := make(chan string)
	go sendGreeting(messages)
	fmt.Println(<-messages)
	fmt.Println()

	fmt.Println("Buffered channel and worker")
	jobs := make(chan string, 3)
	results := make(chan string, 3)

	var workerWG sync.WaitGroup
	workerWG.Add(1)
	go worker(1, jobs, results, &workerWG)

	jobs <- "job one"
	jobs <- "job two"
	jobs <- "job three"
	close(jobs)

	workerWG.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}
	fmt.Println()

	fmt.Println("Select with timeout")
	select {
	case result := <-slowResult():
		fmt.Println(result)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("timed out")
	}
}
