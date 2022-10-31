package main

import "fmt"

func main() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil) // <nil> true
	fmt.Println(in, in == nil)     // <nil> true

	in = data
	fmt.Println(in, in == nil) // <nil> false 	// 这时的in的类型已经不为nil

	var data2 = in.(*byte)
	fmt.Println(data2, data2 == nil)

	error_example()
}

// 容易错误的地方，返回值为interface{}时
func error_example() {
	doIt := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		}

		return result
	}

	if res := doIt(-1); res != nil {
		fmt.Println("Good result: ", res) // Good result: <nil>
		fmt.Printf("%T\n", res)           // *struct{} 	// res的类型不为nil, res的值为nil
		fmt.Printf("%v\n", res)           // <nil>
	}
}

// 上一个错误实例的正确版本
func right_example() {
	doIt := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		} else {
			return nil // 直接返回nil，而不是result
		}
		return result
	}

	if res := doIt(-1); res != nil {
		fmt.Println("Good result: ", res)
	} else {
		fmt.Println("Bad result: ", res)
	}
}
