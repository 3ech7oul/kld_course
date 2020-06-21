package main

import (
	"fmt"
	"os"
)

func main() {
	var input int
	var msg string

	for true {
		fmt.Fscan(os.Stdin, &input)

		switch {
		case input > 0:
			msg = "greater 0 "
		case input < 0:
			msg = "less 0 "
		case input == 0:
			msg = "equal 0 "
		}

		fmt.Println("Number is", input, msg)
	}

}
