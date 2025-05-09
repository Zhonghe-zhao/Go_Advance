package main

import (
	"fmt"
	"time"
)

/*
func receiver(ch <-chan int) {
	val := <-ch
	fmt.Println("接收到：", val)
}*/

// 因为发送操作会被阻塞，直到有数据接受它
//func sent(ch chan<- int) {
//	for a := 1; a <= 5; a++ {
//		ch <- a
//	}
//	close(ch)
//}
//
//func main() {
//	ch := make(chan int)
//	go sent(ch)
//	for val := range ch {
//		fmt.Println("goroutine:", val)
//	}
//}

//
//func main() {
//	ch := make(chan int) //无缓冲通道
//	go receiver(ch)
//	ch <- 12
//}

//初级生产者 消费者模型

func sent(ch chan<- int) {
	for i := 0; i < 5; i++ {
		now := time.Now().Second()
		ch <- now
		time.Sleep(time.Second)
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go sent(ch)
	for time := range ch {
		fmt.Println("number:", time)
	}
}
