package main

import (
	"fmt"
	"github.com/karimarttila/go/lib/stringutil"
)

func main() {
	fmt.Println("Hello to my first Go demo!")
	fmt.Println("!oG ,olleH reversed is:")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
}
