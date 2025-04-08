package main

import "fmt"

/*1.-------*/

// func outer() func() int {
// 	count := 0          // 外部变量
// 	return func() int { // 闭包函数
// 		count++ // 操作外部变量
// 		return count
// 	}
// }

/*2.-------*/
// func main() {
// 	msg := "hello" // 1. 初始化 msg 为 "hello"
// 	defer func() { // 2. 注册 defer 闭包（此时未执行）
// 		fmt.Println(msg) // 5. 执行闭包时，msg 已是 "world"
// 	}()
// 	msg = "world" // 3. 修改 msg 为 "world"
// } // 4. main 函数结束，触发 defer

/*3.------*/
// func main() {
// 	msg := "hello"
// 	defer func(msg string) { // 通过参数传递当前 msg 的值
// 		fmt.Println(msg) // 输出 "hello"
// 	}(msg) // 此时 msg 的值被复制
// 	msg = "world"
// }

/*4. ------*/

// func main() {
// 	for i := 0; i < 3; i++ { // i 从 0 递增到 2，循环结束后 i=3
// 		defer func() { // 注册 defer 闭包（此时不执行）
// 			fmt.Println(i) // 闭包内部引用的是外部的 i
// 		}()
// 	} // 循环结束后，i 的值为 3
// 	// 开始执行 defer（此时 i=3）

// }

/**5.------*/

func main() {
	for i := 0; i < 3; i++ {
		defer func(n int) { // n 是闭包的局部变量
			fmt.Println(n) // 输出 2, 1, 0（defer 是栈顺序）
		}(i) // 立即传入当前的 i 值
	}
}
