package main

import "fmt"

func zero(z int) {
	z = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5

	// this makes sense.
	// x passes its value into z within the zero() function, z is altered
	// zero() never returns the new value and x is never altered.
}
