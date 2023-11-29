package main

import "fmt"

func main() {

	// slice adalah potongan dari array
	var bulan = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	// slice mendapatkan nilai dari x index ke index terakhir
	fmt.Println(bulan[2:])

	// slice mendapatkan nilai dari index pertama ke index x
	fmt.Println(bulan[:5])

	// slice mendapatkan nilai dari index x ke index x
	fmt.Println(bulan[2:5])

	// merubah array ke slice
	fmt.Println(bulan[:])

	// deklarasi ini jika tidak diisi datanya akan mengembalikan nil
	var a []string

	if a == nil {
		a = make([]string, 0)
	}

	// another
	data := make([]int, 0)

	for i := 0; i < 10; i++ {
		data = append(data, i)
		a = append(a, fmt.Sprint(i))
	}

	fmt.Println(data)
	fmt.Println(a)
}
