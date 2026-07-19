# Go Learning Flow: Basic to Advanced

This roadmap is organized as modules. Move in order, build small programs often, and keep notes for each module in this repository.

## Current Progress

- Completed: Modules 1 to 16
- Current: Module 17 - Databases
- Next: Module 18 - CLI Applications

## Module 1: Setup and First Steps

**Goal:** Get comfortable running Go programs.

Detailed lesson:
- [Module 1: Setup and First Steps](modules/module-01-setup-and-first-steps.md)

Topics:
- Install Go and verify with `go version`
- Understand `go run`, `go build`, `go fmt`, and `go help`
- Create a basic `main.go`
- Understand packages and the `main` function
- Print output with `fmt.Println`

Practice:
- Write a hello world program
- Create a program that prints your name, city, and learning goal
- Build and run a compiled binary

Checkpoint:
- You can create, run, format, and build a simple Go program.

## Module 2: Variables, Types, and Operators

**Goal:** Understand Go's basic data model.

Detailed lesson:
- [Module 2: Variables, Types, and Operators](modules/module-02-variables-types-and-operators.md)

Topics:
- Variables with `var`
- Short declaration with `:=`
- Constants with `const`
- Basic types: `string`, `int`, `float64`, `bool`
- Zero values
- Type inference
- Type conversion
- Arithmetic, comparison, and logical operators

Practice:
- Build a simple calculator
- Convert temperature between Celsius and Fahrenheit
- Write a program that calculates simple interest

Checkpoint:
- You can choose correct basic types and perform simple calculations.

## Module 3: Control Flow

**Goal:** Make decisions and repeat work.

Detailed lesson:
- [Module 3: Control Flow](modules/module-03-control-flow.md)

Topics:
- `if`, `else if`, and `else`
- `switch`
- `for` loops
- Loop control with `break` and `continue`
- Basic input using `fmt.Scan`

Practice:
- Number guessing game
- Grade calculator
- Multiplication table generator
- Program to check if a number is prime

Checkpoint:
- You can write programs with branches and loops.

## Module 4: Functions

**Goal:** Break programs into reusable logic.

Detailed lesson:
- [Module 4: Functions](modules/module-04-functions.md)

Topics:
- Function declaration
- Parameters and return values
- Multiple return values
- Named return values
- Variadic functions
- Scope
- Anonymous functions

Practice:
- Create reusable math functions
- Write a function that returns quotient and remainder
- Build a small menu-driven calculator

Checkpoint:
- You can design small functions and pass data between them.

## Module 5: Arrays, Slices, and Maps

**Goal:** Work with collections.

Detailed lesson:
- [Module 5: Arrays, Slices, and Maps](modules/module-05-arrays-slices-and-maps.md)

Topics:
- Arrays
- Slices
- `append`
- Slice length and capacity
- Iteration with `range`
- Maps
- Checking if a map key exists
- Deleting map entries

Practice:
- Store and print a list of students
- Find max, min, and average from a slice
- Count word frequency using a map
- Build a simple contact book

Checkpoint:
- You can store, search, update, and iterate over grouped data.

## Module 6: Strings and Basic Standard Library

**Goal:** Use common built-in packages.

Detailed lesson:
- [Module 6: Strings and Basic Standard Library](modules/module-06-strings-and-basic-standard-library.md)

Topics:
- String operations
- `strings` package
- `strconv` package
- `math` package
- `time` package
- Basic error values from library functions

Practice:
- Validate a username
- Convert strings to numbers
- Format dates and times
- Build a simple text analyzer

Checkpoint:
- You can solve common tasks using standard library packages.

## Module 7: Structs and Methods

**Goal:** Model real-world data.

Detailed lesson:
- [Module 7: Structs and Methods](modules/module-07-structs-and-methods.md)

Topics:
- Struct declaration
- Struct literals
- Nested structs
- Methods
- Pointer receivers vs value receivers
- Basic object-style design in Go

Practice:
- Create a `Student` struct
- Create a `BankAccount` struct with deposit and withdraw methods
- Build an employee salary calculator

Checkpoint:
- You can group related data and behavior using structs and methods.

## Module 8: Pointers

**Status:** Completed

**Goal:** Understand references and mutation.

