package main

import "fmt"

func main() {
	x := "text"
	fmt.Println(&x)
	xBytes := []byte(x)
	fmt.Println(xBytes) // 打印出来的不是地址， 那这里怎么获取到 xBytes 的地址？
	//xBytes[0] = 'T' // 注意此时的T 是 rune 类型
	fmt.Println("len: ", len(xBytes), ", cap: ", cap(xBytes))
	xBytes = append(xBytes, []byte("123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")...)
	fmt.Println(&xBytes)
	x = string(xBytes)
	fmt.Println(x, &x) // text
}
