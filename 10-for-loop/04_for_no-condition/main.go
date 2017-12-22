package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
	}

	// I'm not sure why you would want an infinite loop in Go.
	// Possibly a listener service?
}
