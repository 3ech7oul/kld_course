package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	var change = func(a *int, b *int) {
		aPt := *a
		bPt := *b

		*a = bPt
		*b = aPt
	}
	change(&a, &b)
	fmt.Println("a is", a)
	fmt.Println("b is", b)
}
