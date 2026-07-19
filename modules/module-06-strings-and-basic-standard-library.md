# Module 6: Strings and Basic Standard Library

## Goal

Learn how to use common built-in Go packages to solve everyday problems.

By the end of this module, you should understand:
- Basic string operations
- The `strings` package
- The `strconv` package
- The `math` package
- The `time` package
- Basic error values from library functions

## 1. What Is the Standard Library?

The Go standard library is a set of packages that come with Go.

You do not need to install these packages separately.

Examples:
- `fmt` for printing
- `strings` for string operations
- `strconv` for converting strings and numbers
- `math` for math functions
- `time` for dates and times

## 2. Basic String Operations

A string stores text.

```go
name := "Bharathi"
```

You can join strings using `+`.

```go
message := "Hello, " + name
```

You can get the length of a string using `len`.

```go
fmt.Println(len(name))
```

For simple English text, `len` gives the number of bytes, which often looks like the number of characters. For text with emojis or some non-English characters, bytes and characters can be different.

## 3. `strings` Package

The `strings` package gives useful functions for working with text.

```go
import "strings"
```

Common functions:

```go
strings.ToUpper("go")
strings.ToLower("GO")
strings.Contains("golang", "go")
strings.TrimSpace("  hello  ")
strings.ReplaceAll("go is good", "good", "great")
strings.Split("go,java,python", ",")
strings.Join([]string{"go", "java"}, " | ")
```

## 4. `strconv` Package

The `strconv` package converts strings to other types and other types to strings.

```go
import "strconv"
```

String to integer:

```go
age, err := strconv.Atoi("25")
```

Integer to string:

```go
ageText := strconv.Itoa(25)
```

String to float:

```go
price, err := strconv.ParseFloat("99.50", 64)
```

The `64` means parse the value with `float64` precision.

Go has two common floating-point types:

```go
float32
float64
```

Most of the time, use `64` because `float64` is more precise and is commonly used by Go's math-related functions.

You can also pass `32`, but `strconv.ParseFloat` still returns a `float64`. Passing `32` only means the value is rounded as if it were a `float32`.

Many conversion functions return an error because the input might be invalid.

## 5. Basic Error Values

Some standard library functions return two values:
- The result
- An error

Example:

```go
number, err := strconv.Atoi("abc")

if err != nil {
	fmt.Println("Invalid number:", err)
} else {
	fmt.Println(number)
}
```

`nil` means there is no error.

If `err != nil`, something went wrong.

This pattern is very common in Go.

## 6. `math` Package

The `math` package provides math functions.

```go
import "math"
```

Examples:

```go
math.Sqrt(64)
math.Pow(2, 3)
math.Round(10.6)
math.Ceil(10.2)
math.Floor(10.8)
```

Most `math` functions use `float64`.

## 7. `time` Package

The `time` package works with dates, times, and durations.

```go
import "time"
```

Current time:

```go
now := time.Now()
```

Format time:

```go
fmt.Println(now.Format("2006-01-02"))
```

Go uses a special reference date for formatting:

```text
Mon Jan 2 15:04:05 MST 2006
```

So if you want year-month-day, use:

```go
"2006-01-02"
```

## 8. Run the Example

Open this example file:

[main.go](../examples/module-06-standard-library/main.go)

From the module example folder, run:

```bash
cd examples/module-06-standard-library
go run .
```

Expected output will show:
- String joining and length
- `strings` package examples
- `strconv` conversions
- Handling a conversion error
- `math` package examples
- `time` formatting examples
- A small username validation example

## 9. Practice Tasks

1. Convert a name to uppercase.
2. Remove extra spaces from a sentence.
3. Check whether an email contains `@`.
4. Convert `"100"` to an integer and add `50`.
5. Convert an integer age to a string.
6. Calculate the square root of `144`.
7. Print today's date in `YYYY-MM-DD` format.
8. Validate a username using `strings.TrimSpace` and `len`.

## Mini Project: Text Analyzer

Create a program that stores a paragraph in a string.

Print:
- Original text
- Text in lowercase
- Text in uppercase
- Number of bytes using `len`
- Whether the text contains the word `go`
- Number of words using `strings.Fields`

Example:

```go
words := strings.Fields(text)
fmt.Println(len(words))
```

## Checkpoint

You are ready for Module 7 when you can:
- Use common `strings` functions
- Convert strings to numbers with `strconv`
- Check and handle `err`
- Use basic `math` functions
- Get and format the current time
- Build simple text-processing programs
