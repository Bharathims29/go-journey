# Module 15: Advanced Concurrency

## Goal

Learn how to control concurrent programs with cancellation, timeouts, worker pools, fan-in/fan-out, mutexes, atomics, rate limiting, and graceful shutdown.

By the end of this module, you should understand:
- Context cancellation
- Timeouts
- Worker pools
- Fan-in and fan-out patterns
- Mutexes
- Atomic operations basics
- Rate limiting
- Graceful shutdown

## 1. From Starting Goroutines to Controlling Them

Module 14 taught you how to start goroutines and communicate with channels.

Advanced concurrency is about control:
- How does work stop?
- How long should work be allowed to run?
- How many goroutines should run at once?
- How is shared data protected?
- How does the program shut down cleanly?

Beginner rule:

Every goroutine should have a clear reason to stop.

## 2. Context Cancellation

The `context` package carries cancellation signals and deadlines across function calls.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

A goroutine can listen for cancellation:

```go
select {
case <-ctx.Done():
	return
case job := <-jobs:
	fmt.Println(job)
}
```

Use context when work may need to stop early.

## 3. Timeouts

A timeout is a deadline for work.

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
```

If the work takes too long, `<-ctx.Done()` becomes ready.

Timeouts are useful for:
- HTTP requests
- Database calls
- Background jobs
- Waiting for results from goroutines

## 4. Worker Pools

A worker pool runs a fixed number of workers that read jobs from a channel.

This controls how much work runs at once.

```go
jobs := make(chan Job)
results := make(chan Result)

for id := 1; id <= 3; id++ {
	go worker(id, jobs, results)
}
```

Worker pools are useful when you have many jobs but want limited concurrency.

## 5. Fan-Out and Fan-In

Fan-out means sending work to multiple goroutines.

Fan-in means collecting results back into one channel.

Typical shape:
- One input channel
- Many workers
- One results channel

This pattern is common in file processing, API checks, and background queues.

## 6. Mutexes

A mutex protects shared data.

```go
var mu sync.Mutex
count := 0

mu.Lock()
count++
mu.Unlock()
```

Use a mutex when multiple goroutines need to read and write the same variable or map.

Common rule:

Keep locked sections small.

## 7. Atomic Operations

The `sync/atomic` package provides low-level safe operations for simple shared values.

```go
var count atomic.Int64
count.Add(1)
fmt.Println(count.Load())
```

Atomics are good for simple counters and flags.

For grouped state or maps, use a mutex.

## 8. Rate Limiting

Rate limiting controls how often work happens.

```go
limiter := time.Tick(200 * time.Millisecond)

for job := range jobs {
	<-limiter
	process(job)
}
```

This is useful when calling APIs, writing logs, or protecting a service from too much work at once.

## 9. Graceful Shutdown

Graceful shutdown means stopping new work, finishing current work, and then exiting.

The basic idea:
- Listen for cancellation
- Stop accepting new work
- Close channels when senders are done
- Wait for goroutines with `sync.WaitGroup`

Do not leave background goroutines running forever.

## 10. Run the Example

Open this example folder:

[module-15-advanced-concurrency](../examples/module-15-advanced-concurrency)

From the module example folder, run:

```bash
cd examples/module-15-advanced-concurrency
go run .
```

Run the tests:

```bash
go test -v
```

Run the race detector:

```bash
go test -race
```

The example includes:
- A worker pool
- Fan-out and fan-in
- Context timeout
- Rate limiting
- Mutex-protected shared state
- Atomic counter
- Graceful shutdown with context cancellation

## 11. Practice Tasks

1. Build a worker pool with three workers and ten jobs.
2. Add `context.WithTimeout` so slow work stops early.
3. Fan out numbers to workers and fan in squared results.
4. Protect a shared map with `sync.Mutex`.
5. Count completed jobs with `atomic.Int64`.
6. Add a simple rate limiter using `time.Tick`.
7. Run your program with `go test -race`.

## Checkpoint

You are ready for the next module when:
- You can cancel goroutines with `context`.
- You can add timeouts to concurrent work.
- You can build a worker pool.
- You understand fan-out and fan-in.
- You can protect shared data with a mutex.
- You can use atomics for simple counters.
- You can explain how a goroutine should stop.
