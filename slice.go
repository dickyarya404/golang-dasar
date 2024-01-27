package main

import (
	"fmt"
)

func main() {
	names := [...]string{"Jhon", "Joe", "Joni", "Johan", "Juan", "Jox", "Jenis"}

	slice1 := names[4:7]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)


	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // sabtu, minggu

	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu baru" //
	daySlice1[1] = "Minggu baru" //
	

	fmt.Println(daySlice1) // [Sabtu baru Minggu baru]
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jum'at Sabtu baru Minggu baru]

	daySlice2 := append(daySlice1, "Libur Baru") //
	daySlice2[0] = "Tahun Baru" //
	fmt.Println(daySlice1) // [Sabtu baru Minggu baru]
	fmt.Println(daySlice2) // [Sabtu baru Minggu baru Libur Baru]
	

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Jhon"
	newSlice[1] = "Jhon"
	// newSlice[2] = "Jhon" erorr harus menggunakan append
	

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Dicky") // tambah dengan menggunakan append
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Joe"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// copy data slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// perbedaan array dan slice yang paling sering digunakan golang adalah Slice
	iniArray := [...]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4} 

	fmt.Println(iniArray)
	fmt.Println(iniSlice)



	
}