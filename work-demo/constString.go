package main

import "fmt"

const Separator = ":"

const Operation = "rotate"

func main() {
	str := fmt.Sprintf("%s%s", Separator, Operation)
	fmt.Println(str)
}
