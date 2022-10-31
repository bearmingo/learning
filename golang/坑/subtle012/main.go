package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) // 3 3 [1 2 3]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) // 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	}

	// 此时的 s1 与 s2 是指向同一个底层数组的
	fmt.Println(s1) // [1 22 23]
	fmt.Println(s2) // [22 23]

	s2 = append(s2, 4) // 向容量为 2 的 s2 中再追加元素，此时将分配新数组来存

	for i := range s2 {
		s2[i] += 10
	}

	fmt.Println(s1) // [1 22 23]
	fmt.Println(s2) // [32 33 14]
}
