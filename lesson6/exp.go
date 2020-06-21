package main

import (
	"errors"
	"fmt"
	"kld_course/lesson6/validate"
)

func main() {
	var str string
	for true {
		fmt.Scan(&str)
		_, err := handelInput(str)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}

	}
}

func handelInput(s string) (string, error) {
	if !validate.IsString(s) {
		return s, errors.New("Something went wrong! Check input value. Allowed only letters.")
	}
	return s, nil
}
