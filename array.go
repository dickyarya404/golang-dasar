package main

import "fmt"

func main() {
	var names [5]string // array tidak boleh dari 5

	names[0] = "Dicky"
	names[1] = "Arya"
	names[2] = "Aji"
	names[3] = "Prayoga"
	names[4] = "Partama"
	

	fmt.Printf(names[0])
	fmt.Printf(names[1])
	fmt.Printf(names[2])
	fmt.Printf(names[3])
	fmt.Printf(names[4])


	// Tidak ada operasi untuk menghapus data di dalam array, hanya bisa mengkosongkan data nya 
	var values = [...]int {
		90,
		89,
		50,
		90,
		100,

	}
	fmt.Println(values)
	fmt.Println(len(values))
}