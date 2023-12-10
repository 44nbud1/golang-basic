package main

import "fmt"

func main() {
	fHello := hello

	fmt.Println(fHello("Aan"))
}

func hello(s string) string {
	return "hello " + s
}
