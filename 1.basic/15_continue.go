package main

import "fmt"

func main() {
	var isEligible = []bool{
		true,  // 0
		true,  // 1
		true,  // 2
		false, // 3
		true,  // 4
		true,  // 5
		true,  // 6
	}

	for i, v := range isEligible {

		if !v {
			continue
		}

		// index 3 tidak adak ditampilkan, namun perulangan tetap dilanjutkan
		fmt.Println("index: ", i)
	}
}
