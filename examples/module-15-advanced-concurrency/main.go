package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Job struct {
	ID    int
	Value int
}

type Result struct {
	WorkerID int
	JobID    int
	Value    int
}

type SafeCounter struct {
	mu     sync.Mutex
	counts map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{counts: make(map[string]int)}
}

func (c *SafeCounter) Add(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counts[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.counts[key]
}

func worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, completed *atomic.Int64, stats *SafeCounter, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}

			time.Sleep(20 * time.Millisecond)
			results <- Result{
				WorkerID: id,
				JobID:    job.ID,
				Value:    job.Value * job.Value,
			}
			completed.Add(1)
			stats.Add("completed")
		}
	}
}

func runWorkerPool(ctx context.Context, workerCount int, input []Job) ([]Result, int64) {
	jobs := make(chan Job)
	results := make(chan Result)
	stats := NewSafeCounter()
	var completed atomic.Int64
	var wg sync.WaitGroup

	for id := 1; id <= workerCount; id++ {
		wg.Add(1)
		go worker(ctx, id, jobs, results, &completed, stats, &wg)
	}

	go func() {
		defer close(jobs)

		limiter := time.NewTicker(10 * time.Millisecond)
		defer limiter.Stop()

		for _, job := range input {
			select {
			case <-ctx.Done():
				return
			case <-limiter.C:
				jobs <- job
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var collected []Result
	for result := range results {
		collected = append(collected, result)
	}

	return collected, completed.Load()
}

func waitForSlowWork(ctx context.Context) string {
	done := make(chan string, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		done <- "slow work completed"
	}()

	select {
	case result := <-done:
		return result
	case <-ctx.Done():
		return "slow work canceled"
	}
}

func heartbeat(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("heartbeat stopped")
			return
		case <-ticker.C:
			fmt.Println("heartbeat")
		}
	}
}

func main() {
	fmt.Println("Module 15: Advanced Concurrency")
	fmt.Println()

	fmt.Println("Worker pool with fan-out and fan-in")
	jobs := []Job{
		{ID: 1, Value: 2},
		{ID: 2, Value: 3},
		{ID: 3, Value: 4},
		{ID: 4, Value: 5},
	}

	results, completed := runWorkerPool(context.Background(), 2, jobs)
	for _, result := range results {
		fmt.Printf("worker %d squared job %d = %d\n", result.WorkerID, result.JobID, result.Value)
	}
	fmt.Println("completed jobs:", completed)
	fmt.Println()

	fmt.Println("Context timeout")
	timeoutCtx, cancelTimeout := context.WithTimeout(context.Background(), 80*time.Millisecond)
	defer cancelTimeout()
	fmt.Println(waitForSlowWork(timeoutCtx))
	fmt.Println()

	fmt.Println("Graceful shutdown")
	shutdownCtx, stop := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go heartbeat(shutdownCtx, &wg)

	time.Sleep(120 * time.Millisecond)
	stop()
	wg.Wait()
}
