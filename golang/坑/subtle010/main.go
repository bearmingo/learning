package main

import "fmt"

func main() {
	code2()
}

func get() []byte {
	raw := make([]byte, 100000)
	fmt.Println(len(raw), cap(raw), &raw[0])

	return raw[:3]
}

func code1() {
	data := get() // data的容量还是10000，
	fmt.Println(len(data), cap(data), &data[0])
}

func get2() (res []byte) {
	raw := make([]byte, 100000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	res = make([]byte, 3)
	copy(res, raw[:3])
	return
}

func code2() {
	data := get2()
	fmt.Println(len(data), cap(data), &data[0])
}
