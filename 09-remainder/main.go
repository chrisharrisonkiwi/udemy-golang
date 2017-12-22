package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println(x)
	if x == 1 {
		fmt.Println("It's Odd")
	} else {
		fmt.Println("It's Even")
	}

	for i := 1; i <= 5; i++ {

		// note: parentheses are optional in Go. You could, if (i%2 == 1) { but why add code...
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}
