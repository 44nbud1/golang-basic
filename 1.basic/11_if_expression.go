package main

import "fmt"

func main() {

	//i := 0
	//i := 2
	i := 3

	if i == 2 {
		fmt.Println("dua")
	} else if i == 3 {
		fmt.Println("tiga")
	} else {
		fmt.Println("nol")
	}

	var name = "budi"
	if l := len(name); l == 4 {
		fmt.Println("empat")
	} else {
		fmt.Println("nol")
	}
}
