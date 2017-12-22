package main

import (
	"fmt"
)

func main() {

	a := 43 // store a value

	fmt.Println(a)  // print the value
	fmt.Println(&a) // print the stored memory address

	var b *int = &a // create a reference aka: pointer to the memory address
	// (note the * ref type in the type declaration - remember: types are optional in Go)

	fmt.Println(b)  // print the memory address
	fmt.Println(*b) // *var - * give me the value stored within this reference

	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int
}
