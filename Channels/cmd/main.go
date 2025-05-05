package main

import "fmt"

/*无缓冲通道*/

// func main() {

// 	messages := make(chan string)

// 	go func() { messages <- "ping" }()
// 	//我们在程序结束时等待 “ping” 消息，而无需使用任何其他同步机制。
// 	msg := <-messages
// 	fmt.Println(msg)
// }

/*缓冲通道*/

// func main() {

// 	messages := make(chan string, 2)

// 	messages <- "buffered"
// 	messages <- "channel"

// 	fmt.Println(<-messages)
// 	fmt.Println(<-messages)
// 	//fmt.Println(<-messages)

// }

/*同步通道*/

// func worker(done chan bool) {
// 	fmt.Print("working...")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")

// 	done <- true
// }

// func main() {

// 	done := make(chan bool, 1)
// 	go worker(done)

// 	// 关键
// 	<-done
// }

/*通道方向*/

// func ping(pings chan<- string, msg string) {
// 	pings <- msg
// }

// func pong(pings <-chan string, pongs chan<- string) {
// 	msg := <-pings
// 	pongs <- msg
// }

// func main() {
// 	pings := make(chan string, 1)
// 	pongs := make(chan string, 1)
// 	ping(pings, "passed message")
// 	pong(pings, pongs)
// 	//fmt.Println(<-pongs)
// }

/*对已经关闭的channel进行写操作*/

// func main() {
// 	c := make(chan int, 3)
// 	close(c)
// 	c <- 1
// }

/*对已经关闭的channel进行读操作*/

// func main() {
// 	c1 := make(chan int, 3)
// 	c1 <- 1
// 	close(c1)
// 	num, ok := <-c1
// 	fmt.Printf("读chan的协程结束, num=%v, ok=%v\n", num, ok)
// 	num1, ok1 := <-c1
// 	fmt.Printf("读chan的协程结束, num1=%v, ok1=%v\n", num1, ok1)
// 	num2, ok2 := <-c1
// 	fmt.Printf("读chan的协程结束, num2=%v, ok2=%v\n", num2, ok2)
// }

/*未初始化chan*/

func main() {
	c := make(chan int) // 初始化 channel

	go func() {
		c <- 1
	}()

	fmt.Println(<-c)
}
