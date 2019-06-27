package main

import "fmt"

func defer_Call_1() {
	defer func() {
		fmt.Println("打印前=--=-=-=-=-=-")
	}()
	defer func() {
		fmt.Println("打印中=--=-=-=-=-=-")
	}()
	defer func() {
		fmt.Println("打印后=--=-=-=-=-=-")
	}()

	panic("出发异常 call1")

	// 结果
	// 打印后=--=-=-=-=-=-
	// 打印中=--=-=-=-=-=-
	// 打印前=--=-=-=-=-=-
	// panic: 出发异常 call1
}

func defer_Call_2() {
	defer func() {
		fmt.Println("打印前=--=-=-=-=-=-")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(" recover2 接受到异常 err = ", err)
		}
		fmt.Println("打印中=--=-=-=-=-=-")
	}()
	defer func() {
		fmt.Println("打印后=--=-=-=-=-=-")
	}()

	panic("触发异常")

	// 结果. recover捕获异常,异常在 打印中的那个 defer中打印出来.
	// 	打印后=--=-=-=-=-=-
	//  recover2 接受到异常 err =  触发异常
	// 打印中=--=-=-=-=-=-
	// 打印前=--=-=-=-=-=-
}

func defer_Call_3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(" recover3 前 接受到异常 err = ", err)
		}
		fmt.Println("打印前=--=-=-=-=-=-")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(" recover3 种 接受到异常 err = ", err)
		}
		fmt.Println("打印中=--=-=-=-=-=-")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(" recover3 后 接受到异常 err = ", err)
		}
		fmt.Println("打印后=--=-=-=-=-=-")
	}()

	panic("触发异常")

	// 如果同时有多个defer，那么异常会被最近的recover()捕获并正常处理。
	// recover3 后 接受到异常 err =  触发异常
	// 打印后=--=-=-=-=-=-
	// 打印中=--=-=-=-=-=-
	// 打印前=--=-=-=-=-=-
}

func main() {
	// defer_Call_1()
	// defer_Call_2()
	defer_Call_3()
}
