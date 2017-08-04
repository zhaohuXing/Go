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

type MyInt int
type WaterMarkImageParam struct {
	Left    int     `json:"l"`
	Top     int     `json:"t"`
	Opacity float64 `json:"p,omitempty"`
	Url     string  `json:"u"`
	Fairel  MyInt   `json:"f,omitempty"`
	Flag    bool    `json:"fg,omitempty"`
}

func main() {
	param := &WaterMarkImageParam{
		Url:     "xxx",
		Opacity: 0.25,
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
		//if isEmptyValue(vf) &&

		if VStr == "0" &&
			len(tagValues) == 2 &&
			tagValues[1] == "omitempty" {
			continue
		}
		key = tagValues[0]
		kvPairs = append(kvPairs, fmt.Sprintf("%v%s%v", key, KVSep, value))
	}
	return strings.Join(kvPairs, KVPairSep)
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array,
		reflect.Map,
		reflect.Slice,
		reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return v.Int() == 0
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32,
		reflect.Float64:
		return v.Float() == 0
	case reflect.Ptr:
		return v.IsNil()
	}
	return false
}
