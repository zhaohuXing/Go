package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowB()
}

// 简单的组合， 这正是 Go 的亮点
// 如果 structA 存在 structB, 那么 structA "继承" structB 的方法。
// 这种方式叫做 组合， "继承" 更容易理解些(针对有 Java 经验的同学)
