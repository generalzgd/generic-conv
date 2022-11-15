/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://code.dobest.com/research-go
 * @software: GoLand
 * @file: any.go
 * @time: 2022/7/29 16:45
 * @project: generic-conv
 */

package gconv

import (
	"fmt"
	"reflect"
	"strconv"
)

// Any2Int
//  @Description: 任意类型转换成int
//  @param v
//  @return int
func Any2Int(v any) int {
	return int(Any2Int64(v))
}

// Any2Int32
//  @Description: 任意类型转换成int32
//  @param v
//  @return int32
func Any2Int32(v any) int32 {
	return int32(Any2Int64(v))
}

// Any2Uint64
//  @Description: 任意类型转换成unit64
//  @param v
//  @return uint64
func Any2Uint64(v any) uint64 {
	return uint64(Any2Int64(v))
}

// Any2Uint32
//  @Description: 任意类型转换成unit32
//  @param v
//  @return uint32
func Any2Uint32(v any) uint32 {
	return uint32(Any2Int64(v))
}

// Any2Int64
//  @Description: 任意类型转换成int64, 数值转换，字符串解析转换，其他返回0
//  @param v
//  @return int64
func Any2Int64(v any) int64 {
	if v == nil {
		return 0
	}
	switch d := v.(type) {
	case string:
		v, _ := strconv.ParseInt(d, 10, 64)
		return int64(v)
	case float32, float64:
		return int64(reflect.ValueOf(d).Float())
	case int, int8, int16, int32, int64:
		return int64(reflect.ValueOf(d).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(d).Uint())
	}
	return 0
}

// Any2Float64
//  @Description: 任意类型转换成float
//  @param v
//  @return float64
func Any2Float64(v any) float64 {
	if v == nil {
		return 0
	}
	switch d := v.(type) {
	case string:
		t, _ := strconv.ParseFloat(d, 64)
		return t
	case float32, float64:
		return reflect.ValueOf(d).Float()
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(d).Int())
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(d).Uint())
	}
	return 0
}

// Any2Float32
//  @Description: 任意类型转换成float
//  @param v
//  @return float32
func Any2Float32(v any) float32 {
	return float32(Any2Float64(v))
}

// Any2Bool
//  @Description: 任意类型转换成bool
//  @param v
//  @return bool
func Any2Bool(v any) bool {
	if v == nil {
		return false
	}
	switch d := v.(type) {
	case bool:
		return d
	case string:
		t, _ := strconv.ParseBool(d)
		return t
	case float32, float64:
		return reflect.ValueOf(d).Float() > 0.0
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(d).Int() > 0
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(d).Uint() > 0
	}
	return false
}

// Any2Str
//  @Description: 任意类型转换成字符串
//  @param v
//  @return string
func Any2Str(v any) string {
	if v == nil {
		return ""
	}
	switch d := v.(type) {
	case string:
		return d
	default:
		return fmt.Sprintf("%v", d)
	}
}
