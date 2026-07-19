# Module 3: Control Flow

## Goal

Learn how to make decisions and repeat work in Go.

By the end of this module, you should understand:
- `if`, `else if`, and `else`
- `switch`
- `fallthrough`
- `for` loops
- `break`
- `continue`
- Basic input with `fmt.Scan`

## 1. What Is Control Flow?

Control flow means controlling which code runs and how many times it runs.

In Go, the main control flow tools are:
- `if` for decisions
- `switch` for choosing between many cases
- `for` for loops

## 2. `if`, `else if`, and `else`

Use `if` when your program needs to make a decision.

```go
age := 20

if age >= 18 {
	fmt.Println("Adult")
} else {
	fmt.Println("Minor")
}
```

You can use `else if` when there are multiple conditions.

```go
marks := 82

if marks >= 90 {
	fmt.Println("Grade A")
} else if marks >= 75 {
	fmt.Println("Grade B")
} else if marks >= 50 {
	fmt.Println("Grade C")
} else {
	fmt.Println("Needs improvement")
}
```

## 3. `switch`

Use `switch` when you want to compare one value against many possible cases.

```go
day := "Monday"

switch day {
case "Monday":
	fmt.Println("Start of the week")
case "Saturday", "Sunday":
	fmt.Println("Weekend")
default:
	fmt.Println("Regular weekday")
}
```

`default` runs when no case matches.

## 4. `fallthrough`

In Go, a `switch` stops after the matching case runs.

Unlike some other languages, Go does not automatically continue into the next case.

If you want the next case to run, use `fallthrough`.

```go
level := 2

switch level {
case 1:
	fmt.Println("Beginner")
case 2:
	fmt.Println("Intermediate")
	fallthrough
case 3:
	fmt.Println("Advanced practice unlocked")
default:
	fmt.Println("Unknown level")
}
```

Output:

```text
Intermediate
Advanced practice unlocked
```

Use `fallthrough` carefully. It runs the next case body without checking that case condition.

Most Go programs do not need `fallthrough` often, but you should know it exists.

## 5. `for` Loops

Go has only one loop keyword: `for`.

Basic loop:

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
```

This means:
- Start with `i := 1`
- Continue while `i <= 5`
- Increase `i` after each loop using `i++`

## 6. Looping Like `while`

Go does not have a separate `while` keyword.

You can write a while-style loop using `for`:

```go
count := 1

for count <= 3 {
	fmt.Println(count)
	count++
}
```

## 7. `break`

Use `break` to stop a loop early.

```go
for i := 1; i <= 10; i++ {
	if i == 5 {
		break
	}

	fmt.Println(i)
}
```

This stops when `i` becomes `5`.

## 8. `continue`

Use `continue` to skip the current loop step and move to the next one.

```go
for i := 1; i <= 5; i++ {
	if i == 3 {
		continue
	}

	fmt.Println(i)
}
```

This skips printing `3`.

## 9. Basic Input with `fmt.Scan`

You can read input from the terminal using `fmt.Scan`.

```go
var name string

fmt.Print("Enter your name: ")
fmt.Scan(&name)

fmt.Println("Hello,", name)
```

The `&name` part means Go should store the typed value inside the `name` variable.

You will understand `&` more deeply when you learn pointers. For now, remember that `fmt.Scan` needs it.

## 10. Run the Example

Open this example file:

[main.go](../examples/module-03-control-flow/main.go)

From the module example folder, run:

```bash
cd examples/module-03-control-flow
go run .
```

Expected output will show:
- Age decision using `if`
- Grade decision using `else if`
- Day selection using `switch`
- Explicit switch fallthrough
- Counting with `for`
- Skipping a value with `continue`
- Stopping a loop with `break`
- Prime number checking

## 11. Practice Tasks

1. Write a program that checks whether a number is positive, negative, or zero.
2. Write a grade calculator using `if`, `else if`, and `else`.
3. Write a program that prints numbers from `1` to `10`.
4. Write a multiplication table for any number.
5. Write a program that prints only even numbers from `1` to `20`.
6. Write a program that checks whether a number is prime.
7. Write a small menu using `switch`.
8. Write a `switch` example that uses `fallthrough`.

## Mini Project: Number Guessing Game

Build a simple number guessing game.

Rules:
- Store a secret number in a variable.
- Ask the user to guess the number.
- If the guess is too high, print `Too high`.
- If the guess is too low, print `Too low`.
- If the guess is correct, print `Correct`.

Start with one guess only. Later, improve it using a loop so the user can keep guessing.

## Checkpoint

You are ready for Module 4 when you can:
- Use `if`, `else if`, and `else`
- Use `switch`
- Explain when `fallthrough` runs the next case
- Write a `for` loop
- Stop a loop with `break`
- Skip a loop step with `continue`
- Read simple user input with `fmt.Scan`
- Write small decision-based programs
