package main

import "fmt"

func main() {
	// Buatlah 3 buah variabel dengan tipe data string, int, dan boolean.
	name := "John Doe"
	age := 10
	isGoLearn := true

	// Mencetak nilai variabel tersebut dengan menggunakan GO verb yang sesuai dengan tipe datanya
	fmt.Printf("%s\n", name)      // String
	fmt.Printf("%d\n", age)       // Integer
	fmt.Printf("%t\n", isGoLearn) // Boolean

	// Mencetak tipe data dari masing-masing variabel
	fmt.Printf("%T\n", name)      // String
	fmt.Printf("%T\n", age)       // Integer
	fmt.Printf("%T\n", isGoLearn) // Boolean
}
