package main

import "fmt"

func main() {
	variadicParam()
	variadicParam("aan", "budi", "setiawan")
}

func variadicParam(names ...string) {

	for _, v := range names {
		fmt.Println("name: ", v)
	}
}
