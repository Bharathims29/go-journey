package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	fmt.Println("Module 10: Files and JSON")
	fmt.Println()

	textPath := filepath.Join("notes.txt")
	textContent := []byte("Learning Go file handling.\nLearning Go JSON.\n")

	fmt.Println("Write text file")
	err := os.WriteFile(textPath, textContent, 0644)
	if err != nil {
		fmt.Println("Error writing text file:", err)
		return
	}
	fmt.Println("Created:", textPath)
	fmt.Println()

	fmt.Println("Read text file")
	readContent, err := os.ReadFile(textPath)
	if err != nil {
		fmt.Println("Error reading text file:", err)
		return
	}
	fmt.Println(string(readContent))

	contacts := []Contact{
		{Name: "Bharath", Email: "bharath@example.com"},
		{Name: "Luffy", Email: "luffy@example.com"},
	}

	fmt.Println("Write JSON file")
	jsonPath := filepath.Join("contacts.json")
	jsonData, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return
	}

	err = os.WriteFile(jsonPath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
	fmt.Println("Created:", jsonPath)
	fmt.Println()

	fmt.Println("Read JSON file")
	storedJSON, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var decodedContacts []Contact
	err = json.Unmarshal(storedJSON, &decodedContacts)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, contact := range decodedContacts {
		fmt.Printf("Name: %s, Email: %s\n", contact.Name, contact.Email)
	}
	fmt.Println()

	fmt.Println("Command-line arguments")
	if len(os.Args) > 1 {
		fmt.Println("First argument:", os.Args[1])
	} else {
		fmt.Println("No extra command-line argument provided")
	}
}
