package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

var test = 10 << 5 // this is like saying: (10 * 2) ^ 5 or 10 * 2, 5 times

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t%d\n", KB, KB)
	fmt.Printf("%b\t%d\n", MB, MB)
	fmt.Printf("%b\t%d\n", GB, GB)
	fmt.Printf("%b\t%d\n", TB, TB)

	fmt.Printf("%d\n", test)
}
