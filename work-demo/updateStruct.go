package main

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	KVSep     = "_"
	KVPairSep = ","
)

type WaterMarkImageParam struct {
	Left    int     `json:"l"`
	Top     int     `json:"t"`
	Opacity float64 `json:"p,omitempty"`
	Url     string  `json:"u"`
}

func main() {
	param := &WaterMarkImageParam{
		Url: "xxx",
	}
	fmt.Println(buildOptParamStr(param))
}

func buildOptParamStr(param interface{}) string {
	v := reflect.ValueOf(param).Elem()
	var kvPairs []string

	for i := 0; i < v.NumField(); i++ {
		vf := v.Field(i)
		tf := v.Type().Field(i)
		key := tf.Tag.Get("json")
		value := vf.Interface()

		VStr := fmt.Sprintf("%v", value)
		tagValues := strings.Split(key, ",")
		if VStr == "0" && len(tagValues) == 2 {
			continue
		}
		key = tagValues[0]
		kvPairs = append(kvPairs, fmt.Sprintf("%v%s%v", key, KVSep, value))
	}
	return strings.Join(kvPairs, KVPairSep)
}
