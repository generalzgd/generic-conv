/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://code.dobest.com/research-go
 * @software: GoLand
 * @file: slice.go
 * @time: 2022/7/29 15:24
 * @project: generic-gconv
 */

package gconv

import "fmt"

// Integer
//  @Description: 整数定义
type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Number
// @Description: 数字定义 整数+浮点数
type Number interface {
	Integer | float32 | float64
}

// NumStr
// @Description: 数字和字符串
type NumStr interface {
	Number | string
}

// Slice2Any [T comparable]
//  @Description: 任意类型切片转成[]any
//  @param in []T
//  @return []any
func Slice2Any[T comparable](in []T) []any {
	out := make([]any, len(in))
	for i, it := range in {
		out[i] = it
	}
	return out
}

// SliceNum2Str [T Number]
//  @Description: 将数值切片转换成字符串切片
//  @param in []T
//  @return []string
func SliceNum2Str[T Number](in []T) []string {
	out := make([]string, len(in))
	for i, it := range in {
		out[i] = fmt.Sprintf("%v", it)
	}
	return out
}

// SliceNum2Int [T Number]
//  @Description: 将数值切片转换成int切片
//  @param in
//  @return []int
func SliceNum2Int[T Number](in []T) []int {
	out := make([]int, len(in))
	for i, it := range in {
		out[i] = int(it)
	}
	return out
}

// SliceNum2Uint [T Number]
//  @Description: 将数值切片转换成uint切片
//  @param in
//  @return []uint
func SliceNum2Uint[T Number](in []T) []uint {
	out := make([]uint, len(in))
	for i, it := range in {
		out[i] = uint(it)
	}
	return out
}

// SliceNum2Int32 [T Number]
//  @Description: 将数值切片转换成int32切片
//  @param in
//  @return []int32
func SliceNum2Int32[T Number](in []T) []int32 {
	out := make([]int32, len(in))
	for i, it := range in {
		out[i] = int32(it)
	}
	return out
}

// SliceNum2Uint32 [T Number]
//  @Description: 将数值切片转换成uint32切片
//  @param in
//  @return []uint32
func SliceNum2Uint32[T Number](in []T) []uint32 {
	out := make([]uint32, len(in))
	for i, it := range in {
		out[i] = uint32(it)
	}
	return out
}

// SliceNum2Int64 [T Number]
//  @Description: 将数值切片转换成int64切片
//  @param in
//  @return []int64
func SliceNum2Int64[T Number](in []T) []int64 {
	out := make([]int64, len(in))
	for i, it := range in {
		out[i] = int64(it)
	}
	return out
}

// SliceNum2Uint64 [T Number]
//  @Description: 将数值切片转换成uint64切片
//  @param in
//  @return []uint64
func SliceNum2Uint64[T Number](in []T) []uint64 {
	out := make([]uint64, len(in))
	for i, it := range in {
		out[i] = uint64(it)
	}
	return out
}

// SliceMinus [T comparable]
//  @Description: 切片减法
//  @param in1
//  @param in2
//  @return []T
func SliceMinus[T comparable](in1, in2 []T) []T {
	if len(in1) < 1 || len(in2) < 2 {
		return in1
	}
	tmp := make(map[T]struct{}, len(in1))
	for _, it := range in1 {
		tmp[it] = struct{}{}
	}
	for _, it := range in2 {
		delete(tmp, it)
	}
	out := make([]T, 0, len(tmp))
	for k := range tmp {
		out = append(out, k)
	}
	return out
}

// SliceMerge [T comparable]
//  @Description: 切片合并，且保证唯一
//  @param ins
//  @return []T
func SliceMerge[T comparable](ins ...[]T) []T {
	if len(ins) < 1 {
		return nil
	}
	tmp := map[T]struct{}{}
	for _, in := range ins {
		for _, it := range in {
			tmp[it] = struct{}{}
		}
	}
	//
	out := make([]T, 0, len(tmp))
	for k := range tmp {
		out = append(out, k)
	}
	return out
}

// SliceUnique [T comparable]
//  @Description: 惟一化处理
//  @param in
//  @return []T
func SliceUnique[T comparable](in []T) []T {
	keys := make(map[T]struct{}, len(in))
	for _, ent := range in {
		keys[ent] = struct{}{}
	}
	return MapKeys(keys)
}

// IsSubSlice [T comparable]
//  @Description: 切片子集校验
//  @param sub
//  @param tank
//  @return bool
func IsSubSlice[T comparable](sub, tank []T) bool {
	if sub == nil {
		return true
	}
	if tank == nil {
		return false
	}
	tmp := make(map[T]struct{}, len(tank))
	for _, v := range tank {
		tmp[v] = struct{}{}
	}
	for _, v := range sub {
		if _, ok := tmp[v]; !ok {
			return false
		}
	}
	return true
}
