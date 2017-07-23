package main

import "fmt"

func main() {
	weekMap := make(map[string]string)

	weekMap["Monday"] = "星期一"
	weekMap["Tuesday"] = "星期二"
	weekMap["Wednesday"] = "星期三"
	weekMap["Thursday"] = "星期四"
	weekMap["Friday"] = "星期五"
	weekMap["Saturday"] = "星期六"
	weekMap["Sunday"] = "星期七"

	for k, v := range weekMap {
		fmt.Println("key = " + k + ", value = " + v)
	}
	isExits(weekMap, "Tuesday")
	isExits(weekMap, "Holiday")

}

func isExits(m map[string]string, key string) {

	if value, ok := m[key]; ok {
		fmt.Println(key, " is exits, value is :", value)
	} else {
		fmt.Println(key, "is not exits")
	}

}
