package main

import (
	"fmt"
	"sync"
)

// gitbub 夜读,   面试汇总1
var URL = "https://reading.developerlearning.cn/interview/articles/interview_analysis_1/"

func main() {
	/*
		// --------------------------------------------------------------------
		// 测试第一个考点
		t := threadSafeSet{
			s: []interface{}{"1", "2", "3"},
		}

		// v := t.Iter_Wrong()
		// 这个测试发现的问题就是 创建channel 的时候 channel是无缓冲的channel, 写入一个数据后, 没有外部去菜从channel 取数据,就会造成阻塞, 所以循环阻塞,
		// <======================>
		v := t.Iter_Right()
		// <=======================>
		time.Sleep(1)
		fmt.Printf("%s <> %v \n", "ch", v)
		time.Sleep(time.Second)
		// --------------------------------------------------------------------
	*/

	/*
		// 考察的是defer的使用, defer 是在return 上一步执行.
		// 还有一个考察点是 返回值 若已经声明变量了.那么会被初始化为零值, 且作用于为这个函数.若没有声明返回值.则没有初始化.
	*/
	fmt.Println("结果 = ", DeferFunc1(1))
	fmt.Println("结果 = ", DeferFunc2(1))
	fmt.Println("结果 = ", DeferFunc3(1))
}

// threadSafeSet struct
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter_Wrong() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		// 加锁
		set.RLock()

		for elem := range set.s {
			fmt.Println("Iter_Wrong --- > elem = ", elem)
			ch <- elem
		}

		// 关闭channel 解锁.
		close(ch)
		set.RUnlock()
	}()

	return ch
}

// 现在的方法 是正确的.
func (set *threadSafeSet) Iter_Right() <-chan interface{} {
	ch := make(chan interface{}, len(set.s))

	set.RLock()
	defer set.RUnlock()
	for index, elem := range set.s {
		fmt.Println("Iter_Right --- > index= , elem = ", index, elem)
		ch <- elem
	}

	close(ch)

	return ch
}

// 考点2 defer的使用.
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	// fmt.Println("DeferFunc1 () --- t = ", t)
	return t
}

func DeferFunc2(i int) int {
	t := i

	defer func() {
		t += 3
		fmt.Println("DeferFunc2 () --defer-- t = ", t)
	}()

	func() {
		t += 100
	}()

	fmt.Println("DeferFunc2 () --- t = ", t)
	return t
}

// 最后一个比较特殊, 因为结尾是 return 2, 其实是在执行defer的上一步, return 变量t已经是 被赋值2了. 所以defer 函数内部执行的是 2+=1  t最后的结果是3
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
		// fmt.Println("DeferFunc3 ()  defer--- t = ", t)
	}()

	// fmt.Println("DeferFunc3 () --- t = ", t)
	return 2
}
