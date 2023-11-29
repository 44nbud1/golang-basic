package main

import "fmt"

func main() {

	// array ada kumpulan data yang memiliki tipe data yang sama
	var arrData [5]string

	arrData[0] = "data1"
	arrData[1] = "data2"
	arrData[2] = "data3"
	arrData[3] = "data4"

	fmt.Println(arrData)
	fmt.Println(len(arrData))

	// selain itu untuk membuat array bisa menggunakan variadic (...) tp data nya harus lgsung di init
	var data2 = [...]string{
		"data1",
		"data2",
		"data3",
		"data4",
		"data5",
	}
	fmt.Println(data2)
}
