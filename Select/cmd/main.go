package main

import (
	"fmt"
)

// func main() {

// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		c1 <- "one"
// 	}()
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		c2 <- "two"
// 	}()
// 	//select 允许您等待多个通道作
// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println("received", msg1)
// 		case msg2 := <-c2:
// 			fmt.Println("received", msg2)
// 		default:
// 			time.Sleep(10000 * time.Millisecond) // 避免 CPU 忙等待
// 			fmt.Println("no message received")
// 		}

// 	}
// }

/*timeout超时机制*/

// func main() {

// 	c1 := make(chan string, 1)
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		c1 <- "result 1"
// 	}()

// 	select {
// 	case res := <-c1:
// 		fmt.Println(res)
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("timeout 1")
// 	}

// 	c2 := make(chan string, 1)
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		c2 <- "result 2"
// 	}()
// 	select {
// 	case res := <-c2:
// 		fmt.Println(res)
// 	case <-time.After(3 * time.Second):
// 		fmt.Println("timeout 2")
// 	}
// }

/*Non-Blocking Channel Operations*/

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	signals <- true
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
