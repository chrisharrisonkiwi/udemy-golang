package main

import "fmt"

func main() {
	fmt.Println(x) // nope. x has yet to be instantiated
	fmt.Println(y)
	x := 42
}

var y = 42
