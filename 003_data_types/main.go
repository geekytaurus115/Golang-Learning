package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
		// *********  Data types in Golang  *********
		--> Basic Types (int, uint, float32, bool, string)
		--> Composite Types (array, map, struct)
		--> Reference Types (pointers, slices)
		--> Interface Types (interfaces)
	*/

	// Numeric Types
	// // Signed Integer (int, int8, int16, int32, int64)
	var age int = 25
	var temparature int32 = 32

	// // Unsigned Integers (uint, uint8, uint32, uint16, uint32, uint64)
	var score uint = 100
	var distance uint64 = 1500

	/*
			The main difference between different integer types in Go (like int, int8, int16, int32, int64, and their unsigned counterparts) is
			--> the range of values they can represent.
			--> The size of these integer types determines the maximum and minimum values they can store.

			### int and uint:
			--> These types are platform-dependent, and their size is determined by the underlying architecture of the system.
		    --> They are usually 32 or 64 bits on most systems.
			--> The int type is signed, while uint is unsigned.

			### int8 and uint8 (byte):
			--> int8 represents signed 8-bit integers, ranging from -128 to 127.
			--> uint8 represents unsigned 8-bit integers (byte), ranging from 0 to 255.

			### int16 and uint16:
		    --> int16 represents signed 16-bit integers, ranging from -32,768 to 32,767.
		    --> uint16 represents unsigned 16-bit integers, ranging from 0 to 65,535.

			### int32 and uint32:
			--> int32 represents signed 32-bit integers, ranging from -2,147,483,648 to 2,147,483,647.
		    --> uint32 represents unsigned 32-bit integers, ranging from 0 to 4,294,967,295.

			### int64 and uint64:
		    --> int64 represents signed 64-bit integers, ranging from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.
		    --> uint64 represents unsigned 64-bit integers, ranging from 0 to 18,446,744,073,709,551,615.

	*/
	var x int64
	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(x))

	var y uint
	fmt.Printf("Size of uint: %d bytes\n", unsafe.Sizeof(y))

	// Floating-Point
	var pi float32 = 3.1415
	var price float64 = 588.99

	// Boolean Types
	var isTrue bool = true
	var isFalse bool = false

	// String Types
	var myName string = "Geekytaurus"
	var message string = `Hi
	Hope you are well.`

	// Rune (character) Types
	var char rune = 'A'

	// Complex Number -- (complex64 and complex128)
	var complexNum complex64 = 5 + 7i

}
