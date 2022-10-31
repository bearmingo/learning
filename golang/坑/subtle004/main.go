package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 错误实例
func error_example1() {
	resp, err := http.Get("https://test.com")
	defer resp.Body.Close() // 这里可能崩溃。请求失败，resp可能为nil，不能读取Body
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	fmt.Println(string(body))
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func error_example2() {
	resp, err := http.Get("https://test.com")
	checkErr(err)

	defer resp.Body.Close() // 绝大多数情况下能够正常关闭，但在得到重定向的错误时，resp和err的值都不为空。

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	fmt.Println(string(body))
}

func main() {
	resp, err := http.Get("https://test.com")
	// 不为空时都需关闭
	if resp != nil {
		defer resp.Body.Close()
	}

	checkErr(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	fmt.Println(string(body))
}
