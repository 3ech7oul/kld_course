package main

import "fmt"

func main() {

	periodicTable := make(map[int]string)

	periodicTable[1] = "H"
	periodicTable[2] = "He"
	periodicTable[3] = "Le"
	periodicTable[4] = "Be"
	periodicTable[8] = "O"

	for key, value := range periodicTable {
		if "O" == value {
			fmt.Println(key)
		}
	}

}
