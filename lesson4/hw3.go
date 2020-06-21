package main

import (
	"fmt"
)

func main() {
	a, b := 10, 4

	aPt := &a
	bPt := &b
	fmt.Println(*aPt * *bPt)

}
