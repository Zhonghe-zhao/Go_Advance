package executor

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestStopJobImmediately(t *testing.T) {
	tasks := make([]TaskFunc, 0, 100)
	for i := 0; i < 100; i++ {
		func(val int) {
			tasks = append(tasks, func(ctx context.Context) error {
				fmt.Println(val)
				if val == 51 {
					return errors.New("missing")
				}
				return nil
			})
		}(i)
	}

	err := ExecuteAll(0, tasks...)
	if err == nil {
		t.Error("missing error")
	}

}
