package main

import "fmt"

type Customer struct {
	Name, Address, Email string // struct adalah template data atau prototype data tdk dpt langusunf digunkan
	Age, Nim             int
}

func (customer Customer) sayHello(name string) { // method struct fuunc sayhello
	fmt.Println("Hello: ", name, "my name is ", customer.Name)
}

func main() {
	var Joe Customer
	Joe.Name = "Joe Rose"
	Joe.Address = "Jakarta"
	Joe.Email = "joe@gmail.com"
	Joe.Age = 23
	Joe.Nim = 202221007
	fmt.Println(Joe)
	// represintasi ddata

	Jhon := Customer{
		Name: "jhon",
		Address: "Indonesia",
		Age: 34,
		Nim: 202332043,
	}

	fmt.Println(Jhon)

	Jhon = Customer{Name: "jhon", Address: "Indonesia", Age: 34, Nim: 202332043} // simple
	fmt.Println(Jhon)

	Joe.sayHello("John")

}