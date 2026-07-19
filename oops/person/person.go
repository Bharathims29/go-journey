package person

import "strconv"

// Person represents a person in an object-style design.
// The fields are unexported, so code outside this package cannot change them directly.
type Person struct {
	firstName string
	lastName  string
	age       int
}

// New creates a Person with all required values at once.
// This is similar to using a constructor in class-based OOP languages.
func New(firstName, lastName string, age int) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

// SetFirstName updates the first name of an existing Person.
// A pointer receiver is used because this method changes the original value.
func (p *Person) SetFirstName(fName string) {
	p.firstName = fName
}

// SetLastName updates the last name of an existing Person.
// A pointer receiver is used because this method changes the original value.
func (p *Person) SetLastName(lName string) {
	p.lastName = lName
}

// SetAge updates the age of an existing Person.
// A pointer receiver is used because this method changes the original value.
func (p *Person) SetAge(a int) {
	p.age = a
}

// FirstName returns the person's first name.
// A value receiver is enough because this method only reads data.
func (p Person) FirstName() string {
	return p.firstName
}

// LastName returns the person's last name.
// A value receiver is enough because this method only reads data.
func (p Person) LastName() string {
	return p.lastName
}

// Age returns the person's age as a string for display.
// A value receiver is enough because this method only reads data.
func (p Person) Age() string {
	return strconv.Itoa(p.age)
}

// getAge is unexported because it starts with a lowercase letter.
// It can be used only inside the person package.
func (p Person) getAge() string {
	return p.firstName + " " + p.lastName + " age is " + strconv.Itoa(p.age)
}

// FamilyName returns a display message using the person's last name.
func (p Person) FamilyName() string {
	return p.lastName + " is his family name"
}
