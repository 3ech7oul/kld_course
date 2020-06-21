package main

import (
	"fmt"
)

func main() {
	var strCount int
	var moreFive int
	str0 := "counter"
	str1 := "bull"
	str2 := "steal"
	str3 := "bear"
	strCount, moreFive = counter(str0, str1, str2, str3)

	fmt.Println(strCount, moreFive)

}

func counter(str ...string) (int, int) {
	var five int

	for _, v := range str {
		if len(v) >= 5 {
			five++
		}
	}

	return len(str), five
}
