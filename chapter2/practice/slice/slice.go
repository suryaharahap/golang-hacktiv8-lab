package main

import "fmt"

func main() {
	names := [...]string{"Nana", "Nindo", "Budi", "Silsi", "Yiren", "Gotun"}
	fruits := [...]string{"apple", "banana", "mango", "grape", "orange"}

	/*
		// Implementasi/deklarasi slice secara manual
		var slice1 []string = names[1:3]

		# Slice tidak mementukan nilainya/ukurannya sedangkan Array membutuhkannya.
	*/

	slice1 := names[1:3]
	fmt.Println(slice1) // [Nindo Budi]

	slice2 := names[:3]
	fmt.Println(slice2) // [Nana Nindo Budi]

	var slice3 []string = names[2:]
	fmt.Println(slice3) // [Budi Silsi Yiren Gotun]

	var slice4 []string = names[:]
	fmt.Println(slice4) // [Nana Nindo Budi Silsi Yiren Gotun]

	fruits1 := fruits[3:]
	fmt.Println(fruits1)

}
