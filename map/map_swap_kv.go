package main

import "fmt"

var (
	KValue = map[string]int{
		"lu":   2014,
		"dong": 9,
	}
)

func main() {
	VKey := make(map[int]string, len(KValue))
	for k, v := range KValue {
		VKey[v] = k
	}

	for k, v := range VKey {
		fmt.Println(k, v)
	}
}
