# Module 4: Functions

## Goal

Learn how to break a program into reusable blocks of logic.

By the end of this module, you should understand:
- Function declaration
- Parameters
- Return values
- Multiple return values
- Named return values
- Variadic functions
- Scope
- Anonymous functions

## 1. What Is a Function?

A function is a reusable block of code that performs a specific task.

Instead of writing the same logic again and again, you put the logic inside a function and call it when needed.

Example:

```go
func greet() {
	fmt.Println("Hello, Go learner!")
}
```

Call the function like this:

```go
greet()
```

## 2. Function Declaration

Basic syntax:

```go
func functionName() {
	// code
}
```

Example:

```go
func sayWelcome() {
	fmt.Println("Welcome to Go")
}
```

## 3. Parameters

Parameters let you pass values into a function.

```go
func greetUser(name string) {
	fmt.Println("Hello,", name)
}
```

Call it:

```go
greetUser("Bharath")
```

The value `"Bharath"` is passed into the `name` parameter.

## 4. Return Values

A function can return a value.

```go
func add(a int, b int) int {
	return a + b
}
```

Call it:

```go
total := add(10, 20)
fmt.Println(total)
```

The final `int` in the function declaration means the function returns an integer.

## 5. Multiple Return Values

Go functions can return more than one value.

```go
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}
```

Call it:

```go
q, r := divide(10, 3)
fmt.Println(q, r)
```

This is common in Go, especially when returning a result and an error.

## 6. Named Return Values

You can name the return values in the function signature.

```go
func rectangle(width int, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}
```

The plain `return` returns the named values.

Use named returns carefully. They are useful for small, clear functions, but they can make long functions harder to read.

## 7. Variadic Functions

A variadic function accepts any number of values of the same type.

```go
func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
```

Call it:

```go
fmt.Println(sum(10, 20, 30))
fmt.Println(sum(5, 15))
```

The `...int` means the function can accept many `int` values.

## 8. Scope

Scope means where a variable can be used.

```go
func showScope() {
	message := "I exist only inside this function"
	fmt.Println(message)
}
```

The `message` variable can be used only inside `showScope`.

If you try to use `message` outside the function, Go will show an error.

## 9. Anonymous Functions

An anonymous function is a function without a name.

```go
double := func(number int) int {
	return number * 2
}

fmt.Println(double(5))
```

This is useful when you need a small function for a short time.

## 10. Run the Example

Open this example file:

[main.go](../examples/module-04-functions/main.go)

From the module example folder, run:

```bash
cd examples/module-04-functions
go run .
```

Expected output will show:
- A simple function call
- A function with parameters
- A function with one return value
- A function with multiple return values
- A function with named return values
- A variadic function
- A scoped variable
- An anonymous function

## 11. Practice Tasks

1. Write a function that prints your name.
2. Write a function that accepts a name and prints a greeting.
3. Write a function that adds two numbers and returns the result.
4. Write a function that returns the square of a number.
5. Write a function that returns quotient and remainder.
6. Write a variadic function that returns the total of any number of prices.
7. Write an anonymous function that triples a number.

## Mini Project: Menu Calculator

Create a small calculator using functions.

Required functions:
- `add(a int, b int) int`
- `subtract(a int, b int) int`
- `multiply(a int, b int) int`
- `divide(a int, b int) (int, int)`

The `divide` function should return quotient and remainder.

After that, use `fmt.Scan` and `switch` to let the user choose an operation.

## Checkpoint

You are ready for Module 5 when you can:
- Create and call functions
- Pass values using parameters
- Return one value
- Return multiple values
- Understand named return values
- Use a variadic function
- Explain basic variable scope
- Use an anonymous function
