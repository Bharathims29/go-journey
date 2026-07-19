# Module 11: Packages, Modules, and Project Structure

## Goal

Learn how to split Go code into packages, understand module files, and organize small projects clearly.

By the end of this module, you should understand:
- What a Go module is
- What `go.mod` and `go.sum` do
- How import paths work
- How to create your own packages
- Exported and unexported identifiers
- A simple project folder structure
- How `go get` adds dependencies

## 1. What Is a Module?

A module is a Go project with its own dependency list.

Create one with:

```bash
go mod init example.com/my-project
```

This creates a `go.mod` file.

The module path is the root import path for your project.

Example:

```go
module example.com/my-project

go 1.22
```

If your module path is `example.com/my-project`, a package inside `calc` can be imported as:

```go
import "example.com/my-project/calc"
```

## 2. What Is `go.mod`?

`go.mod` stores:
- The module path
- The Go version
- Direct dependencies
- Sometimes indirect dependencies

Example:

```go
module example.com/my-project

go 1.22

require github.com/google/uuid v1.6.0
```

Commit `go.mod` to Git.

## 3. What Is `go.sum`?

`go.sum` stores checksums for downloaded dependencies.

It helps Go verify that dependencies have not changed unexpectedly.

Commit `go.sum` too when it exists.

You usually do not edit `go.sum` manually.

## 4. Packages

A package groups related Go files.

Every Go file starts with a package declaration:

```go
package main
```

The `main` package builds an executable program.

Other package names usually describe what the code does:

```go
package calc
package validate
package store
```

All `.go` files in the same folder must use the same package name.

## 5. Creating Your Own Package

Folder structure:

```text
my-project/
  go.mod
  main.go
  calc/
    calc.go
```

`calc/calc.go`:

```go
package calc

func Add(a, b int) int {
	return a + b
}
```

`main.go`:

```go
package main

import (
	"fmt"

	"example.com/my-project/calc"
)

func main() {
	fmt.Println(calc.Add(2, 3))
}
```

## 6. Exported vs Unexported Identifiers

In Go, names that start with a capital letter are exported.

Exported names can be used from another package:

```go
func Add(a, b int) int {
	return a + b
}
```

Names that start with a lowercase letter are unexported.

Unexported names can only be used inside the same package:

```go
func isPositive(n int) bool {
	return n > 0
}
```

Beginner rule:

Export only what other packages need.

## 7. Import Paths

Standard library imports use short paths:

```go
import "fmt"
```

Your own packages use the module path plus the folder path:

```go
import "example.com/my-project/validate"
```

Third-party packages use their public module path:

```go
import "github.com/google/uuid"
```

## 8. Common Small Project Structure

For beginner projects, keep the structure simple:

```text
project/
  go.mod
  main.go
  calc/
    calc.go
  validate/
    validate.go
```

Avoid creating many folders before the program needs them.

For larger applications, you may later see:

```text
project/
  cmd/
  internal/
  pkg/
```

Do not rush into advanced layouts. Start simple and let the project grow naturally.

## 9. Dependency Management with `go get`

Use `go get` to add a third-party package:

```bash
go get github.com/google/uuid
```

This updates `go.mod` and usually creates or updates `go.sum`.

Use `go mod tidy` to remove unused dependencies and add missing ones:

```bash
go mod tidy
```

Beginner rule:

After changing imports or dependencies, run:

```bash
go mod tidy
```

## 10. Run the Example

Open this example file:

[main.go](../examples/module-11-packages/main.go)

From the module example folder, run:

```bash
cd examples/module-11-packages
go run .
```

Expected output will show:
- A calculator package being used
- A validation package being used
- Exported functions called from `main`
- Unexported helper functions used inside their own package

## 11. Practice Tasks

1. Split a calculator program into a `calc` package and a `main` package.
2. Create a `validate` package with functions for checking names, ages, and emails.
3. Create a `store` package that saves and loads JSON data.
4. Rename one exported function to lowercase and observe the compiler error from another package.
5. Add a third-party dependency with `go get`, use it once, then run `go mod tidy`.

## Checkpoint

You are ready for the next module when:
- You can create a Go module with `go mod init`.
- You can explain why `go.mod` exists.
- You can import your own package from another folder.
- You understand capitalized names are exported.
- You can keep a small Go project organized without overcomplicating it.