Detailed lesson:
- [Module 8: Pointers](modules/module-08-pointers.md)

Topics:
- Address operator `&`
- Dereference operator `*`
- Pointer parameters
- Pointers with structs
- When to use pointers
- Avoiding unnecessary pointer use

Practice:
- Swap two values using pointers
- Update a struct through a function
- Compare value receiver and pointer receiver behavior

Checkpoint:
- You understand how Go passes values and how pointers allow mutation.

## Module 9: Error Handling

**Status:** Completed

**Goal:** Write reliable programs.

Detailed lesson:
- [Module 9: Error Handling](modules/module-09-error-handling.md)

Topics:
- The `error` type
- Returning errors
- Checking errors
- Creating errors with `errors.New`
- Formatting errors with `fmt.Errorf`
- Error wrapping with `%w`
- `panic` and `recover` basics

Practice:
- Validate user input
- Return errors from calculator functions
- Create custom error messages
- Handle file read errors

Checkpoint:
- You can write functions that fail clearly and handle failures safely.

## Module 10: Files and JSON

**Status:** Completed

**Goal:** Read, write, and exchange data.

Detailed lesson:
- [Module 10: Files and JSON](modules/module-10-files-and-json.md)

Topics:
- Reading files
- Writing files
- Working with paths
- JSON encoding and decoding
- Struct tags
- Basic command-line arguments

Practice:
- Save notes to a file
- Read a text file and count lines
- Store contacts as JSON
- Build a small expense tracker using a JSON file

Checkpoint:
- You can persist data and work with JSON APIs or files.

## Module 11: Packages, Modules, and Project Structure

**Status:** Completed

**Goal:** Organize Go code properly.

Detailed lesson:
- [Module 11: Packages, Modules, and Project Structure](modules/module-11-packages-modules-project-structure.md)

Topics:
- `go mod init`
- `go.mod` and `go.sum`
- Import paths
- Creating your own packages
- Exported vs unexported identifiers
- Common folder structure
- Dependency management with `go get`

Practice:
- Split a calculator into packages
- Create a reusable validation package
- Add and use a third-party dependency

Checkpoint:
- You can organize code across files and packages.

## Module 12: Interfaces

**Status:** Completed

**Goal:** Use abstraction the Go way.

Detailed lesson:
- [Module 12: Interfaces](modules/module-12-interfaces.md)

Topics:
- Interface declaration
- Implicit implementation
- Interface values
- Empty interface / `any`
- Type assertions
- Type switches
- Small interfaces
- Dependency inversion basics

Practice:
- Create a `Shape` interface
- Build different payment method implementations
- Create a logger interface
- Write functions that accept behavior instead of concrete types

Checkpoint:
- You can use interfaces to make code flexible without over-designing.

## Module 13: Testing

**Status:** Completed

**Goal:** Prove your code works.

Detailed lesson:
- [Module 13: Testing](modules/module-13-testing.md)

Topics:
- `testing` package
- Writing unit tests
- Table-driven tests
- Test helpers
- Running tests with `go test`
- Test coverage
- Benchmark basics

Practice:
- Test calculator functions
- Write table-driven tests for validation logic
- Add tests for error cases
- Benchmark a string operation

Checkpoint:
- You can write useful tests for normal and failure paths.

## Module 14: Concurrency Basics

**Status:** Completed

**Goal:** Understand goroutines and channels.

Detailed lesson:
- [Module 14: Concurrency Basics](modules/module-14-concurrency-basics.md)

Topics:
- Goroutines
- Channels
- Buffered vs unbuffered channels
- `select`
- `sync.WaitGroup`
- Race conditions
- Running tests with `-race`

Practice:
- Run multiple tasks concurrently
- Build a worker example
- Send results through channels
- Fix a race condition

Checkpoint:
- You can safely run concurrent work and collect results.

## Module 15: Advanced Concurrency

**Status:** Completed

**Goal:** Build robust concurrent systems.

Detailed lesson:
- [Module 15: Advanced Concurrency](modules/module-15-advanced-concurrency.md)

Topics:
- Context cancellation
- Timeouts
- Worker pools
- Fan-in and fan-out patterns
- Mutexes
- Atomic operations basics
- Rate limiting
- Graceful shutdown

