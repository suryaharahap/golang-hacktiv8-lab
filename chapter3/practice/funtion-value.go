package main

import "fmt"

func main() {
	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("Budi"))
	fmt.Println(misal("Lord"))
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
