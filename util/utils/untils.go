package utils

import (
	"fmt"
	"reflect"
)

// GetColumn  泛型方法，直接返回指定字段的具体类型切片
func Pluck[T any, R any](slice []T, fieldName string) ([]R, error) {
	v := reflect.ValueOf(slice)

	// 检查输入是否是一个切片
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected a slice, got %s", v.Kind())
	}

	var result []R
	for i := 0; i < v.Len(); i++ {
		// 获取切片中每个元素的值
		elem := reflect.ValueOf(slice[i])

		// 检查元素是否是结构体
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		if elem.Kind() != reflect.Struct {
			return nil, fmt.Errorf("expected a struct, got %s", elem.Kind())
		}

		// 获取指定字段的值
		fieldValue := elem.FieldByName(fieldName)
		if !fieldValue.IsValid() {
			return nil, fmt.Errorf("no such field: %s in struct", fieldName)
		}

		result = append(result, fieldValue.Interface().(R))
	}

	return result, nil
}

func Difference[T comparable](a, b []T) []T {
	var diff []T
	setB := make(map[any]bool)

	// 将 b 数组的元素存入 map 中，方便查找
	for _, val := range b {
		setB[val] = true
	}

	// 遍历 a 数组，如果元素不在 b 数组中，就将其添加到差集中
	for _, val := range a {
		if !setB[val] {
			diff = append(diff, val)
		}
	}

	return diff
}
