package main

import (
	"fmt"
	"os"
)

func main() {
	months := make(map[int]string)
	var input int
	fmt.Fscan(os.Stdin, &input)

	months[1] = "Январь"
	months[2] = "Февраль"
	months[3] = "Март"
	months[4] = "Апрель"
	months[5] = "Май"
	months[6] = "Июнь"
	months[7] = "Июль"
	months[8] = "Август"
	months[9] = "Сентябрь"
	months[10] = "Октябрь"
	months[11] = "Ноябрь"
	months[12] = "Декабрь"

	if input > 12 {
		fmt.Println("Введите число меньше или равное 12")
		return
	}

	fmt.Println(months[input])

}
