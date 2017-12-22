package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	// I get it, but im not sure the benefit as yet.
	// zero could simply return a value
	// this is prob a performance tool, that or my brain needs to break free of PHP more
}
