package main

import (
	"fmt"
	"kld_course/lesson6/validate"
)

func main() {
	textRu := `ываываыв`
	textEn := `sdfdsfdsfsdf`

	fmt.Println(validate.IsString(textEn))
	fmt.Println(validate.IsString(textRu))
	fmt.Println(validate.IsNumber("23434"))
	fmt.Println(validate.IsNumber("2343sdfd4"))
	fmt.Println(validate.IsEmail("2343sdfd4"))
	fmt.Println(validate.IsEmail("2343sdfd4@mail.com"))

}
