package main

import "fmt"

// Person struct dengan properti public dan private
type Person struct {
	Name string // Public (Exported)
	age  int    // Private (Unexported)
}

// NewPerson adalah fungsi public yang mengembalikan pointer ke Person
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		age:  age,
	}
}

// Greet adalah metode public yang dapat diakses dari luar paket
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s.\n", p.Name)
}

// Age adalah metode public untuk mendapatkan nilai age yang private
func (p Person) Age() int {
	return p.age
}

// setAge adalah metode private untuk mengatur nilai age
func (p *Person) setAge(newAge int) {
	p.age = newAge
}
