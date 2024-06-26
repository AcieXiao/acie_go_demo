package main

import "fmt"

// func main() {
// 	var a int
// 	fmt.Println(&a)
// 	var p *int
// 	p = &a
// 	*p = 20
// 	fmt.Println(a)
// }

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
