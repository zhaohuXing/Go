package main

import "fmt"

func main() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)

	map1["BeiJing"] = 20170715
	map1["ShangHai"] = 20180130
	value, isPresent = map1["BeiJing"]
	if isPresent {
		fmt.Println("我来到北京的日期:", value)
	}

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	delete(map1, "BeiJing")
	for k, v := range map1 {
		fmt.Println(k, v)
	}
}
