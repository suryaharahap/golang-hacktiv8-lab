package main

import "fmt"

func main() {

	// MARK: Membuat sebuah perulangan dengan for loop
	fmt.Println("PIRAMIDA: \n ")
	rows := 10
	var x int

	for i := 1; i <= rows; i++ {
		x = 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print(" ")
		}

		for {
			fmt.Print("*")
			x++
			if x == 2*i-1 {
				break
			}

		}
		fmt.Println("")
	}
}
