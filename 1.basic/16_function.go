package main

import "fmt"

func main() {

	voidFunc()
	voidFuncParam("Hello world")
	fmt.Println(returnFunc())
	fmt.Println(returnFuncParam("World"))
	fmt.Println(returnMultipleVal("Hello"))

	//tidak menggunakan value
	hello, _ := returnMultipleVal("Hello")
	fmt.Println(hello)
	fmt.Println(returnNameParam())

}

func voidFunc() {
	fmt.Println("Hello world")
}

func voidFuncParam(hello string) {
	fmt.Println(hello)
}

func returnFunc() string {
	return "Hello world"
}

func returnFuncParam(name string) string {
	return "hello " + name
}

func returnMultipleVal(hello string) (string, string) {
	return hello, "aan"
}

func returnNameParam() (hello, name string) {

	hello = "hello"
	name = "Aan"

	return hello, name
}
