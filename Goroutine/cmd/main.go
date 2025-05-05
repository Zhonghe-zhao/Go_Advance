package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	i := 0
	for i <= 3 {
		fmt.Println("main:", i)
		i++
	}
	f("direct")

	go f("goroutine")

	go func(msg int) {
		fmt.Println(msg)
	}(i)

	time.Sleep(time.Second)
	fmt.Println("done")
}
