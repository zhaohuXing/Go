package main

import "fmt"
import "reflect"

type MyInt int

func main() {
	var a MyInt = 2
	fmt.Println(reflect.TypeOf(a).Kind())
}
