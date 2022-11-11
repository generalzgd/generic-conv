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
	"reflect"
)

// Map2KvPairs [T1, T2 comparable]
//  @Description: 转换map为键值对
//  @param in
//  @return []any
func Map2KvPairs[T1, T2 comparable](in map[T1]T2) []any {
	out := make([]any, 0, len(in)*2)
	for k, v := range in {
		out = append(out, k, v)
	}
	return out
}

// GetMapKeys [T1, T2 comparable]
//  @Description: 提取map的键
//  @param in
//  @return []T1
func GetMapKeys[T1, T2 comparable](in map[T1]T2) []T1 {
	out := make([]T1, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

// MapMerge [T1, T2 comparable]
//  @Description: map合并操作，如果存在同key,不同value则会被后者覆盖
//  @param in
//  @param others
//  @return map[T1]T2
func MapMerge[T1, T2 comparable](in map[T1]T2, others ...map[T1]T2) map[T1]T2 {
	if len(in) < 1 {
		return map[T1]T2{}
	}

	for _, other := range others {
		for k, v := range other {
			in[k] = v
		}
	}

	return in
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
// StructToMap
//  @Description: 将一级结构体转换成map, 且只支持一级。如果有深层结构嵌套请使用deepcopy库
//  @param in 任意一级结构体对象
//  @return map[string]any 对应的map
//  @return error 异常错误
//
func StructToMap(in interface{}) (map[string]any, error) {
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
