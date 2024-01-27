package helper

var version = "1.0.0" // tidsak bisa di akses
var Application = "backend golang"

func sayGoodBye(name string) string { // tidak bisa di akses
	return "Good bye" + name
}

// func huruf kapital bisa diakses di package yang berbeda
func SayHello(name string) string {
	return "Hello " + name
}