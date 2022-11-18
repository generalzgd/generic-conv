/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://code.dobest.com/research-go
 * @software: GoLand
 * @file: map.go
 * @time: 2022/7/29 16:22
 * @project: generic-gconv
 */

package gconv

import (
	"fmt"
	"math"
	"reflect"
)

// MapKvPairs [T1, T2 comparable]
//  @Description: 转换map为键值对
//  @param in
//  @return []any
func MapKvPairs[T1, T2 comparable](in map[T1]T2) []any {
	out := make([]any, 0, len(in)*2)
	for k, v := range in {
		out = append(out, k, v)
	}
	return out
}

// MapKeys [T1, T2 comparable]
//  @Description: 提取map的键
//  @param in
//  @return []T1
func MapKeys[T1, T2 comparable](in map[T1]T2) []T1 {
	out := make([]T1, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

// MapValues [T1, T2 comparable]
//  @Description: 提取map的值
//  @param in
//  @return []T2
func MapValues[T1, T2 comparable](in map[T1]T2) []T2 {
	out := make([]T2, 0, len(in))
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

// MapUnion [T1, T2 comparable]
//  @Description: map合并操作，如果存在同key,不同value则会被后者覆盖
//  @param ins
//  @return map[T1]T2
func MapUnion[T1, T2 comparable](ins ...map[T1]T2) map[T1]T2 {
	if len(ins) < 1 {
		return map[T1]T2{}
	}
	count := 0
	for _, in := range ins {
		count += len(in)
	}

	out := make(map[T1]T2, count)
	for _, other := range ins {
		for k, v := range other {
			out[k] = v
		}
	}
	return out
}

// MapIntersect [T1, T2 comparable]
//  @Description: map交集操作，返回交集的key列表。在两个map中都存在的key
//  @param first
//  @param second
//  @return []T1
func MapIntersect[T1, T2 comparable](first, second map[T1]T2) []T1 {
	out := make([]T1, 0, int(math.Max(float64(len(first)), float64(len(second)))))
	for k := range first {
		if _, ok := second[k]; ok {
			out = append(out, k)
		}
	}
	return out
}

// MapExcept [T1, T2 comparable]
//  @Description: map差集操作，返回差集的key列表。key在第一个map中，且不在第二个map中。
//  @param first
//  @param second
//  @return []T1
func MapExcept[T1, T2 comparable](first, second map[T1]T2) []T1 {
	out := make([]T1, 0, int(math.Max(float64(len(first)), float64(len(second)))))
	for k := range first {
		if _, ok := second[k]; !ok {
			out = append(out, k)
		}
	}
	return out
}

// MapVal2Int [T1, T2 comparable]
//  @Description: 将map的value转换为int
//  @param in
//  @return map[T1]int
func MapVal2Int[T1, T2 comparable](in map[T1]T2) map[T1]int {
	out := make(map[T1]int, len(in))
	for k, v := range in {
		out[k] = Any2Int(v)
	}
	return out
}

// MapVal2Str [T1, T2 comparable]
//  @Description: 将map的value转换为string
//  @param in
//  @return map[T1]string
func MapVal2Str[T1, T2 comparable](in map[T1]T2) map[T1]string {
	out := make(map[T1]string, len(in))
	for k, v := range in {
		out[k] = Any2Str(v)
	}
	return out
}

// MapVal2Any [T1, T2 comparable]
//  @Description: 将map的value转换为any
//  @param in
//  @return map[T1]any
func MapVal2Any[T1, T2 comparable](in map[T1]T2) map[T1]any {
	out := make(map[T1]any, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

//
// MapFromStruct
//  @Description: 将一级结构体转换成map, 且只支持一级。如果有深层结构嵌套请使用deepcopy库
//  @param in 任意一级结构体对象
//  @return map[string]any 对应的map
//  @return error 异常错误
//
func MapFromStruct(in interface{}) (map[string]any, error) {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("StructToMap only accepts structs; got %T", v)
	}
	out := make(map[string]any, v.NumField())
	//typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := v.Field(i)
		if fi.Kind() == reflect.Ptr {
			fi = fi.Elem()
		}
		ft := v.Type().Field(i)
		if name := ft.Name; name != "" {
			out[name] = fi.Interface()
		}
	}
	return out, nil
}
