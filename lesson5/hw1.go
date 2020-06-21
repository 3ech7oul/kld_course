package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(i int, c chan int) {
	fmt.Println(i)
	c <- i

}

func reciver(c chan int) {

	data := make([]int, 5)
	i := <-c
	data = append(data, i)
	time.Sleep(time.Second)
	fmt.Println(i)

}

func main() {
	var c chan int = make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	for i := 0; i <= 7; i++ {
		go sender(i, c)
	}

	go reciver(c)

}
