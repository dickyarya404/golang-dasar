package main

import "fmt"

func getHello(name string) string {
	hello := "Hello" + name
	return hello // return untuk mengembalikan datanya

}

func main() {
	result := getHello("Joe")
	fmt.Println(result)

	fmt.Println(getHello("Dicky"))
	fmt.Println(getHello("Arya"))
}