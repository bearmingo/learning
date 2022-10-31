package main

import "fmt"

type data struct {
	name string
}

func main() {
	m := map[string]data{
		"x": {"Test"},
	}

	// m["x"].name = "Test2"	// <- map的值为struct时，struct的内的值不可修改
	fmt.Printf(m["x"].name)

	s := []data{{"Test"}}
	s[0].name = "Test2" // slice的元素是可以寻址
}
