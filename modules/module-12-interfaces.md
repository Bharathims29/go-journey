# Module 12: Interfaces

## Goal

Learn how interfaces let Go functions depend on behavior instead of concrete types.

By the end of this module, you should understand:
- Interface declarations
- Implicit implementation
- Interface values
- Empty interface / `any`
- Type assertions
- Type switches
- Small interfaces
- Dependency inversion basics

## 1. What Is an Interface?

An interface describes behavior.

It lists method signatures that a type must have:

```go
type Shape interface {
	Area() float64
}
```

Any type with an `Area() float64` method satisfies this interface.

## 2. Implicit Implementation

Go does not use `implements`.

If a type has the methods required by an interface, it automatically satisfies the interface.

```go
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
```

`Rectangle` now satisfies `Shape`.

## 3. Using Interface Values

A function can accept an interface instead of one concrete type:

```go
func printArea(shape Shape) {
	fmt.Println(shape.Area())
}
```

Now `printArea` can work with any type that has an `Area` method.

## 4. Small Interfaces

Go prefers small interfaces.

Common examples:

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

Beginner rule:

Accept interfaces where behavior is needed. Return concrete types when creating values.

## 5. Empty Interface and `any`

`any` means the same thing as `interface{}`.

It can hold a value of any type:

```go
func printValue(value any) {
	fmt.Println(value)
}
```

Use `any` carefully. When everything is `any`, the compiler cannot help you as much.

## 6. Type Assertions

A type assertion gets a concrete value out of an interface value.

```go
value := any("go")

text, ok := value.(string)
if ok {
	fmt.Println(text)
}
```

Use the two-value form with `ok` to avoid panics.

## 7. Type Switches

A type switch handles several possible concrete types.

```go
func describe(value any) {
	switch v := value.(type) {
	case string:
		fmt.Println("string:", v)
	case int:
		fmt.Println("int:", v)
	default:
		fmt.Println("unknown type")
	}
}
```

This is useful when input can be one of several known types.

## 8. Dependency Inversion Basics

Interfaces help code depend on behavior instead of details.

Example:

```go
type Logger interface {
	Log(message string)
}

func processPayment(logger Logger) {
	logger.Log("payment processed")
}
```

`processPayment` does not care whether logs go to the terminal, a file, or a remote service.

## 9. Run the Example

Open this example file:

[main.go](../examples/module-12-interfaces/main.go)

From the module example folder, run:

```bash
cd examples/module-12-interfaces
go run .
```

Expected output will show:
- Different shapes used through one interface
- A logger interface used by a function
- A type assertion
- A type switch

## 10. Practice Tasks

1. Create a `Shape` interface with `Area` and implement it for `Circle` and `Rectangle`.
2. Create a `PaymentMethod` interface and implement `Pay` for card and cash payments.
3. Create a `Logger` interface and pass it into a function that performs work.
4. Write a function that accepts `any` and uses a type switch for `string`, `int`, and `bool`.
5. Refactor one previous example so a function accepts a small interface instead of a concrete type.

## Checkpoint

You are ready for the next module when:
- You can define an interface.
- You understand interfaces are satisfied implicitly.
- You can pass different concrete types into one interface-based function.
- You can use a safe type assertion.
- You know why small interfaces are preferred in Go.
