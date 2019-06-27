package main

import (
	"errors"
	"fmt"
)

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
	q
)

// 考点:变量简短模式
// 变量简短模式限制：

// 定义变量同时显式初始化
// 不能提供数据类型
// 只能在函数内部使用
var (
	// size := 1024	// 错误写法 这里不是函数内部, 不能使用简短模式 进行变量初始化.
	size     = 1024
	max_size = size * 2
)

func main() {
	// 所以这个编译 均不通过.因为这个new 返回值是一个 指针类型, 而append需要的参数;类型是[]int的切片. 而非*[]int
	// list := new([]int)
	// list = append(list, 1)
	// list = append(*list, 1)
	// fmt.Println(list)

	// 编译不通过. 因为append 操作切片的时候, 如果2个都是切片一定要在后面的那个切片 加上 ...
	// s1 := []int{1, 2, 3}
	// s2 := []int{4, 5}
	// // s1 = append(s1, s2) //.\main.go:15:13: cannot use s2 (type []int) as type int in append
	// s1 = append(s1, s2...) // 正确写法.
	// fmt.Println("s1 = ", s1)

	// 编译不通过, 结果体的比较,首先要结构体的类型相同, 其次比较每个字段 是否相同. 这里的第二个字段map[string]string 没办法进行比较,所以编译报错. 类似不能比较的 map slice interface channel
	// 结构体是否相同不但与属性类型个数有关,还与属性顺序相关。
	// sn1 := struct {
	// 	age  int
	// 	name string
	// }{age: 11, name: "qq"}
	// sn2 := struct {
	// 	age  int
	// 	name string
	// }{age: 11, name: "qq"}

	// sn3 := struct {
	// 	name string
	// 	age  int
	// }{name: "zxc", age: 10}

	// 这里的sn3 与 sn1 就已经不是不同的类型的结果体了.
	// if sn1 == sn2 {
	// 	fmt.Println("sn1 == sn2")
	// }

	// sm1 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}
	// sm2 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}

	// if sm1 == sm2 { //.\main.go:41:9: invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)
	// 	fmt.Println("sm1 == sm2")
	// }

	// var x *int = nil
	// Foo(x)
	// non-empty interface

	// 考点是 iota 的理解.
	// fmt.Println(x, y, z, k, p, q)

	// 考点是 简短模式 只能在函数内部使用.
	// fmt.Println(size, max_size)

	// 考点是 函数内部的变量的作用域.
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing() // 注意这里err 是重新生成的一个变量, 这里的err不同于外面的err
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}
