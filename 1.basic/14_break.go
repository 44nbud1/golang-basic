package main

import "fmt"

//break digunakan untuk menghentikan perulangan

func main() {

	var isEligible = []bool{
		true,
		true,
		true,
		false,
		true,
		true,
		true,
	}

	for i, v := range isEligible {

		// index 4 kebawah tidak adak ditampilkan
		fmt.Println("index: ", i)

		if !v {
			break
		}
	}
	fmt.Println(len(isEligible))
}
