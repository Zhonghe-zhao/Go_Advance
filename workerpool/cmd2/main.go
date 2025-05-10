package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func producer(id int, ch chan<- int, wg *sync.WaitGroup, total *int32) {
	defer wg.Done()
	for {
		current := atomic.LoadInt32(total)
		if current >= 100 {
			return
		}
		newVal := atomic.AddInt32(total, 1)
		if newVal > 100 {
			return
		}
		fmt.Println("生产者:", id, "生产:", newVal)
		ch <- int(newVal)

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
	var counter int32
	ch := make(chan int)
	var wgPro sync.WaitGroup
	var wgCon sync.WaitGroup

	wgPro.Add(3)
	// 28-31行的数据顺序也会导致不同的结果 对于没有 go 启动的生产者来说，生产者执行的是一个阻塞操作，它会依次执行以下步骤
	for a := 1; a <= 3; a++ {
		go producer(a, ch, &wgPro, &counter)
	}

	//统一处理通道关闭
	go func() {
		wgPro.Wait()
		close(ch)
	}()

	wgCon.Add(5)
	for i := 1; i <= 5; i++ {
		go consumer(i, ch, &wgCon)
	}

	wgCon.Wait() //等待所有消费者处理完
}
