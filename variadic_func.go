package main

import "fmt"

// varaible argument kirimkan dengan bentuk parameter
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(10, 20, 30))
	fmt.Println(sumAll(10, 20, 30, 10))
	fmt.Println(sumAll(10, 20, 30, 90))


	numbers := []int{1, 2, 3, 4, 5} // menggunakan slice := [] {1, 2,}
	fmt.Println(sumAll(numbers...))
}