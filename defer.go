package main

import "fmt"

func logging() {
	fmt.Println("funciton selesai di panggil")
}

func runApllication () {
	defer logging() 
	// defer func akan selalu di eseskusi walaupun terjad eror di func
	// defer func merupakan func yg bisa di jadwalkan di esekusi setelah sebuah func selesai di esekusi
	fmt.Println("Run Appllication")

} 

func main() {

	runApllication()
}