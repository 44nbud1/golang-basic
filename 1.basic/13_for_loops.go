package main

import (
	"fmt"
)

func main() {

	name := []string{"names", "nugi", "budi", "adam"}

	for n := range name {
		fmt.Println(n)
	}

	for i, n := range name {
		fmt.Println(i)
		fmt.Println(n)
	}
}
