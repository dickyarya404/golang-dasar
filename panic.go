package main

import "fmt"

func endApp() {
	fmt.Println("End Application")
	message := recover()
	fmt.Println("Terjadi panic", message)
	// cara penanganan eror yg bisa menghentika aplikasi 
	// recover merupakan func yg bisa digunakan untuk mendapatkan data panic
	// panic merupakan func bisa digunaknan untuk menghentikan program
	

}

func runApp(error bool) {
	defer endApp() // func endApp()

	if error {
		panic("Ups Error")
	}

}

func main() {
	runApp(true)
	fmt.Println("Dicky Arya Aji Prayoga")

}