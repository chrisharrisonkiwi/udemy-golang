package main

import (
	"fmt"
	myAlias "github.com/Udemy/02-package/icomefromnz"
	"github.com/Udemy/02-package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Println(stringutil.MyFullName)
	fmt.Println(myAlias.KiwiName)
}
