# Module 10: Files and JSON

## Goal

Learn how to read and write files in Go and how to work with JSON data using structs.

By the end of this module, you should understand:
- Reading files
- Writing files
- Working with paths
- JSON encoding and decoding
- Struct tags
- Basic command-line arguments

## 1. Reading Files

Go provides the `os` package for common file operations.

Use `os.ReadFile` to read the full content of a small file:

```go
data, err := os.ReadFile("notes.txt")
if err != nil {
	fmt.Println("Error:", err)
	return
}

fmt.Println(string(data))
```

The result is returned as `[]byte`, so convert it to `string` when printing text.

## 2. Writing Files

Use `os.WriteFile` to write data into a file.

```go
content := []byte("Go is fun to learn.\n")
err := os.WriteFile("notes.txt", content, 0644)
if err != nil {
	fmt.Println("Error:", err)
	return
}
```

`0644` is a common file permission for text files.

## 3. Working with Paths

Use the `path/filepath` package for portable path handling.

```go
filePath := filepath.Join("data", "contacts.json")
```

This helps your code work across operating systems.

## 4. JSON Encoding

JSON is commonly used to store and exchange structured data.

Use `json.MarshalIndent` to convert Go data into formatted JSON:

```go
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

contact := Contact{Name: "Bharath", Email: "bharath@example.com"}

data, err := json.MarshalIndent(contact, "", "  ")
```

## 5. JSON Decoding

Use `json.Unmarshal` to convert JSON into a Go struct.

```go
var contact Contact
err := json.Unmarshal(data, &contact)
if err != nil {
	fmt.Println("Error:", err)
}
```

Notice that `&contact` is passed, because `Unmarshal` must update the original value.

## 6. Struct Tags

Struct tags control how fields are encoded and decoded.

```go
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
```

Without tags, JSON uses the Go field names directly.

Tags are useful when:
- JSON field names should be lowercase
- External APIs expect specific names
- You want to omit fields or rename them

## 7. Basic Command-Line Arguments

Go provides command-line arguments through `os.Args`.

```go
fmt.Println(os.Args)
```

`os.Args[0]` is the program name.

Values after that are the arguments passed by the user.

Example:

```bash
go run . notes.txt
```

Then:

- `os.Args[0]` is the compiled program path
- `os.Args[1]` is `notes.txt`

## 8. Run the Example

Open this example file:

[main.go](../examples/module-10-files-json/main.go)

From the module example folder, run:

```bash
cd examples/module-10-files-json
go run .
```

Expected output will show:
- A text file being created
- The text file being read back
- A JSON file being written
- The JSON file being decoded into Go structs
- A command-line argument check

## 9. Practice Tasks

1. Write a program that saves your daily notes to a text file.
2. Read a text file and count how many lines it contains.
3. Store a list of contacts in a JSON file.
4. Load JSON data from a file into a struct.
5. Accept a file name using `os.Args` and print its contents.

## Checkpoint

You are ready for the next module when:
- You can read and write a file using `os.ReadFile` and `os.WriteFile`.
- You can encode structs into JSON.
- You can decode JSON into structs.
- You understand why struct tags are used.
- You can read simple command-line arguments.

