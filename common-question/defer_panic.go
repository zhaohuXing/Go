package main

import "fmt"

func main() {
	defer_call()
}

func defer_call() {
	defer fmt.Println("最后打印")
	defer fmt.Println("中间打印")
	defer fmt.Println("最先打印")

	panic("抛出个异常")
}

//知识点：
//	1. 多个 defer 按照 先进后出 的执行顺序执行
//  2. return value 并不是原子操作，先是赋值，后是return, 然而 defer 执行发生在 二者中间
