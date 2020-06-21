package main

import (
	"fmt"
	"os"
)

var maxAttemp int = 7

func main() {
	arInp := make([]int, maxAttemp)
	var input int
	var atemp int

	for true {
		atemp++
		var res int
		fmt.Fscan(os.Stdin, &input)
		res, arInp = biggestNum(input, arInp, atemp)

		if atemp == maxAttemp {
			fmt.Println("Max", res)
			return
		}
	}
}

func biggestNum(num int, arInp []int, attemp int) (int, []int) {
	arInp = append(arInp, num)

	if attemp == maxAttemp {
		var prev int
		for _, v := range arInp {
			if v > prev {
				prev = v
			}
		}
		num = prev
	}

	return num, arInp
}
