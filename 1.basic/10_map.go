package main

import (
	"fmt"
)

func main() {

	var ms = map[string]string{
		"keys": "values",
	}
	fmt.Println(ms)
	ms["key2"] = "val2"
	fmt.Println(ms)

	m := make(map[string]string)
	m["key"] = "value"
	fmt.Println(m)

	for k, v := range ms {
		fmt.Println("key: ", k)
		fmt.Println("value: ", v)
	}

	dup := make(map[int]interface{})

	dup[1] = 10

	fmt.Println(dup[1])
	fmt.Println(dup[2])

	duplicate([]int{
		1, 2, 3, 4, 5, 3, 2, 3, 23, 42, 1, 32, 2, 2, 4, 1, 3, 4,
	})

}

func duplicate(v []int) {
	fmt.Println("DUPLICATE")
	dup := make(map[int]interface{})
	for _, val := range v {
		count := dup[val]
		if count == nil {
			dup[val] = 1
		} else {
			cs := count.(int) + 1
			dup[val] = cs
		}
	}

	delete(dup, 1)
	for k, val := range dup {
		if val.(int) > 1 {
			fmt.Println("duplicate: ", k, " value ", val)
		}
	}
}
