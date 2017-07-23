package main

import "fmt"

func main() {
	fmt.Println(testDefer(2))
}

func testDefer(a int) int {
	defer func() {
		a = a * 2
	}()
	return a
}

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
