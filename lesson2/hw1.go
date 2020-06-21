package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	var arNumb [4]int
	isUnig := true
	msg := "Yes"

	fmt.Fscan(os.Stdin, &num)
	numString := strconv.FormatInt(int64(num), 10)

	for i := 0; i < len(numString); i++ {
		el := string(numString[i])
		num, _ := strconv.Atoi(el)
		if i > 3 {
			fmt.Println("Will be compare only first 4 digits numbers")
			continue
		}

		if contains(arNumb, num) {
			isUnig = false
		}

		arNumb[i] = num
	}

	if !isUnig {
		msg = "No"
	}

	fmt.Println(msg)
}

func contains(a [4]int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
