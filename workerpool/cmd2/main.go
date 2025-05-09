package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i < 11; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)

}

func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("val:", <-ch)
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	// 28-31行的数据顺序也会导致不同的结果 对于没有 go 启动的生产者来说，生产者执行的是一个阻塞操作，它会依次执行以下步骤
	producer(ch)
	for i := 1; i <= 2; i++ {
		go consumer(i, ch, &wg)
	}

	wg.Wait()
}
