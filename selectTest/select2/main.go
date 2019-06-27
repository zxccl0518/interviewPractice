package main

import (
	"fmt"
	"sync"
)

func doFunc(ch chan interface{}, done chan struct{}, wg *sync.WaitGroup, workID int) {
	fmt.Println("协程 启动, 编号为 = ", workID)
	defer wg.Done()

	for {
		select {
		case v := <-ch:
			fmt.Printf("wordID = %v,  v = %v\n", workID, v)
		case <-done:
			fmt.Printf("编号 wordID =%v, 线程 结束.\n", workID)
			return
		}
	}
}

func main() {
	chanint := make(chan interface{})
	chanDone := make(chan struct{})
	wg := sync.WaitGroup{}

	workCount := 2
	for i := 0; i < workCount; i++ {
		wg.Add(1)
		go doFunc(chanint, chanDone, &wg, i)
	}
	for i := 0; i < workCount; i++ {
		chanint <- i
	}

	close(chanDone)
	wg.Wait()
	close(chanint)
	fmt.Println("循环结束.---------")
}
