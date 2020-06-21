package main

import (
	"fmt"
)

func main() {

	numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	var odd, even int

	for _, v := range numbers {
		if v%2 == 0 {
			even++

		} else {
			odd++

		}
	}

	fmt.Println("Even count", even)
	fmt.Println("Odd count:", odd)

}
