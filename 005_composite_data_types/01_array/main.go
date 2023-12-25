package main

import (
	"fmt"
	"reflect"
)

func main() {

	// ********* Arrays *********

	// Declaration without Initialization
	var arr [3]int // all elements initialized to 0
	fmt.Println(arr)

	// Declaration with Initialization
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)

	// Declaration with Implicit Size
	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr2)

	// Manual Initializations
	var arr3 [3]int
	arr3[0] = 1
	arr3[1] = 2
	arr3[2] = 3
	fmt.Println(arr3)

	// Partial Initializations
	arr4 := [4]int{1}
	fmt.Println(arr4) //[1 0 0 0]

	// ********** Operations on Array **********
	arr5 := [3]int{1, 2, 3}

	// Accessing element
	element := arr5[0]
	fmt.Println(element)

	// Modifying element
	arr5[0] = 42
	fmt.Println(arr5)

	// Iterating Over Elements
	for i, v := range arr5 {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// Length of Array
	fmt.Printf("Length of arr: %d\n", len(arr5))

	// Comparing Array
	isEqual := reflect.DeepEqual(arr1, arr5)
	fmt.Println(isEqual)

	// Multidimensions Array
	var matrix [3][3]int
	fmt.Println(matrix)

}
