package main

import "fmt"

type Filter func(string) string // ini tpe declaration untuk mengganti type data : string, bool, int

func sayHelloWithFilter(name string, filter Filter) { // Filter berasal dari type
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}


func spamFilter(name string) string {
	if name == "Toxic" {
		return "***" // logic sensor filter
	} else {  
		return name
	}
}

func main() {
	sayHelloWithFilter("Jhon", spamFilter)

	filter := spamFilter // func filter
	
	sayHelloWithFilter("Toxic", filter)

}