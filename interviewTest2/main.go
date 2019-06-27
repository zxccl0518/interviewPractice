package main

import (
	"fmt"
	"time"
	"unsafe"
)

var url = "https://www.jianshu.com/p/e1ed301a3599"

func main() {
	// test1()

	test4()
	time.Sleep(3 * time.Second)
}

// 测试1
// 测试结果, 是变量不携带任何信息, 且值是0, 那么变量的地址是 相同的.
func test1() {
	// 变量类型不携带任何信息 且0值, 地址是否相同.?
	s1 := struct{}{}
	s2 := struct{}{}
	d := [0]int{}
	var a map[int]int
	var sli = []int{}

	if unsafe.Pointer(&s1) == unsafe.Pointer(&d) {
		fmt.Println("地址 相同. ---- ")
	} else {
		fmt.Println("地址不相同. xxx --- xxx")
	}

	fmt.Printf("s1 address = %p\n", &s1)
	fmt.Printf("s2 address = %p\n", &s2)
	fmt.Printf("d address = %p\n", &d)
	fmt.Printf("a address = %p\n", &a)
	fmt.Printf("sli address = %p\n", &sli)
}

// 测试2
// new 初始化T类型的零值. 返回的是指针.
// make 初始化T类型,返回T类型.

// 测试3
// 当变量或者对象在方法中分配后,其指针被返回或者被全局引用.(这样就会对其他的过程或者线程所引用,)这种现象称作指针(或者引用)的逃逸.

// 测试4
// 隐式赋值,查看下面输出的结果是什么
// 通过这个 隐式赋值,学习了,最后return 那个位置是直接吧1 赋值给了ret
// 其实这个打印的结果还是看 谁 执行快,若第一个协程在return之前打印了 那么结果为0 第二个协程在return 之后执行了, 打印的结果就为1
func test4() (ret int) {
	i := 0
	ret = 0
	for i < 3 {
		go func() {
			// 如果延时先执行, 结果是1 ,1 ,1, 因为打印的时候, ret 已经变成了1
			fmt.Println("get value = ", ret)
			time.Sleep(time.Second)

			// 若是延时在打印之后, 则先打印,后延时,结果就是 0, 0, 0, 因为打印的时候,ret 值还没有改变.
			// time.Sleep(time.Second)
			// fmt.Println("get value = ", ret)
		}()

		fmt.Println("i = ", i)
		i++
	}

	return 1
}
