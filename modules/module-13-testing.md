# Module 13: Testing

## Goal

Learn how to prove that your Go code works by writing clear automated tests.

By the end of this module, you should understand:
- The `testing` package
- Unit tests
- Table-driven tests
- Test helpers
- Error case tests
- Running tests with `go test`
- Test coverage
- Basic benchmarks

## 1. Why Testing Matters

Tests let you check behavior automatically.

Instead of running a program and checking output manually every time, you write test functions that Go can run for you.

Good tests help you:
- Catch mistakes early
- Refactor with confidence
- Document expected behavior
- Check both success and failure cases

## 2. Test Files

Go test files end with `_test.go`.

Example:

```text
calculator.go
calculator_test.go
```

Test files usually sit beside the code they test.

## 3. Test Functions

A test function starts with `Test` and receives `*testing.T`.

```go
func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Fatalf("Add(2, 3) = %d, want %d", got, want)
	}
}
```

Use `t.Fatalf` when the test should stop immediately.

Use `t.Errorf` when the test can continue after reporting the problem.

## 4. Table-Driven Tests

Table-driven tests are common in Go.

They let you test many inputs with one test function:

```go
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive numbers", a: 2, b: 3, want: 5},
		{name: "negative number", a: -2, b: 3, want: 1},
		{name: "zero", a: 0, b: 7, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
```

`t.Run` creates a named subtest for each case.

## 5. Testing Errors

When a function returns an error, test both paths:
- Valid input should return no error
- Invalid input should return an error

Example:

```go
func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
```

Do not only test happy paths. Error behavior is part of your program too.

## 6. Test Helpers

A helper is a small function used by tests.

Call `t.Helper()` inside helper functions. This makes failure messages point to the test line instead of the helper line.

```go
func requireNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
```

Helpers keep tests readable.

## 7. Running Tests

Run tests in the current package:

```bash
go test
```

Run tests with detailed output:

```bash
go test -v
```

Run every package below the current folder:

```bash
go test ./...
```

## 8. Test Coverage

Coverage shows how much code was executed by tests.

```bash
go test -cover
```

Coverage is useful, but it is not the full story. A test can execute a line without checking the right behavior.

Beginner rule:

Use coverage as a guide, but focus on meaningful test cases.

## 9. Benchmarks

Benchmarks measure performance.

A benchmark function starts with `Benchmark` and receives `*testing.B`.

```go
func BenchmarkNormalizeName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalizeName("  bharath  ")
	}
}
```

Run benchmarks:

```bash
go test -bench=.
```

`b.N` is chosen by Go. Your benchmark body should run the operation being measured.

## 10. Run the Example

Open this example folder:

[module-13-testing](../examples/module-13-testing)

From the module example folder, run:

```bash
cd examples/module-13-testing
go test -v
```

Run coverage:

```bash
go test -cover
```

Run the benchmark:

```bash
go test -bench=.
```

The example includes:
- Unit tests for calculator functions
- Table-driven tests for validation logic
- Error case tests
- A test helper
- A string benchmark

## 11. Practice Tasks

1. Write tests for an `Add`, `Subtract`, `Multiply`, and `Divide` calculator.
2. Write table-driven tests for a function that validates usernames.
3. Add tests for invalid input, such as division by zero or empty names.
4. Create a helper function that checks errors.
5. Write a benchmark for a string cleanup function.

## Checkpoint

You are ready for the next module when:
- You can create a `_test.go` file.
- You can write a basic test using `testing.T`.
- You can write table-driven tests with subtests.
- You can test expected errors.
- You can run `go test`, `go test -cover`, and `go test -bench=.`
