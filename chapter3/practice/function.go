package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("From function great: ")
	great("Surya", 22)

	fmt.Println("From function greatAddress: ")
	greatAddress("Dani", "Bandung")

	fmt.Println("Function return")
	names := []string{"Airell", "Hunis\n"}
	printMsg := greats2("Heii", names)
	fmt.Println(printMsg)

	// MARK: Return multiple value
	fmt.Println("Return multiple value")
	diameter := 15
	var area, circumference = calculate(float64(diameter))
	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)

	// MARK: variadic function
	fmt.Println("Variadic function")
	studentList := printStudentList("Airel", "Marion", "Danu", "Listiow", "Narse")
	fmt.Printf("%v", studentList)

	// MARK: variadic function
	fmt.Println("\n\nVariadic function 2")
	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := sum(numberLists...)
	fmt.Println("Result:", result)

	profile("Nnidi", "Pasta", "Ayam Geprek", "Sate Padang")

	result2 := getHello("Surya")
	fmt.Println(result2)

	nameUniverity, _ := getFullUniversity()
	fmt.Println(nameUniverity)

}

func great(name string, age int8) {
	fmt.Printf("Hello there!  My name is %s and I'm %d years old.\n\n", name, age)
}

func greatAddress(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address+"\n")
}

func greats2(msg string, names []string) string {
	joinStr := strings.Join(names, " ")

	result := fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

func calculate(d float64) (float64, float64) {
	// Menghitung luas
	area := math.Pi * math.Pow(d/2, 2)

	// Menghitung keliling
	circumference := math.Pi * d

	// Mengembalikan 2 nilai
	return area, circumference
}

func printStudentList(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Printf("Hello there! I'm  %s I realy love to eat %s\n", name, mergeFavFoods)
}

func getHello(name string) string {
	hello := "Hello " + name

	return hello
}

func getFullUniversity() (string, string) {
	return "Universitas Indonesia", "Jakarta"
}
