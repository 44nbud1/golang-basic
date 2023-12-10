package main

import "fmt"

func main() {
	fmt.Println(add(1))
}

func add(i int) int {

	fmt.Println(i)

	if i >= 50 {
		return i
	}

	return i + add(i+1)
}
