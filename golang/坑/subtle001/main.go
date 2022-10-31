package main

import "fmt"

type data struct {
	name string
}

type printer interface {
	print()
	print2()
}

func (p *data) print() {
	fmt.Println("print: ", p.name)
}

func (p data) print2() {
	fmt.Println("print2:", p.name)
}

func main() {
	d1 := data{"one"}
	d1.print()

	var in printer = &data{"two"}
	in.print()

	m := map[string]data{
		"x": {"three"},
	}
	// m["x"].print()  // <- map的data是不可寻址的
	// var a = &m["x"] //
	m["x"].print2()
}
