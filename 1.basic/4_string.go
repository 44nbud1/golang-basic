package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "aan budi setiawan"

	// 1. length of name
	fmt.Println("panjang dari namaku ", len(name))

	// 2. get character, in bit
	fmt.Printf("%v\n", name[0])

	// 3. hasPrefix
	fmt.Println(strings.HasPrefix(name, "aan"))
}