Practice:
- Build a worker pool
- Add timeout support with `context`
- Implement graceful shutdown for a long-running process
- Build a concurrent URL checker

Checkpoint:
- You can control concurrent programs instead of just starting goroutines.

## Module 16: HTTP and APIs

**Status:** Completed

**Goal:** Build web services.

Detailed lesson:
- [Module 16: HTTP and APIs](modules/module-16-http-apis.md)

Topics:
- `net/http`
- Handlers
- Routing basics
- Request and response handling
- JSON APIs
- HTTP status codes
- Middleware basics
- Environment variables

Practice:
- Build a hello API
- Create CRUD endpoints for tasks
- Add JSON request validation
- Add simple logging middleware

Checkpoint:
- You can build a basic REST-style API in Go.

## Module 17: Databases

**Status:** Current

**Goal:** Store application data in a real database.

Detailed lesson:
- [Module 17: Databases](modules/module-17-databases.md)

Topics:
- `database/sql`
- SQL basics
- Connecting to PostgreSQL or SQLite
- Queries and commands
- Scanning rows
- Transactions
- Migrations
- Repository pattern basics

Practice:
- Build a task API backed by a database
- Add create, list, update, and delete operations
- Use transactions for multi-step updates

Checkpoint:
- You can connect Go code to a database and manage persistent records.

## Module 18: CLI Applications

**Goal:** Build useful terminal tools.

Topics:
- Command-line arguments
- Flags with the `flag` package
- Reading from stdin
- Writing clean terminal output
- Exit codes
- Config files

Practice:
- Build a todo CLI
- Build a file search CLI
- Build a JSON formatter CLI

Checkpoint:
- You can create practical command-line programs.

## Module 19: Generics

**Goal:** Write reusable type-safe code.

Topics:
- Type parameters
- Constraints
- Generic functions
- Generic structs
- When generics help
- When interfaces are better

Practice:
- Create generic `Min` and `Max` functions
- Create a generic stack
- Create reusable map/filter helpers

Checkpoint:
- You can use generics for real reuse without making code harder to read.

## Module 20: Advanced Go Internals

**Goal:** Understand how Go behaves under the hood.

Topics:
- Memory allocation basics
- Stack vs heap intuition
- Escape analysis
- Garbage collection basics
- Slices internals
- Map internals overview
- Scheduler basics
- Profiling with `pprof`

Practice:
- Inspect escape analysis output
- Benchmark allocation-heavy code
- Profile CPU usage
- Profile memory usage

Checkpoint:
- You can reason about performance and memory at a practical level.

## Module 21: Production Practices

**Goal:** Build maintainable Go applications.

Topics:
- Logging
- Configuration
- Structured errors
- Dependency injection without heavy frameworks
- Clean package boundaries
- Observability basics
- Linting
- Formatting
- CI basics

Practice:
- Add structured logging to an API
- Add configuration through environment variables
- Add GitHub Actions for tests
- Run linting locally

Checkpoint:
- You can prepare Go projects for real-world use.

## Module 22: Capstone Projects

**Goal:** Combine everything into complete applications.

Project ideas:
- Todo REST API with database, tests, and graceful shutdown
- Expense tracker CLI with JSON or SQLite storage
- URL health checker with concurrency and reports
- Blog API with authentication basics
- File organizer CLI
- Chat server using WebSockets

Capstone requirements:
- Use modules and packages
- Include tests
- Handle errors clearly
- Read configuration from environment or flags
- Include a README with setup and usage
- Use Git commits for each milestone

Checkpoint:
- You can design, build, test, and explain a complete Go project.

## Suggested Weekly Plan

Week 1:
- Modules 1 to 3

Week 2:
- Modules 4 to 6

Week 3:
- Modules 7 to 9

Week 4:
- Modules 10 to 13

Week 5:
- Modules 14 and 15

Week 6:
- Modules 16 and 17

Week 7:
- Modules 18 to 21

Week 8:
- Module 22 capstone project

## Recommended Learning Habit

For every module:
1. Read the topic.
2. Write at least three small programs.
3. Refactor one program after learning a new concept.
4. Add notes about mistakes and fixes.
5. Commit your work with a clear message.

The best way to learn Go is to write small, boring programs first, then slowly make them more realistic.
