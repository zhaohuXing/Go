package main

import "fmt"

func main() {
	fmt.Println(calc())
	fmt.Println(calc1())
	fmt.Println(calc2())
}

func calc() (a int) {
	a = 5
	defer func() {
		a = a * 2
	}()
	return a
}

func calc1() int {
	a := 5
	defer func() {
		a = a * 2
	}()
	return a
}

func calc2() int {
	a := 5
	defer func(i int) {
		i = i * 2
	}(a)
	return a
}

// 不理解
// TODO
/*
func main() {
	fmt.Println(doubleScore(0))    //0
	fmt.Println(doubleScore(20.0)) //40
	fmt.Println(doubleScore(50.0)) //50
}
func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			//将影响返回值
			score = source
		}
	}()
	score = source * 2
	return

	//或者
	//return source * 2
}*/

// 闭包
// 作用域
// defer 运行原理
