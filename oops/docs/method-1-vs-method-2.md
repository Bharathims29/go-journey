# Difference Between Method 1 and Method 2

This note explains the two object creation styles used in `oops/main.go`.

## Method 1: Create a Person Using `New`

```go
p1 := person.New("Bharathi", "MS", 25)
p2 := person.New("Luffy", "Monkey", 30)
```

Here, `New` creates and returns a complete `Person`.

This is useful when you already have all required values.

### Benefits

- The object is ready immediately.
- The setup code is short.
- It is harder to forget a field.
- It keeps creation logic inside the `person` package.

### Best Use Case

Use this style when a value should be created with valid data from the beginning.

## Method 2: Create an Empty Person and Use Setters

```go
p := person.Person{}
p.SetFirstName("Bharathi")
p.SetLastName("MS")
p.SetAge(25)
```

Here, an empty `Person` is created first. After that, setter methods update the fields one by one.

This works because the setter methods use pointer receivers:

```go
func (p *Person) SetFirstName(fName string) {
	p.firstName = fName
}
```

The `*Person` receiver allows the method to change the original `Person`.

### Benefits

- Fields can be updated later.
- It is useful when values arrive step by step.
- It shows how pointer receiver methods modify a struct.

### Risk

The object can exist in an incomplete state before all setters are called.

For example, after this line:

```go
p := person.Person{}
```

the `Person` exists, but its fields still contain zero values.

## Main Difference

| Topic | Method 1: `New` | Method 2: Setters |
| --- | --- | --- |
| Creation style | Create with values immediately | Create empty, then update |
| Object state | Complete from the start | May be incomplete at first |
| Code length | Shorter | Longer |
| Receiver lesson | Shows constructor-style function | Shows pointer receiver methods |
| Best for | Required data | Step-by-step updates |

## Simple Rule

Use `New` when you want to create a ready-to-use value.

Use setters when you need to update a value after it already exists.

