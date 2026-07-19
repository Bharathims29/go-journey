package main

import (
	"context"
	"testing"
	"time"
)

func TestRunWorkerPool(t *testing.T) {
	jobs := []Job{
		{ID: 1, Value: 2},
		{ID: 2, Value: 3},
		{ID: 3, Value: 4},
	}

	results, completed := runWorkerPool(context.Background(), 2, jobs)

	if completed != int64(len(jobs)) {
		t.Fatalf("completed = %d, want %d", completed, len(jobs))
	}

	if len(results) != len(jobs) {
		t.Fatalf("got %d results, want %d", len(results), len(jobs))
	}
}

func TestWaitForSlowWorkTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	got := waitForSlowWork(ctx)
	want := "slow work canceled"

	if got != want {
		t.Fatalf("waitForSlowWork() = %q, want %q", got, want)
	}
}

func TestSafeCounter(t *testing.T) {
	counter := NewSafeCounter()

	counter.Add("completed")
	counter.Add("completed")

	got := counter.Value("completed")
	want := 2

	if got != want {
		t.Fatalf("counter value = %d, want %d", got, want)
	}
}
