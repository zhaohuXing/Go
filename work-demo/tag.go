package main

import (
	"fmt"
	"reflect"
)

type CorpGravity int
type CropParam struct {
	Width   int         `schema:"w"`
	Height  int         `schema:"h"`
	Gravity CorpGravity `schema:"g" omitempty:"true"`
}

func main() {
	crop := &CropParam{Width: 100, Height: 0, Gravity: 0}
	v := reflect.ValueOf(crop).Elem()

	for i := 0; i < v.NumField(); i++ {
		vf := v.Field(i)
		tf := v.Type().Field(i)
		value := vf.Interface()
		key := tf.Tag.Get("omitempty")
		null := fmt.Sprintf("%v", value)

		if null == "" || null == "0" {
			isOmitEmpty := tf.Tag.Get("omitempty")
			if isOmitEmpty == "" {
				continue
			}

		}
		fmt.Println(key)
		fmt.Println(value)
	}
}
