# Module 2: Variables, Types, and Operators

## Goal

Understand how Go stores values and performs basic operations.

By the end of this module, you should understand:
- How to create variables
- How to create constants
- Common basic types in Go
- Zero values
- Type inference
- Type conversion
- Arithmetic operators
- Comparison operators
- Logical operators

## 1. What Is a Variable?

A variable is a named place where you store a value.

Example:

```go
var name string = "Bharath"
```

This means:
- `var` creates a variable
- `name` is the variable name
- `string` is the type
- `"Bharath"` is the value

## 2. Different Ways to Create Variables

### Full variable declaration

```go
var city string = "Bengaluru"
```

This is clear and explicit.

### Type inference

```go
var age = 25
```

Go understands that `age` is an `int` because the value is a whole number.

### Short declaration

```go
language := "Go"
```

This is the most common style inside functions.

You can use `:=` only inside a function.

## 3. Constants

A constant is a value that should not change.

```go
const country = "India"
```

Use constants for fixed values such as app names, tax rates, limits, or messages.

## 4. Basic Types

Common Go types:

```go
string
int
float64
bool
```

Examples:

```go
var name string = "Bharath"
var age int = 25
var height float64 = 5.9
var isLearning bool = true
```

## 5. Zero Values

If you create a variable without giving it a value, Go gives it a default value.

```go
var name string
var age int
var price float64
var active bool
```

Zero values:
- `string` becomes `""`
- `int` becomes `0`
- `float64` becomes `0`
- `bool` becomes `false`

Go does this to keep variables predictable.

## 6. Type Conversion

Go does not automatically convert between different numeric types.

Example:

```go
var marks int = 95
var average float64 = float64(marks)
```

Here, `float64(marks)` converts the integer into a floating-point number.

## 7. Arithmetic Operators

```go
+ addition
- subtraction
* multiplication
/ division
% remainder
```

Example:

```go
a := 10
b := 3

fmt.Println(a + b)
fmt.Println(a - b)
fmt.Println(a * b)
fmt.Println(a / b)
fmt.Println(a % b)
```

## 8. Comparison Operators

Comparison operators return `true` or `false`.

```go
== equal
!= not equal
>  greater than
<  less than
>= greater than or equal
<= less than or equal
```

Example:

```go
age := 20
fmt.Println(age >= 18)
```

## 9. Logical Operators

Logical operators combine boolean values.

```go
&& and
|| or
!  not
```

Example:

```go
age := 20
hasID := true

fmt.Println(age >= 18 && hasID)
```

This prints `true` only when both conditions are true.

## 10. Run the Example

Open this example file:

[main.go](../examples/module-02-variables/main.go)

From the module example folder, run:

```bash
cd examples/module-02-variables
go run .
```

Expected output will show:
- Variable values
- Constant values
- Zero values
- Type conversion
- Arithmetic results
- Comparison results
- Logical results

## 11. Practice Tasks

1. Create variables for your name, age, city, and learning status.
2. Print all variables.
3. Create constants for country and course name.
4. Calculate the total price of three items.
5. Convert an `int` value to `float64`.
6. Check whether a person is eligible to vote.
7. Check whether a student passed using marks and attendance.

## Mini Project: Simple Profile and Score Calculator

Create a program that stores:
- Name
- Age
- City
- Course name
- Three subject marks

Then print:
- The profile details
- Total marks
- Average marks
- Whether the student passed

Passing rule:
- Average should be at least `40`

## Checkpoint

You are ready for Module 3 when you can:
- Create variables using `var`
- Create variables using `:=`
- Create constants using `const`
- Explain zero values
- Use `string`, `int`, `float64`, and `bool`
- Convert one type to another
- Use arithmetic operators
- Use comparison operators
- Use logical operators
