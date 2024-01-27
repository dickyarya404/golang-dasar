package main

import "fmt"

func main() {
	counter := 0 

	increment := func() { // func berinteraksi dengan data disekitar nya
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}