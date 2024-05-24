package main

import "fmt"

func main() {
	sayHelloWithFilter("Surya", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
