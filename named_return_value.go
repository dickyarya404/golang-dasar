package main

import "fmt"

func getComplateName() (firstName bool, lastName int, middleName string) {
	firstName = true
	lastName = 20
	middleName = "Python"
	return firstName, lastName, middleName
}

func main() {
	a, b, c := getComplateName()
	fmt.Println(a, b, c)
}