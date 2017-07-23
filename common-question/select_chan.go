package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)

	intChan <- 1
	stringChan <- "hello ofo"

	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		fmt.Println(value)
		panic(value)
	}
}

// 单个 chan 如果没有缓冲， 将阻塞。但结合 select 可以在多个 chan 间等待执行.
// 但是它有三个原则:
// 1. select 中只要有一个 case 能 return, 就可以执行
// 2. 如果没有 case 能 return，则执行 default
// 3. 当如果同一时间有多个 case 均能 return 则伪随机方式抽取任意一个执行
