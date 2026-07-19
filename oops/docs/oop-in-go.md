# Object-Oriented Programming in Go

## What Is OOP?

Object-oriented programming is a way to organize code by grouping data and behavior together.

In many languages, this is done with classes and objects. Go does not have classes, but it still supports object-style programming using structs, methods, interfaces, and packages.

## OOP Concepts in Go

### 1. Encapsulation

Encapsulation means hiding data and allowing access through methods.

In this example, the `Person` struct has lowercase fields:

```go
type Person struct {
	firstName string
	lastName  string
	age       int
}
```

Because the fields start with lowercase letters, they are private to the `person` package. Other packages must use methods like `FirstName`, `SetFirstName`, and `Age`.

### 2. Abstraction

Abstraction means showing only the important behavior and hiding internal details.

For example, code in `main.go` does not need to know how age is stored. It only calls:

```go
fmt.Println(p.Age())
```

The `Age` method decides how to return the value.

### 3. Composition

Go prefers composition over inheritance.

Composition means building a larger type by placing one struct inside another struct.

```go
type Address struct {
	City string
}

type Employee struct {
	Person  Person
	Address Address
}
```

This is the common Go way to reuse structure and behavior.

### 4. Polymorphism

Polymorphism means different types can be used through the same behavior.

In Go, this is usually done with interfaces.

```go
type Speaker interface {
	Speak() string
}
```

Any type with a `Speak() string` method automatically satisfies this interface.

## Types Used for OOP Style in Go

- `struct`: stores related data
- `method`: attaches behavior to a type
- `pointer receiver`: changes the original value
- `value receiver`: reads data from a copy
- `interface`: describes behavior without depending on a concrete type
- `package`: controls visibility and groups related code

## Important Go Difference

Go does not have:

- Classes
- Constructors as a language feature
- Inheritance
- Method overriding like Java or C++

Go uses simple building blocks instead:

- Structs for data
- Methods for behavior
- Interfaces for flexible design
- Composition for reuse

