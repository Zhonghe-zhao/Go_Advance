package main

import (
	"fmt"
)

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Println("worker", id, "started  job", j)
// 		time.Sleep(time.Second)
// 		fmt.Println("worker", id, "finished job", j)
// 		results <- j * 2
// 	}
// }

// func main() {

// 	const numJobs = 5
// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	for w := 1; w <= 3; w++ {
// 		go worker(w, jobs, results)
// 	}

// 	for j := 1; j <= numJobs; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)

// 	for a := 1; a <= numJobs; a++ {
// 		<-results
// 	}
// }

func worker(id int, jobs <-chan int, results chan<- int /*wg *sync.WaitGroup*/) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)

		fmt.Println("worker", id, "finished job", j)
		//wg.Done()
		results <- j * 2
	}
}

func main() {
	//var wg sync.WaitGroup
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for i := 0; i < 2; i++ {
		go worker(i, jobs, results /*&wg*/)
	}

	for j := 1; j <= 5; j++ {
		//wg.Add(1)
		jobs <- j
	}
	close(jobs)
	// 如果不等待woker处理完毕，main函数就结束了，worker还在继续处理任务，导致结果不对
	//wg.Wait()

	//2.等待多个 goroutine
	for i := 0; i < 5; i++ {
		fmt.Println("result:", <-results)
	}

}
