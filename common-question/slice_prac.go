package main

import "fmt"

func main() {
	nums := make([]int, 5)
	nums = append(nums, 1, 3, 4)
	fmt.Println(nums)
}

// slice 默认初设容量，默认值
// 自动扩容
