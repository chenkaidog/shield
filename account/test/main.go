package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type SensitiveMarshal struct {
	sensitiveSet map[string]bool
}

func (sm *SensitiveMarshal) Marshal(obj interface{}) interface{} {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	objType := objValue.Type()

	mapper := make(map[string]interface{})

	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		fieldName := objType.Field(i).Name
		jsonTag := objType.Field(i).Tag.Get("json")
		if jsonTag != "" {
			if jsonTag == "-" {
				continue
			}
			fieldName = jsonTag
		}

		if sm.sensitiveSet[fieldName] {
			mapper[fieldName] = "******"
		} else {
			if field.Kind() == reflect.Struct ||
				field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct {
				mapper[fieldName] = sm.Marshal(field.Interface())
			} else {
				mapper[fieldName] = field.Interface()
			}
		}
	}

	return mapper
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"-"`
	Info1    Info1
	Info2    *Info2
}

type Info1 struct {
	Password string
}

type Info2 struct {
	List     []string
	Password string
}

func main() {
	user := &User{
		Name: "Jack",
		Info1: Info1{
			Password: "123456",
		},
		Info2: &Info2{
			List: []string{"elem1", "elem2"},
		},
	}

	sm := &SensitiveMarshal{
		sensitiveSet: map[string]bool{"Password": true},
	}

	fmt.Println(sm.Marshal(user))

	data, _ := json.Marshal(sm.Marshal(user))
	fmt.Println(string(data))
}
