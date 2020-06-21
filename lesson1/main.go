package main

import (
	"fmt"
	"strconv"
)

func main() {

	var first, last int
	num := 123456
	numString := strconv.FormatInt(int64(num), 10)

	for i := 0; i < len(numString); i++ {
		el := string(numString[i])
		num, _ := strconv.Atoi(el)

		if i <= 2 {
			first = first + num
		}

		if i > 2 {
			last = last + num
		}

	}

	fmt.Println(first)
	fmt.Println(last)

}
