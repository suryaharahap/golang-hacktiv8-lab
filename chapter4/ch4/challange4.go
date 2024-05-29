package main

import (
	"fmt"
	"sync"
)

func main() {
	// Menampung data slice dengan tipe data integer 1 sampai 10
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Membuat `waitgroup` untuk untuk melacak penyelesaian goroutine
	var wg sync.WaitGroup

	// Melooping slice
	for _, number := range numbers {
		// Menambah `waitgroup` untuk setiap goroutine
		wg.Add(1)
		go func(n int) {
			defer wg.Done() // Mengurangi `waitgroup` setiap goroutine selesai
			fmt.Println(n)
		}(number)
	}

	// Menunggu semua goroutine selesai
	wg.Wait()
}
