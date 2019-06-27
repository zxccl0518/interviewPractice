package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func pase_student() map[string]*student {
	m := make(map[string]*student)
	// 錯誤寫法
	stus := []student{
		student{Name: "zhou", Age: 24},
		student{Name: "li", Age: 23},
		student{Name: "wang", Age: 22},
	}

	// 解决方案1: 就是定义切片的时候 就是 *student 而不是 student
	// stus := []*student{
	// 	&student{Name: "zhou", Age: 24},
	// 	&student{Name: "li", Age: 23},
	// 	&student{Name: "wang", Age: 22},
	// }
	for _, stu := range stus {
		// 解决方案2: 给map赋值的时候,不能取地址.
		// var stuTemp = stu
		// m[stu.Name] = &stuTemp

		m[stu.Name] = &stu
	}
	return m
}
func main() {
	students := pase_student()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}
