# Module 7: Structs and Methods

## Goal

Learn how to model real-world data using structs and attach behavior using methods.

By the end of this module, you should understand:
- Struct declaration
- Struct literals
- Nested structs
- Methods
- Value receivers
- Pointer receivers
- Basic object-style design in Go

## 1. What Is a Struct?

A struct groups related values together.

Example:

```go
type Student struct {
	Name  string
	Age   int
	Marks float64
}
```

This creates a new type named `Student`.

The fields are:
- `Name`
- `Age`
- `Marks`

Structs are useful when one thing has multiple pieces of data.

## 2. Struct Literals

A struct literal creates a value from a struct type.

```go
student := Student{
	Name:  "Bharath",
	Age:   25,
	Marks: 88.5,
}
```

You can access fields using dot syntax:

```go
fmt.Println(student.Name)
fmt.Println(student.Age)
```

You can also update fields:

```go
student.Marks = 91.0
```

## 3. Nested Structs

A struct can contain another struct.

```go
type Address struct {
	City    string
	Country string
}

type Employee struct {
	Name    string
	Role    string
	Address Address
}
```

Create it:

```go
employee := Employee{
	Name: "Anu",
	Role: "Developer",
	Address: Address{
		City:    "Bengaluru",
		Country: "India",
	},
}
```

Access nested fields:

```go
fmt.Println(employee.Address.City)
```

## 4. Methods

A method is a function attached to a type.

```go
func (s Student) Display() {
	fmt.Println(s.Name, s.Age, s.Marks)
}
```

Here, `(s Student)` is called the receiver.

Call the method:

```go
student.Display()
```

Methods help keep behavior close to the data it belongs to.

## 5. Value Receivers

A value receiver gets a copy of the struct.

```go
func (s Student) IsPassed() bool {
	return s.Marks >= 40
}
```

This is good when the method only reads data and does not change the struct.

## 6. Pointer Receivers

A pointer receiver can change the original struct value.

```go
func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}
```

Use a pointer receiver when:
- The method should modify the original value
- The struct is large and copying it would be inefficient
- You want consistency across methods of the same type

Pointers will be explained more deeply in Module 8.

## 7. Basic Object-Style Design in Go

Go does not use classes.

Instead, Go commonly uses:
- Structs for data
- Methods for behavior
- Interfaces for abstraction

Example:

```go
type BankAccount struct {
	Owner   string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}
```

This gives you a clean way to group data and actions.

## 8. Run the Example

Open this example file:

[main.go](../examples/module-07-structs-methods/main.go)

From the module example folder, run:

```bash
cd examples/module-07-structs-methods
go run .
```

Expected output will show:
- Creating a `Student` struct
- Updating struct fields
- Calling methods
- Nested structs
- A `BankAccount` with deposit and withdraw methods
- Value receiver and pointer receiver behavior

## 9. Practice Tasks

1. Create a `Book` struct with title, author, and price.
2. Create a method that prints book details.
3. Create a `Student` struct and an `IsPassed` method.
4. Create an `Employee` struct with nested `Address`.
5. Create a `BankAccount` struct.
6. Add `Deposit` and `Withdraw` methods to `BankAccount`.
7. Create a method that updates a student's marks.

## Mini Project: Employee Salary Calculator

Create an `Employee` struct with:
- Name
- Role
- Base salary
- Bonus

Add methods:
- `TotalSalary() float64`
- `Display()`
- `AddBonus(amount float64)`

Use a pointer receiver for `AddBonus` because it changes the employee.

## Checkpoint

You are ready for Module 8 when you can:
- Create a struct type
- Create struct values
- Read and update struct fields
- Create nested structs
- Write methods
- Choose a value receiver for read-only methods
- Choose a pointer receiver for methods that update data
