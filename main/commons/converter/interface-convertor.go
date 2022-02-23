package converter

import (
	"fmt"
	"reflect"
	"time"
)

func InterfaceToString(data interface{}) (string, error) {
	return fmt.Sprintf("%v", data), nil
}

func InterfaceToUint(data interface{}) (uint, error) {
	res, ok := data.(uint)
	if !ok {
		return 0, fmt.Errorf("convert error: '%v' can't convert to uint", data)
	}
	return res, nil
}

func InterfaceToFloat32(data interface{}) (float32, error) {
	res, ok := data.(float32)
	if !ok {
		return 0, fmt.Errorf("convert error: '%v' can't convert to float32", data)
	}
	return res, nil
}

func InterfaceToDate(data interface{}, format string) (time.Time, error) {
	if str, ok := data.(string); ok {
		return StrToTime(str, format)
	} else {
		return time.Now(), fmt.Errorf("convert error: '%v' can't convert to date with format %s", data, format)
	}
}

func InterfaceToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		fmt.Println(tag, field)
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = InterfaceToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}
