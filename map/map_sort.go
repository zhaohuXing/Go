package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "Jravo": 56,
		"charlie": 23, "delta": 45}
)

func main() {
	fmt.Println("Unsorted:")
	keys := make([]string, len(barVal))
	i := 0
	for k, v := range barVal {
		keys[i] = k
		i++
		fmt.Println(k, ":", v)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, ":", barVal[k])
	}

}
