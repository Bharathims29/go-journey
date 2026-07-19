# Module 8: Pointers

## Goal

Understand how Go passes values and how pointers allow functions and methods to update original data.

By the end of this module, you should understand:
- Address operator `&`
- Dereference operator `*`
- Pointer parameters
- Pointers with structs
- When to use pointers
- Avoiding unnecessary pointer use

## 1. What Is a Pointer?

A pointer stores the memory address of another value.

Normal variable:

```go
age := 25
```

Pointer to that variable:

```go
agePointer := &age
```

The pointer does not store `25` directly. It stores the address where `age` lives.

## 2. Address Operator `&`

Use `&` to get the address of a variable.

```go
name := "Bharath"
namePointer := &name

fmt.Println(namePointer)
```

The printed address may look different every time you run the program.

That is normal.

## 3. Dereference Operator `*`

Use `*` to get or change the value stored at a pointer address.

```go
age := 25
agePointer := &age

fmt.Println(*agePointer)
```

This prints the value inside `age`.

You can also update the original value:

```go
*agePointer = 30
fmt.Println(age)
```

This prints `30`.

## 4. Go Passes Values by Copy

When you pass a normal variable to a function, Go passes a copy.

```go
func changeValue(number int) {
	number = 100
}

func main() {
	number := 10
	changeValue(number)
	fmt.Println(number)
}
```

This still prints `10`.

The function changed only its copy.

## 5. Pointer Parameters

If a function should update the original value, pass a pointer.

```go
func changeValue(number *int) {
	*number = 100
}

func main() {
	number := 10
	changeValue(&number)
	fmt.Println(number)
}
```

This prints `100`.

`&number` sends the address.

`*number = 100` updates the value at that address.

## 6. Swap Two Values Using Pointers

Pointers are useful when a function must change multiple original values.

```go
func swap(a *int, b *int) {
	*a, *b = *b, *a
}
```

Call it:

```go
x := 10
y := 20

swap(&x, &y)
```

After the call:
- `x` becomes `20`
- `y` becomes `10`

## 7. Pointers with Structs

Pointers are commonly used with structs.

```go
type Student struct {
	Name  string
	Marks float64
}

func updateMarks(student *Student, marks float64) {
	student.Marks = marks
}
```

Go allows this shortcut:

```go
student.Marks = marks
```

You do not need to write:

```go
(*student).Marks = marks
```

Go automatically dereferences struct pointers when accessing fields.

## 8. Pointer Receivers

You already used pointer receivers in Module 7.

Example:

```go
func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}
```

This method changes the original `BankAccount`.

Use pointer receivers when the method should update the struct.

## 9. When to Use Pointers

Use pointers when:
- A function or method must change the original value
- Passing a large struct by copy would be inefficient
- You want a method to update struct fields

Avoid pointers when:
- The function only needs to read a small value
- A normal return value is simpler
- Pointer use makes the code harder to understand

Beginner rule:

If you only need to read data, start without a pointer.

If you need to update original data, use a pointer.

## 10. Run the Example

Open this example file:

[main.go](../examples/module-08-pointers/main.go)

From the module example folder, run:

```bash
cd examples/module-08-pointers
go run .
```

Expected output will show:
- Address with `&`
- Dereferencing with `*`
- Pass-by-value behavior
- Pointer parameter behavior
- Swapping values
- Updating a struct through a pointer
- Pointer receiver behavior

## 11. Practice Tasks

1. Create an integer variable and print its value and address.
2. Create a pointer to a string and print the original value using `*`.
3. Write a function that tries to change an integer without a pointer and observe the result.
4. Write a function that changes an integer using a pointer.
5. Write a `swap` function using pointers.
6. Create a `Student` struct and update marks using a pointer function.
7. Create a `BankAccount` struct and update balance using pointer receiver methods.

## Mini Project: Profile Updater

Create a `Profile` struct with:
- Name
- City
- Skill

Add functions:
- `updateCity(profile *Profile, city string)`
- `updateSkill(profile *Profile, skill string)`
- `displayProfile(profile Profile)`

Use pointer parameters for the update functions because they change the original profile.

## Checkpoint

You are ready for Module 9 when you can:
- Explain what `&` does
- Explain what `*` does
- Pass a pointer to a function
- Update original values using pointers
- Swap two values using pointers
- Update struct fields through a pointer
- Choose when not to use pointers
