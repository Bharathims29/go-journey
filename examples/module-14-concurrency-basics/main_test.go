package main

import (
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	jobs := make(chan string, 2)
	results := make(chan string, 2)

	var wg sync.WaitGroup
	wg.Add(1)
	go worker(1, jobs, results, &wg)

	jobs <- "job one"
	jobs <- "job two"
	close(jobs)

	wg.Wait()
	close(results)

	got := len(results)
	want := 2
	if got != want {
		t.Fatalf("worker produced %d results, want %d", got, want)
	}
}
