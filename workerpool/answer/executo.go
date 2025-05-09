package executor

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

type TaskFunc func(ctx context.Context) error

// Go 的可变参数 ...TaskFunc 到切片的转换
func ExecuteAll(numCPU int, tasks ...TaskFunc) error {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if numCPU == 0 {
		numCPU = runtime.NumCPU()
	}

	wg := sync.WaitGroup{}
	wg.Add(numCPU)

	//if numCPU == 0 {
	//	numCPU = runtime.NumCPU()
	//}

	fmt.Println("nuwCPU:", numCPU)
	//queue := make(chan TaskFunc, numCPU)

	queue := make(chan TaskFunc, len(tasks))

	for i := 0; i < numCPU; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case task, ok := <-queue:
					if !ok || ctx.Err() != nil {
						return
					}
					if e := task(ctx); e != nil {
						err = e
						cancel()
					}
				case <-ctx.Done():
					return
				}
			}
			//for task := range queue {
			//	fmt.Println("get task")
			//	if err == nil {
			//		taskErr := task(ctx)
			//		if taskErr != nil {
			//			err = taskErr
			//			cancel()
			//		}
			//	}
			//	wg.Done()
			//}
		}()
	}
	for _, task := range tasks {
		queue <- task
	}
	close(queue)

	wg.Wait()
	return err
}
