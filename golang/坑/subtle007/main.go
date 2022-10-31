package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type data struct {
	num    int
	checks [10]func() bool   // 无法比较
	doIt   func() bool       // 无法比较
	m      map[string]string // 无法比较
	bytes  []byte            // 无法比较
}

func main() {
	code2()
}

func code1() {
	v1 := data{}
	v2 := data{}

	// fmt.Println("v1 == v2: ", v1 == v2) 无法编译

	fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2)) // v1 == v2: true

	m1 := map[string]string{"1": "a", "2": "b"}
	m2 := map[string]string{"2": "b", "1": "a"}
	fmt.Println("m1 == m2: ", reflect.DeepEqual(m1, m2)) // m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}

	// 注意两个slice相等，值和顺序必须一致
	fmt.Println("s1 == s2: ", reflect.DeepEqual(s1, s2)) // s1 == s2:  true
}

func code2() {
	var str = "one"
	var in interface{} = "one"
	fmt.Println("str == in: ", reflect.DeepEqual(str, in)) // str == in: true

	v1 := []string{"one", "two"}
	v2 := []string{"two", "one"}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2)) // v1 == v2: false

	data := map[string]interface{}{
		"code":  200,
		"value": []string{"one", "two"},
	}

	encoded, _ := json.Marshal(data)
	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)
	fmt.Println("data == decoded: ", reflect.DeepEqual(data, decoded)) // data == decoded:
}

func code3() {
	var b1 []byte = nil
	b2 := []byte{}

	// b1 与 b2 长度相等、有相同的字节序
	// nil 与 slice 在字节上是相同的
	fmt.Println("b1 == b2: ", bytes.Equal(b1, b2)) // b1 == b2: true
}
