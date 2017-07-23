package main

import "fmt"

func main() {

	//Versin: A
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int)
		items[i][1] = 2
		items[i][2] = 3
	}
	fmt.Println("Version A :", items)

	//Version: B
	items1 := make([]map[int]int, 5)
	for _, item := range items {
		item = make(map[int]int)
		item[1] = 2
	}
	fmt.Println("Version B(bad style):", items1)
}

// 这里需要特别注意： 在 for 中定义的变量 它是地址唯一的
// 所以 Version B 都是操作的同一块地址
