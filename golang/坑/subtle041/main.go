package main

import "fmt"

func main() {
	data := "Aé\xfe\x02\xff\x04"
	// data := "我的"
	fmt.Println(data)
	for i, v := range data {
		fmt.Printf("%d: %#x, %d\n", i, v, v)
	}

	for i, v := range []byte(data) {
		fmt.Printf("%d: %#x\n", i, v)
	}
}
