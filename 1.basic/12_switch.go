package main

import "fmt"

func main() {
	expression := 10

	switch expression {
	case 1:
		fmt.Println("Januari")
	case 2:
		fmt.Println("Februari")
	case 3:
		fmt.Println("Maret")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("Mei")
	case 6:
		fmt.Println("Juni")
	case 7:
		fmt.Println("Juli")
	case 8:
		fmt.Println("Agustus")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("Oktober")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("Desember")
	default:
		fmt.Println("Tidak ditemukan")
	}

	switch l := len("aan"); l == 4 {
	case true:
		fmt.Println("Benar")
	default:
		fmt.Println("Salah")
	}

	// tidak menggunakan switch expression

	l := len("name")

	switch {
	case l == 1:
		fmt.Println("satu")
	case l == 2:
		fmt.Println("dua")
	case l == 3:
		fmt.Println("tiga")
	case l == 4:
		fmt.Println("empat")
	case l == 5:
		fmt.Println("lima")
	default:
		fmt.Println("tidak ada")
	}
}
