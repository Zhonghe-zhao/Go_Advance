package executor

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

type TaskFunc func(ctx context.Context) error

func ExecuteAll(numCPU int, tasks ...TaskFunc) error {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(len(tasks))

	if numCPU == 0 {
		numCPU = runtime.NumCPU()
	}
	fmt.Println("nuwCPU:", numCPU)
	queue := make(chan TaskFunc, numCPU)

	for i := 0; i < numCPU; i++ {
		go func() {
			for task := range queue {
				fmt.Println("get task")
				if err == nil {
					taskErr := task(ctx)
					if taskErr != nil {
						err = taskErr
						cancel()
					}
				}
				wg.Done()
			}
		}()
	}
	for _, task := range tasks {
		queue <- task
	}
	close(queue)

	wg.Wait()
	return err
}
