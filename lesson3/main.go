package main

import (
	"fmt"
)

func changeValue(x int) {
	x = x + 100
}

func changeValuePtr(x *int) {
	*x = *x + 100
}

func main() {
	var x int = 4
	fmt.Println("Value", x)

	changeValue(x)
	fmt.Println("changeValue", x)

	changeValuePtr(&x)
	fmt.Println("changeValuePtr", x)

}
