package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Println(i, "FIZZ BUZZ")
		} else if i % 3 == 0 {
			fmt.Println(i, "FIZZ")
		} else if i % 5 == 0 {
			fmt.Println(i, "BUZZ")
		} else {
			fmt.Println(i)
		}
	}

	// alternatively (more complex but interesting)

	for i := 0; i <= 100; i++ {

		fmt.Print(i)

		switch {
		case i % 3 == 0 :
			fmt.Print(" FIZZ")
			if i % 15 != 0 {
				break
			}
			fallthrough
		case i % 5 == 0 :
			fmt.Print(" BUZZ")
		}

		fmt.Println()
	}
}
