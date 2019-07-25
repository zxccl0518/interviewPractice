package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		fmt.Println("hello world")
	}()

	go func() {
		for {
			// fmt.Println("go func 2")
			// time.Sleep(time.Second)
		}
	}()

	fmt.Println("结束---")
}
