package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type MyInt int64

func main() {

	var a uint8 = 2
	var b int = 37
	var c string = "3.2"
	res := sum(a, b, c)
	fmt.Printf("result: %f\n", res)
	var d MyInt = 1
	res = sum(a, b, c, d)
	fmt.Printf("result: %f\n", res)
}

func sum(v ...interface{}) float64 {
	var res float64 = 0
	for _, val := range v {
		ref := reflect.ValueOf(val)
		switch ref.Kind() {
		case reflect.Int, reflect.Int64:
			res += float64(ref.Int())
		case reflect.Uint8:
			res += float64(ref.Uint())
		case reflect.String:
			a, err := strconv.ParseFloat(ref.String(), 64)
			if err != nil {
				panic(err)
			}
			res += a
		default:
			fmt.Println("Unsupported type %T. Ignoring.\n", val)
		}
	}
	return res
}
