package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye // goodbye jadi function
	contoh := getGoodBye

	fmt.Println(goodbye("Javascript"))
	fmt.Println(contoh("Golang"))
}