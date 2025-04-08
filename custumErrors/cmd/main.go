package main

import (
	"fmt"
	"io"
	"net/http"
)

type urlError struct {
	url     string
	Err     error
	message string
}

func (e *urlError) Error() string {
	return fmt.Sprintf("%s,%s,%s", e.url, e.message, e.Err)
}

// func f(arg int) (int, error) {
// 	if arg == 42 {

// 		return -1, &argError{arg, "can't work with it"}
// 	}
// 	return arg + 3, nil
// }

// func main() {

// 	_, err := f(42)
// 	var ae *argError
// 	if errors.As(err, &ae) {
// 		fmt.Println(ae.arg)
// 		fmt.Println(ae.message)
// 	} else {
// 		fmt.Println("err doesn't match argError")
// 	}
// }

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	// 也就是断言会在这里判断返回的error是否是*urlError类型
	if err != nil {
		return "", &urlError{
			message: "获取网页失败",
			url:     url,
			Err:     err,
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

func main() {
	content, err := fetchURL("https://gobyexam2323ple.com/")
	if urlErr, ok := err.(*urlError); ok {
		fmt.Println("Error: ", urlErr)
	} else {
		fmt.Println("err:", err)
	}
	fmt.Println(content)

}
