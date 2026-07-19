# Module 14: Concurrency Basics

## Goal

Learn how Go runs work concurrently using goroutines, channels, `select`, and `sync.WaitGroup`.

By the end of this module, you should understand:
- Goroutines
- Channels
- Buffered and unbuffered channels
- `select`
- `sync.WaitGroup`
- Race conditions
- Running tests with `-race`

## 1. What Is Concurrency?

Concurrency means a program can make progress on multiple tasks during the same time period.

In Go, concurrency is usually built with:
- Goroutines for running work
- Channels for sending values between goroutines
- Synchronization tools like `sync.WaitGroup`

Concurrency is useful for tasks like:
- Running slow work in the background
- Waiting for multiple operations
- Processing jobs
- Handling many requests

## 2. Goroutines

A goroutine is a lightweight concurrent function.

Start one with the `go` keyword:

```go
go printMessage("learning Go")
```

Example:

```go
func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	go printMessage("hello from a goroutine")
	fmt.Println("hello from main")
}
```

Important beginner rule:

The program exits when `main` exits. If `main` finishes too early, goroutines may not get time to run.

## 3. WaitGroup

Use `sync.WaitGroup` when you need to wait for goroutines to finish.

```go
var wg sync.WaitGroup

wg.Add(1)
go func() {
	defer wg.Done()
	fmt.Println("work finished")
}()

wg.Wait()
```

Common pattern:
- `Add` before starting the goroutine
- `Done` when the goroutine finishes
- `Wait` where you need to pause until all work is complete

## 4. Channels

A channel sends values between goroutines.

```go
messages := make(chan string)
```

Send a value:

```go
messages <- "hello"
```

Receive a value:

```go
message := <-messages
```

Channels let goroutines communicate safely.

## 5. Unbuffered Channels

An unbuffered channel has no storage space.

```go
messages := make(chan string)
```

A send waits until another goroutine receives.

A receive waits until another goroutine sends.

This is useful when you want communication and synchronization at the same time.

## 6. Buffered Channels

A buffered channel has limited storage space.

```go
jobs := make(chan string, 2)
```

You can send values into the buffer until it is full.

```go
jobs <- "job one"
jobs <- "job two"
```

Buffered channels are useful when producers and consumers do not need to move at exactly the same speed.

## 7. Closing Channels and Range

Close a channel when no more values will be sent:

```go
close(jobs)
```

Receive all values with `range`:

```go
for job := range jobs {
	fmt.Println(job)
}
```

Beginner rule:

The sender closes the channel, not the receiver.

## 8. Select

`select` waits on multiple channel operations.

```go
select {
case message := <-messages:
	fmt.Println(message)
case <-time.After(time.Second):
	fmt.Println("timeout")
}
```

This is useful for:
- Waiting on whichever result arrives first
- Adding timeouts
- Coordinating multiple channels

## 9. Race Conditions

A race condition can happen when multiple goroutines access the same data at the same time and at least one goroutine writes to it.

Unsafe example:

```go
counter++
```

If many goroutines run this at the same time, the final value may be wrong.

Run the race detector:

```bash
go test -race
```

The race detector helps find unsafe shared memory access.

## 10. Run the Example

Open this example folder:

[module-14-concurrency-basics](../examples/module-14-concurrency-basics)

From the module example folder, run:

```bash
cd examples/module-14-concurrency-basics
go run .
```

Expected output will show:
- Goroutines running with a `WaitGroup`
- Values sent through channels
- A simple worker processing jobs
- `select` receiving a result or timeout

Run the tests:

```bash
go test -v
```

Run the race detector:

```bash
go test -race
```

## 11. Practice Tasks

1. Start three goroutines that print different messages and wait for all of them.
2. Create a channel that sends numbers from one goroutine to another.
3. Build a worker that receives jobs from a channel and prints each job.
4. Use `select` with `time.After` to add a timeout.
5. Write a small test and run it with `go test -race`.

## Checkpoint

You are ready for the next module when:
- You can start goroutines with `go`.
- You can wait for goroutines using `sync.WaitGroup`.
- You can send and receive values with channels.
- You understand the difference between buffered and unbuffered channels.
- You can use `select` for timeouts or multiple channels.
- You know why shared data can cause race conditions.
