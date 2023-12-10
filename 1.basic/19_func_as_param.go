package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	a := func(s string) {
		fmt.Println("hello ", s)
	}

	s := a

	s("Aan")

	helper()
}

func printNameToJson(errFn func(stringName string) error) {

	var err error

	if err = errFn("name"); err != nil {
		fmt.Println("error, ", err)
		return
	}

	fmt.Println("Ok, ", err)
}

func helper() {

	a := func(s string) error {
		if bytes, err := json.Marshal(s); err != nil {
			return err
		} else {
			fmt.Println(fmt.Sprint(bytes))
		}

		return nil
	}

	printNameToJson(a)
	printNameToJson(printJson())
}

type generate func(s string) error

func printJson() generate {
	return func(s string) error {
		if bytes, err := json.Marshal(s); err != nil {
			return err
		} else {
			fmt.Println(fmt.Sprint(bytes))
		}

		return nil
	}
}
