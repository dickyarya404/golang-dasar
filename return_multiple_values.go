package main

import "fmt"

func getFullName() (string, string) {
	return "Golang", "Java"
}
// Multiple return value wajib ditangkap semua value
func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullName() // jika tidak butuh gunakan _
	fmt.Println(firstName)


}