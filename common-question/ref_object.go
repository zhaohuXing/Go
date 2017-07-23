package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func parseStudent1() map[string]student { // why not : map[string]student
	m := make(map[string]student)
	stus := []student{
		{Name: "yu", Age: 18},
		{Name: "ming", Age: 17},
		{Name: "zhu", Age: 16},
	}

	for _, stu := range stus {
		m[stu.Name] = stu
	}

	return m
}

func parseStudent2() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "xing", Age: 18},
		{Name: "zhao", Age: 18},
		{Name: "hu", Age: 18},
	}
	for i, _ := range stus {
		stu := stus[i]
		m[stu.Name] = &stu
	}

	return m
}

func main() {
	stus := parseStudent1()
	for k, v := range stus {
		fmt.Printf("key = %s, value = %v\n", k, v)
	}

	stus1 := parseStudent2()
	for k, v := range stus1 {
		fmt.Printf("key = %s, value = %v\n", k, *v)
	}
}

// 知识点：
//  0. 在 for 中定义的 val 在遍历过程中指针都是唯一的
//	1. map 的读取 : 见./map/
//  2. 那些可以用 range 遍历 : slice, map
//  3. 为什么要用 pointer ? : 做修改时用的，如果结构体较大时，用 pointer 会节约内存的

// map[keyType][valueType]
// 1. keyType 可以用 == or != 来比较的类型。常用: string, int, float,... 不符合标准的: slice, map,结构体..
// 2. valueType 任意类型, 对，什么都可以
