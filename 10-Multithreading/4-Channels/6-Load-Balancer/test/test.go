package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for i := range c {
		fmt.Printf("worker %d, msg %d\n", id, i)
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan int)
	for i := 0; i < 100; i++ {
		go worker(i, c)
	}

	for i := 0; i < 1000; i++ {
		c <- i
	}
}
