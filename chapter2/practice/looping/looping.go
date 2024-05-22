package main

import "fmt"

func main() {
	// MARK: LOOPINGS (FIRST WAYS OF LOOPING)
	fmt.Println("FIRST WAYS OF LOOPING: ")
	for i := 0; i < 3; i++ {
		fmt.Println("Hello World", i)
	}

	// MARK: LOOPINGS (SECOND WAYS OF LOOPING)
	fmt.Println("\nSECOND WAYS OF LOOPING: ")
	x := 0
	for x < 3 {
		fmt.Println("Hello x", x)
		x++
	}

	// MARK: LOOPINGS (THIRD WAYS OF LOOPING)
	fmt.Println("\nTHIRD WAYS OF LOOPING: ")
	y := 0
	for {
		fmt.Println("Hello y", y)

		y++
		if y == 3 {
			break
		}
	}

	// MARK: LOOPING (NESTED LOOPING)
	fmt.Println("\n NESTED LOOPING: ")
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
