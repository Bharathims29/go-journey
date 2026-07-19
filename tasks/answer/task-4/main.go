package main

import "fmt"

func addContacts(contacts map[string]string) {
	var choice string

	for {
		var name, phone string

		fmt.Print("Enter Contact Name: ")
		fmt.Scan(&name)

		fmt.Print("Enter Phone Number: ")
		fmt.Scan(&phone)

		contacts[name] = phone

		fmt.Print("Add another contact? (y/n): ")
		fmt.Scan(&choice)

		if choice == "n" || choice == "N" {
			break
		}
	}
}

func searchContacts(contacts map[string]string) {
	var choice string

	for {
		var name string

		fmt.Print("\nEnter Contact Name to Search: ")
		fmt.Scan(&name)

		if phone, found := contacts[name]; found {
			fmt.Println("Phone Number:", phone)
		} else {
			fmt.Println("Contact not found.")
		}

		fmt.Print("Search another contact? (y/n): ")
		fmt.Scan(&choice)

		if choice == "n" || choice == "N" {
			break
		}
	}
}

func updateContact(contacts map[string]string) {
	var choice string

	for {
		var name, phone string

		fmt.Print("\nEnter Contact Name to Update: ")
		fmt.Scan(&name)

		if _, found := contacts[name]; found {
			fmt.Print("Enter New Phone Number: ")
			fmt.Scan(&phone)

			contacts[name] = phone
			fmt.Println("Contact updated successfully.")
		} else {
			fmt.Println("Contact not found.")
		}

		fmt.Print("Update another contact? (y/n): ")
		fmt.Scan(&choice)

		if choice == "n" || choice == "N" {
			break
		}

	}
}

func deleteContacts(contacts map[string]string) {
	var choice string

	for {
		var name string

		fmt.Print("\nEnter Contact Name to Delete: ")
		fmt.Scan(&name)

		if _, found := contacts[name]; found {
			delete(contacts, name)
			fmt.Println("Contact deleted successfully.")
		} else {
			fmt.Println("Contact not found.")
		}

		fmt.Print("Delete another contact? (y/n): ")
		fmt.Scan(&choice)

		if choice == "n" || choice == "N" {
			break
		}
	}
}

func printContacts(contacts map[string]string) {
	if len(contacts) == 0 {
		fmt.Println("No contacts available.")
		return
	}

	for name, phone := range contacts {
		fmt.Printf("Name: %-15s Phone: %s\n", name, phone)
	}
}

func main() {
	contacts := make(map[string]string)

	addContacts(contacts)

	fmt.Println("\nAll Contacts:")
	printContacts(contacts)

	searchContacts(contacts)

	updateContact(contacts)

	deleteContacts(contacts)

	fmt.Println("\nFinal Contact List:")
	printContacts(contacts)
}
