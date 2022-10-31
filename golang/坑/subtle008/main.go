package main

import "fmt"

func main() {
	code3()
}

// 错误的 recover 调用示例
func code1() {
	recover()         //
	panic("not good") // 发生 panic，主程序退出
	recover()         // 不会被执行
	println("ok")
}

// 正确的 recover 调用示例
func code2() {
	defer func() {
		fmt.Println("recovered: ", recover())
	}()

	panic("not good")
}

// 错误的调用示例
func code3() {
	defer func() {
		doRecover()
	}()

	panic("not good")
}

func doRecover() {
	fmt.Println("recoverd: ", recover())
}
