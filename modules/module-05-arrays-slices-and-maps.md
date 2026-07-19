# Module 5: Arrays, Slices, and Maps

## Goal

Learn how to store and work with multiple values.

By the end of this module, you should understand:
- Arrays
- Slices
- `append`
- Slice length and capacity
- Iteration with `range`
- Maps
- Checking if a map key exists
- Deleting map entries

## 1. What Are Collections?

Collections help you store multiple values in one variable.

In Go, the most common collection types are:
- Arrays
- Slices
- Maps

Arrays and slices store values in order.

Maps store values using keys.

## 2. Arrays

An array has a fixed size.

```go
var numbers [3]int = [3]int{10, 20, 30}
```

This array can store exactly three integers.

You can access values using an index:

```go
fmt.Println(numbers[0])
```

Indexes start from `0`, so:
- `numbers[0]` is the first value
- `numbers[1]` is the second value
- `numbers[2]` is the third value

Arrays are useful when the size should never change.

## 3. Slices

A slice is like a flexible array.

```go
scores := []int{80, 90, 75}
```

Slices are used much more often than arrays because they can grow.

## 4. `append`

Use `append` to add values to a slice.

```go
scores := []int{80, 90, 75}
scores = append(scores, 88)
```

Important: `append` returns a new slice value, so assign it back.

```go
scores = append(scores, 95)
```

## 5. Length and Capacity

`len` gives the number of items in a slice.

`cap` gives the capacity of the slice before Go may need to allocate more space.

```go
fmt.Println(len(scores))
fmt.Println(cap(scores))
```

For beginners, focus more on `len`. Capacity becomes more useful when learning performance.

## 6. Iteration with `range`

Use `range` to loop over arrays, slices, and maps.

```go
for index, score := range scores {
	fmt.Println(index, score)
}
```

If you do not need the index, use `_`.

```go
for _, score := range scores {
	fmt.Println(score)
}
```

## 7. Maps

A map stores key-value pairs.

```go
studentMarks := map[string]int{
	"Bharath": 90,
	"Anu":     85,
}
```

Here:
- The key type is `string`
- The value type is `int`

You can read a value using its key:

```go
fmt.Println(studentMarks["Bharath"])
```

## 8. Add or Update Map Values

Use the key to add or update a value.

```go
studentMarks["Ravi"] = 88
studentMarks["Anu"] = 91
```

If the key does not exist, Go adds it.

If the key already exists, Go updates it.

## 9. Check If a Key Exists

When reading from a map, Go can return two values:

```go
marks, exists := studentMarks["Bharath"]
```

`exists` is `true` when the key is present.

```go
if exists {
	fmt.Println("Marks:", marks)
} else {
	fmt.Println("Student not found")
}
```

This is useful because if a key does not exist, Go returns the zero value for the value type.

## 10. Delete Map Entries

Use `delete` to remove a key from a map.

```go
delete(studentMarks, "Anu")
```

If the key does not exist, `delete` does nothing.

## 11. Run the Example

Open this example file:

[main.go](../examples/module-05-collections/main.go)

From the module example folder, run:

```bash
cd examples/module-05-collections
go run .
```

Expected output will show:
- Array values
- Slice values
- Slice length and capacity
- Iteration using `range`
- Map creation
- Map lookup
- Map update
- Map delete
- Finding max, min, and average from a slice

## 12. Practice Tasks

1. Create an array with five numbers and print each value.
2. Create a slice of student names and add two more names using `append`.
3. Create a slice of marks and calculate total and average.
4. Find the maximum number in a slice.
5. Find the minimum number in a slice.
6. Create a map of student names and marks.
7. Check whether a student exists in the map.
8. Delete one student from the map.

## Mini Project: Contact Book

Create a simple contact book using a map.

Use:

```go
map[string]string
```

The key should be the contact name.

The value should be the phone number.

Required features:
- Add a contact
- Update a contact
- Search for a contact
- Delete a contact
- Print all contacts

Start with fixed data in code. Later, you can improve it with user input.

## Checkpoint

You are ready for Module 6 when you can:
- Create an array
- Create a slice
- Add values to a slice using `append`
- Use `len`
- Understand basic slice capacity
- Loop using `range`
- Create a map
- Add and update map values
- Check if a map key exists
- Delete a map entry
