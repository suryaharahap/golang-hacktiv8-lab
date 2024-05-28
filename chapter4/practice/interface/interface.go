package main

import "fmt"

func main() {

	/*
		Empty interface merupakan suatu tipe data
		yang dapat menerima segala tipe data pada bahasa Go.

		Variabel dengan tipe data interface dapat diberikan nilai
		dengan tipe data yang berbeda-beda
	*/
	var randomValues interface{}

	_ = randomValues
	randomValues = "Jalan sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Airel", "Nadia"}

	fmt.Println(randomValues)

	/*
		MARK: Empty interface (type assertion)
	*/

	var v interface{}
	v = 20
	// v = v * 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	fmt.Println(v)

	/*
		MARK: Empty interface (map & slice of empty interface)
	*/

	// Variabel `rs` memiliki tipe data slice of empty interface dan memiliki nilai dengan tipe data yang berbeda-beda.
	rs := []interface{}{1, "Noah", true, 2, "Dinda", true}

	// `rm` memiliki tipe data map yang memiliki key dengan tipe data string dan value dengan tipe data empty interface.
	rm := map[string]interface{}{
		"Name":   "Noah",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm

	fmt.Println(rs...)
	fmt.Println(rm)

}
