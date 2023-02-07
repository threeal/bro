package utils

import (
	"reflect"
	"strings"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetJSONFields(j interface{}) []string {
	val := reflect.ValueOf(j)
	var result []string
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		fieldName := t.Name
		switch jsonTag := t.Tag.Get("json"); jsonTag {
		case "-":
		case "":
			result = append(result, fieldName)
		default:
			parts := strings.Split(jsonTag, ",")
			name := parts[0]
			if name == "" {
				name = fieldName
			}
			result = append(result, name)
		}
	}
	return result
}

func GetStructValueByJSON(j interface{}, key string) string {
	val := reflect.ValueOf(j)
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		switch jsonTag := t.Tag.Get("json"); jsonTag {
		case "-":
		case "":
		default:
			parts := strings.Split(jsonTag, ",")
			if parts[0] == key {
				return val.Field(i).String()
			}
		}
	}
	return ""
}

func SetStructValueByJSON(j interface{}, key string, value string) {
	valPtr := reflect.ValueOf(j)
	val := reflect.Indirect(reflect.ValueOf(j))
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		fieldName := t.Name
		switch jsonTag := t.Tag.Get("json"); jsonTag {
		case "-":
		case "":
		default:
			parts := strings.Split(jsonTag, ",")
			if parts[0] == key {
				valPtr.Elem().FieldByName(fieldName).SetString(value)
				return
			}
		}
	}
}
