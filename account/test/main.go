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
	objValue := reflect.ValueOf(obj).Elem()
    objType := objValue.Type()

    mapper := make(map[string]interface{})

    for i := 0; i < objValue.NumField(); i++ {
        field := objValue.Field(i)
        fieldName := objType.Field(i).Name

		if sm.sensitiveSet[fieldName] {
			mapper[fieldName] = "******"
		} else {
			if field.Kind() == reflect.Struct || 
			field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct{
				mapper[fieldName] = sm.Marshal(field.Interface())
			}
			mapper[fieldName] = field.Interface()
		}
	}

	return mapper
}

type User struct {
	Name string 
	Info1 Info1
	Info2 *Info2
}

type Info1 struct {
	Password string 
}

type Info2 struct {
	List []string 
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
}
