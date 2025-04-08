package trace1

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("Hello, world!")
	fmt.Println(runtime.NumCPU())
}
