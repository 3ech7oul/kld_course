package main

import (
	"fmt"
	"os"
)

func main() {
	var input int
	var values []int

	for true {
		fmt.Fscan(os.Stdin, &input)

		if (len(values) == 0) && (input <= 5) {
			values = make([]int, input)
		}

		if (len(values) <= 0) && (input > 5) {
			fmt.Println("Must be less 5")
		}

		if cap(values) > 0 {
			values = append(values, input)
			fmt.Println("Last element:", values[len(values)-1])
		}

	}
}
