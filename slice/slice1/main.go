package main

import (
	"fmt"
	"sync"
)

func sliceTest1() {
	s := []int{1, 2, 3}
	// s := make([]int, 10)
	ss := s[1:]

	// append 一旦对slice 进行扩容了, 那么返回的的slice 与原来的slice 地址引用就不同了.
	// 也就是说是新的 slice ,彼此改动不会影响到彼此.
	ss = append(ss, 4)

	// 说明通过这个range 的方式改变slice 是错的,
	// v 就是一个临时变量, 这个变量里面 不停的复制 slice 的值.
	for _, v := range ss {
		v += 10
	}

	fmt.Println(" s = ", s)
	fmt.Println(" ss = ", ss)

	// 要想改变这个slice的值, 要通过下表索引的方式 去改变值.
	for i := range ss {
		ss[i] += 10
	}

	fmt.Println("--------------------------")
	fmt.Println(" s = ", s)
	fmt.Println(" ss = ", ss)
}

// ---------------------------------------
type S struct {
}
type IF interface {
	F()
}

func (s S) F() {

}

func InitType() S {
	var s S
	return s
}
func InitPointer() *S {
	var s *S
	return s
}

func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func InitIfaceType() IF {
	var s S
	return s
}

func InitIfacePointer() IF {
	var s *S
	return s
}

func Interface_structTest1() {
	// 无法通过编译, 因为无法将 nil 转换为 S
	// cannot convert nil to type Sgo
	// fmt.Println(InitType() == nil)

	// 返回的是结构体的指针,因为结构体没有实例化,只有声明,所以指针 是一个nil的空值.所以返回true
	fmt.Println(InitPointer() == nil)
	// 因为这个函数返回值是 interface{}, 所以return之前 会将变量变成一个空的interface{} 但是空interface 是数据为空,类型不为空.
	// 所以 != nil  返回false
	fmt.Println(InitEfaceType() == nil)
	// 同理 返回的一个空接口 != nil  结果为false
	fmt.Println(InitEfacePointer() == nil)

	// 返回值  是一个已经被实现了的接口, 不为空.
	fmt.Println(InitIfaceType() == nil)
	fmt.Println(InitIfacePointer() == nil)
}

//--------------------

type V struct {
	m string
}

func f() *V {
	return &V{"foo"} //A
}

func returnTest1() {
	p := *f()  //B
	print(p.m) //print "foo"
}

//-----------------------------------------------------------
type S1 struct {
}

func z(x interface{}) {
}

func g(x *interface{}) {
}

func InterfaceTest2() {
	s := S1{}
	p := &s
	var w interface{}
	z(s) //A
	// g(s) //B  cannot use s (type S1) as type *interface {} in argument to g:*interface {} is pointer to interface, not interface
	z(p) //C
	// g(p) //D  cannot use s (type *S1) as type *interface {} in argument to g:*interface {} is pointer to interface, not interface

	// 下面2种情况 是对的.
	w = s
	g(&w) // E
	w = p
	g(&w) // F
}

//-----------------------------------------------------------

func main() {
	// sliceTest1()

	// Interface_structTest1()

	// returnTest1()

	// InterfaceTest2()

	N := 10
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	for i, v := range m {
		fmt.Printf("i = %v, v = %v\n", i, v)
	}
	println(len(m))
}
