# Module 1: Setup and First Steps

## Goal

Get comfortable creating, running, formatting, and building your first Go program.

By the end of this module, you should understand:
- What a Go program looks like
- How to run a Go file
- How to build a Go file into an executable
- Why every executable Go program starts with `package main`
- What the `main` function does
- How to print output using `fmt.Println`

## 1. Check Go Installation

First, check whether Go is installed:

```bash
go version
```

If Go is installed, you will see output similar to:

```text
go version go1.22.0 linux/amd64
```

The exact version may be different. That is okay.

## 2. Your First Go Program

Open this example file:

[main.go](../examples/module-01-setup/main.go)

The file contains:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("My name is Bharath.")
	fmt.Println("I am learning Go from basic to advanced.")
}
```

## 3. Explanation

### `package main`

Every Go file starts with a package declaration.

```go
package main
```

`package main` tells Go that this program can be run directly as an application.

If you want to create a runnable program, use `package main`.

### `import "fmt"`

```go
import "fmt"
```

This imports Go's `fmt` package.

The `fmt` package is used for formatted input and output. In this module, we use it to print text to the terminal.

### `func main()`

```go
func main() {
}
```

This is the starting point of a Go program.

When you run a Go application, Go looks for the `main` function and starts executing code from there.

### `fmt.Println`

```go
fmt.Println("Hello, Go!")
```

`fmt.Println` prints text to the terminal and then moves to a new line.

## 4. Run the Program

From the project root, run:

```bash
go run examples/module-01-setup/main.go
```

Expected output:

```text
Hello, Go!
My name is Bharath.
I am learning Go from basic to advanced.
```

`go run` compiles and runs the program immediately. It is useful while learning and testing small programs.

## 5. Format the Program

Go has a built-in formatter. Run:

```bash
gofmt -w examples/module-01-setup/main.go
```

This automatically formats the code in the standard Go style.

In Go, formatting is not something every developer decides differently. The Go tool gives one standard format.

## 6. Build the Program

To create an executable file, run:

```bash
go build -o hello-go examples/module-01-setup/main.go
```

Then run the executable:

```bash
./hello-go
```

`go build` creates a binary file that you can run later without using `go run`.

## 7. Useful Commands

```bash
go version
```

Shows the installed Go version.

```bash
go run examples/module-01-setup/main.go
```

Runs the Go program.

```bash
gofmt -w examples/module-01-setup/main.go
```

Formats the Go file.

```bash
go build -o hello-go examples/module-01-setup/main.go
```

Builds the Go program into an executable.

```bash
go help
```

Shows help for Go commands.

## 8. Practice Tasks

1. Change the name in the example program.
2. Add one more line that prints your city.
3. Add one more line that prints why you want to learn Go.
4. Run the program again.
5. Format the file using `gofmt`.
6. Build the program using `go build`.

## Checkpoint

You are ready for Module 2 when you can:
- Create a `.go` file
- Explain `package main`
- Explain `func main`
- Use `fmt.Println`
- Run a Go program with `go run`
- Format a Go file with `gofmt`
- Build a Go program with `go build`
