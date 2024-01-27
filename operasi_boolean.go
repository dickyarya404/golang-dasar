package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsen = 78

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = nilaiAbsen > 76

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}