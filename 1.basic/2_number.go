package main

import "fmt"

func main() {
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	i8 = 8
	i16 = 16
	i32 = 32
	i64 = 64

	fmt.Println(i8)
	fmt.Println(i16)
	fmt.Println(i32)
	fmt.Println(i64)

	// tidak boleh minus
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	ui8 = 8
	ui16 = 16
	ui32 = 32
	ui64 = 64

	fmt.Println(ui8)
	fmt.Println(ui16)
	fmt.Println(ui32)
	fmt.Println(ui64)

	// alias byte == uint8
	var aliasUi8 byte
	aliasUi8 = 89
	fmt.Println(aliasUi8)
}
