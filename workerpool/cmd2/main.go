package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		val := id*10 + 1
		fmt.Println("生产者:", id, "生产:", val)
		ch <- val
		time.Sleep(time.Second)
	}
	//close(ch) 不能让每个生产者都能关闭通道 而是要统一关闭

}

func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("消费者: ", id, "消费:", val)
	}
}

func main() {
	ch := make(chan int)
	var wgPro sync.WaitGroup
	var wgCon sync.WaitGroup

	wgPro.Add(2)
	// 28-31行的数据顺序也会导致不同的结果 对于没有 go 启动的生产者来说，生产者执行的是一个阻塞操作，它会依次执行以下步骤
	for a := 1; a <= 2; a++ {
		go producer(a, ch, &wgPro)
	}

	wgCon.Add(3)
	for i := 1; i <= 3; i++ {
		go consumer(i, ch, &wgCon)
	}

	//统一处理通道关闭
	go func() {
		wgPro.Wait()
		close(ch)
	}()

	wgCon.Wait() //等待所有消费者处理完
}
