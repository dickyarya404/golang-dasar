package main

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Banyumas", "Jawa Tengah", "Indonesia"}
	// address2 := address1 // copy value

	// address2.City = "Purwokerto"
	// address2.Country = "Japan"
	// fmt.Println(address1) // tidak berubah 
	// fmt.Println(address2) // berubah

	// var address1 Address = Address{"Banyumas", "Jawa Tengah", "Indonesia"}
	// var address2 *Address = &address1 // pointer menuju data asli *

	// address2.City = "Purwokerto"
	
	// fmt.Println(address1) // ikut berubah
	// fmt.Println(address2) // berubah

	



}

// pointer adalah kemampuan membuat referance ke lokasi data di memory yang sama, 
// tanpa mendupliksi data yang sudah ada
// sederhananya kita dapat memebuat pass bye referance
// tambahan operator &