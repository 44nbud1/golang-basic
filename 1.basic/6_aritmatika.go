package main

import "fmt"

func main() {

	fmt.Println(5 + 5) // penjumlahan = 10
	fmt.Println(7 - 5) // pengurangan = 2
	fmt.Println(5 * 5) // perkalian = 25
	fmt.Println(5 / 5) // pembagian = 1
	fmt.Println(5 % 4) // hasil bagi = 1

	//shortcut nya

	var a = 5
	// augmented assignment
	a += 5 // penjumlahan a = a + 5
	a -= 3 // pengurangan a = a - 3
	a *= 2 // perkalian a = a * 2
	a /= 2 // pembagian a = a / 2
	a %= 4 // hasil bagi a = a % 4

	// unary
	a++ // a = a + 1
	a-- // a = a - 1

}
