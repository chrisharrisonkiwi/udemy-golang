package main

import (
	"fmt"
	"github.com/Udemy/04-scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
