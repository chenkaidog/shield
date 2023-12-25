package sensitive

import (
	"reflect"
)

type SensitiveMarshal struct {
	sensitiveSet map[string]bool
}

func NewSensitiveMarshal(words ...string) *SensitiveMarshal {
	set := make(map[string]bool)
	for _, word := range words {
		set[word] = true
	}

	return &SensitiveMarshal{
		sensitiveSet: set,
	}
}

func (sm *SensitiveMarshal) AddSensitiveWord(words ...string) {
	for _, word := range words {
		sm.sensitiveSet[word] = true
	}
}

func (sm *SensitiveMarshal) SafeMarshal(obj interface{}) string {
	return SafeJson(sm.sensitiveMarshal(obj))
}

func (sm *SensitiveMarshal) sensitiveMarshal(obj interface{}) interface{} {
	if obj == nil {
		return obj
	}

	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		if objValue.IsNil() {
			return obj
		}
		// 如果是指针，需要获取指针指向的结构体值
		objValue = objValue.Elem()
	}
	objType := objValue.Type()

	// 如果是结构体，遍历每个字段
	if objValue.Kind() == reflect.Struct {
		mapper := make(map[string]interface{})

		for i := 0; i < objValue.NumField(); i++ {
			field := objValue.Field(i)                  // 获取字段
			fieldName := objType.Field(i).Name          // 获取字段名称
			jsonTag := objType.Field(i).Tag.Get("json") // 获取字段tag中的json属性
			if jsonTag != "" {
				if jsonTag == "-" {
					continue
				}
				fieldName = jsonTag // 用json名称覆盖字段名称
			}

			if sm.sensitiveSet[fieldName] {
				mapper[fieldName] = "******"
			} else {
				if field.Kind() == reflect.Struct ||
					field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct {
					// 如果字段类型是结构体或者结构体指针，递归序列化
					mapper[fieldName] = sm.sensitiveMarshal(field.Interface())
				} else {
					// 其他类型直接添加
					mapper[fieldName] = field.Interface()
				}
			}
		}

		return mapper
	}

	// 否则直接返回
	return objValue.Interface()
}
