package main

import "fmt"

// interface adalah tipe data abstract, beriskan definsi method

type HasName interface {
	GetName() string  // method GetName dengan parameter string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string // method
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string // method
	
}

func (animal Animal) GetName() string {
	return animal.Name
	
}

func main() {
	person := Person{Name: "Dicky"}
	SayHello(person)
	
	animal := Animal{Name: "Kucing"}
	SayHello(animal)


}