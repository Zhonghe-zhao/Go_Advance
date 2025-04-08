package trace2

import "time"

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		println("Hello, World")
	}

}
