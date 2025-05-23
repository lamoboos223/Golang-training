package main

import "fmt"

// Person represents a simple person with a name
type Person struct {
	Name string
}

// NewPerson creates a new Person and returns a pointer to it
func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}

// UpdateName updates the person's name
func (p *Person) UpdateName(newName string) {
	p.Name = newName
}

func main() {
	// Create a new person using the constructor
	person := NewPerson("John")
	fmt.Println("Original name:", person.Name) // Output: Original name: John

	// Update the name using the method
	person.UpdateName("Jane")
	fmt.Println("Updated name:", person.Name) // Output: Updated name: Jane

	// Create another person
	anotherPerson := NewPerson("Bob")
	fmt.Println("Another person:", anotherPerson.Name) // Output: Another person: Bob
}
