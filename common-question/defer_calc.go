package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

// 多个 defer 执行的顺序是 LIFO
// 但是不代表： defer 后面的 代码不执行。方便理解，描述下 执行 顺序吧

// 上述代码执行顺序
// 执行到 defer calc("1", a, calc("10", a, b))
// 首先执行 calc("10", a, b),( 注意这里 int 值传递), 然后 执行 calc("1", a, 3), 将结果放入等待队列中
// 执行到 defer calc("2", a, calc("20", a, b))，然后执行 calc("2", a, 2), 将结果放入等待队列中
// 然后 倒叙执行
