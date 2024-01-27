package main

import "fmt"

func sayHelloTo(firstName int, lastName string) {
	fmt.Println("Hello gaes", firstName, lastName)

}

func main() {
	sayHelloTo(1, "Arya")
	sayHelloTo(3, "Jhon")

}