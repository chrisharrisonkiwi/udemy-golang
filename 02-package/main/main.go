package main

import (
	"fmt"
	"github.com/Udemy/02-package/stringutil"
	myAlias "github.com/Udemy/02-package/icomefromnz"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Println(stringutil.MyFullName)
	fmt.Println(myAlias.KiwiName)
}
