package main

import "fmt"

var url = "https://studygolang.com/articles/11073?fr=sidebar"

func main() {
	// test1()

	test2()
}

// func test1(){
// 	jsonStr := []byte('{"age":1}')
// 	var value map[string] interface{}
// 	json.Unmarshal(jsonStr,&value)
// 	age:= value["age"]
// 	fmt.Println(reflect.TypeOf(age))
// }

func test2() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]

	for i := range s2 {
		s2[i] += 10
	}
	fmt.Println("s2 = ", s2)

	s2 = append(s2, 4)
	for i := range s2 {
		s2[i] += 10
	}
	fmt.Println("s2 = ", s2)
}

// s2 =  [12 13]
// s2 =  [22 23 14]
