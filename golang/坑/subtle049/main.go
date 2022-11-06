package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		for m := range ch {
			fmt.Println("Processed: ", m)
			time.Sleep(1 * time.Second)
		}
	}()

	ch <- "cmd 1"
	fmt.Println("send cmd 1")
	ch <- "cmd 2"
	fmt.Println("send cmd 2")
}
