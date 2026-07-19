# Module 9: Error Handling

## Goal

Learn how to write Go functions that fail clearly and how to handle those failures safely.

By the end of this module, you should understand:
- The `error` type
- Returning errors from functions
- Checking errors with `if err != nil`
- Creating errors with `errors.New`
- Formatting errors with `fmt.Errorf`
- Wrapping errors with `%w`
- Basic `panic` and `recover`

## 1. What Is an Error?

In Go, an error is a normal return value.

The built-in `error` type represents a problem:

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
```

If there is no problem, return `nil` for the error.

## 2. Checking Errors

Always check the error before using the result.

```go
result, err := divide(10, 0)
if err != nil {
	fmt.Println("Error:", err)
	return
}

fmt.Println(result)
```

Beginner rule:

Handle the error close to where it happens.

## 3. Returning Errors

When a function can fail, return an `error` as the last return value.

```go
func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}

	return nil
}
```

Call it like this:

```go
err := validateAge(-1)
if err != nil {
	fmt.Println(err)
}
```

## 4. Creating Errors with `errors.New`

Use `errors.New` when the message is fixed.

```go
return errors.New("name cannot be empty")
```

This is simple and clear.

## 5. Formatting Errors with `fmt.Errorf`

Use `fmt.Errorf` when the message needs values.

```go
return fmt.Errorf("age %d is below minimum age %d", age, minimumAge)
```

This helps create useful error messages.

## 6. Wrapping Errors with `%w`

Wrapping keeps the original error while adding more context.

```go
func loadConfig() error {
	err := readFile()
	if err != nil {
		return fmt.Errorf("load config failed: %w", err)
	}

	return nil
}
```

Use `errors.Is` to check wrapped errors.

```go
if errors.Is(err, ErrNotFound) {
	fmt.Println("missing file")
}
```

## 7. `panic` and `recover`

Most normal errors should use `error`, not `panic`.

Use `panic` only when the program reaches a state it cannot safely continue from.

`recover` can catch a panic inside a deferred function.

```go
func safeRun() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	panic("something went very wrong")
}
```

For normal application logic, prefer returning errors.

## 8. Run the Example

Open this example file:

[main.go](../examples/module-09-error-handling/main.go)

From the module example folder, run:

```bash
cd examples/module-09-error-handling
go run .
```

Expected output will show:
- A successful division
- A division error
- Input validation errors
- Error wrapping
- Checking wrapped errors with `errors.Is`
- A recovered panic

## 9. Practice Tasks

1. Write a `divide` function that returns an error when the second number is zero.
2. Write a `validateUsername` function that rejects empty names and names shorter than three characters.
3. Write a calculator function that returns errors for invalid operations.
4. Create a package-level error value and check it with `errors.Is`.
5. Write a small function that uses `panic` and recover from it only for practice.

## Checkpoint

You are ready for the next module when:
- You can return an error from a function.
- You can check `err != nil`.
- You can choose between `errors.New` and `fmt.Errorf`.
- You understand why `%w` is useful.
- You know that `panic` is not for normal expected errors.

