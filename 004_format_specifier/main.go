package main

import "fmt"

func main() {

	// Default format of the value --> %v
	fmt.Printf("%v", 42)        //42
	fmt.Printf("\n%v", "Hello") // Hello
	fmt.Printf("\n%v", 3.14)    //3.14

	// Type of the Value --> %T
	fmt.Printf("\n%T", 42)      // int
	fmt.Printf("\n%T", "Hello") // string
	fmt.Printf("\n%T", 3.14)    //float64

	// Decimal integer --> %d
	fmt.Printf("\n%d", 543) //543

	// Floating point decimal --> %f
	fmt.Printf("\n%f", 3.14) //3.140000

	// String --> %s
	fmt.Printf("\n%s", "Hello World!") // Hello World!

	// Quoted String --> %q
	fmt.Printf("\n\n%q", "Hello") // "Hello"

	// Boolean --> %t
	fmt.Printf("\n%t", true) //true

	// Character --> %c
	fmt.Printf("\n%c", 'A') // A

	// Pointer representation --> %p
	var x int
	fmt.Printf("\n%p", &x) //0xc000016120

	// Binary --> %b, Octal --> %o, Hexadecimal --> %x
	fmt.Printf("\n\n%b", 42) // 101010
	fmt.Printf("\n%o", 42)   // 52
	fmt.Printf("\n%x", 42)   //2a

	// Some more advance format specifiers

	// Display additional information such as field names in structs --> %+v

	type Person struct {
		Name string
		Age  int
	}

	p := Person{"Alice", 35}
	fmt.Printf("\n\n%+v", p) // {Name:Alice Age:35}

	// Go-syntax representation --> %#v
	var age int = 8
	fmt.Printf("\n\n%#v", age) // 8
	fmt.Printf("\n%#v", p)     // main.Person{Name:"Alice", Age:35}

}
