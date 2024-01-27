package main

import "fmt"

func main() {

	// var person map[string]string = map[string]string{}
	// person["name"] = "John"
	// persion["addres"] = "Jakarta"
	// key dan vvalue = "person"
	person := map[string]string{
		"name":   "John",
		"addres": "Jakarta",
	}

	fmt.Println(person["name"])
	fmt.Println(person["addres"])
	fmt.Println(person)

	book := make(map[string]string)
	book["Title"] = "Buku Golang"
	book["Author"] = "Dicky"
	book["Content"] = "Buku Programming"

	fmt.Println(book)

	delete(book, "Content")
	fmt.Println(book)


}