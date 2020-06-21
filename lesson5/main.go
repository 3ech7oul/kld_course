package main

import "fmt"

type Machine struct {
	Signal, Env string
}

func (m Machine) beep() {
	fmt.Println(m.Signal)
}

func (m Machine) environment() {
	fmt.Println(m.Env)
}

var car Machine
var airplane Machine

func main() {
	car.Signal = "beep-beep"
	car.Env = "Ground"

	airplane.Signal = "ping"
	airplane.Env = "Air"
	car.beep()
	car.environment()
	airplane.beep()
	airplane.environment()

}
