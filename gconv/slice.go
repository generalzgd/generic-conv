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

// Slice2Map [T comparable]
//  @Description: 切片转map
//  @param in
//  @return map[T]struct{}
func Slice2Map[T comparable](in []T) map[T]struct{} {
	out := make(map[T]struct{}, len(in))
	for _, it := range in {
		out[it] = struct{}{}
	}
	return out
}

// SliceStr2Int [T string]
//  @Description: 字符串切片转成int切片
//  @param in
//  @return []int64
func SliceStr2Int[T string](in []T) []int {
	out := make([]int, len(in))
	for i, it := range in {
		out[i] = Any2Int(it)
	}
	return out
}

// SliceStr2Int64 [T string]
//  @Description: 字符串切片转成int64切片
//  @param in
//  @return []int64
func SliceStr2Int64[T string](in []T) []int64 {
	out := make([]int64, len(in))
	for i, it := range in {
		out[i] = Any2Int64(it)
	}
	return out
}

// SliceStr2Uint64 [T string]
//  @Description: 字符串切片转成uint64切片
//  @param in
//  @return []int64
func SliceStr2Uint64[T string](in []T) []uint64 {
	out := make([]uint64, len(in))
	for i, it := range in {
		out[i] = Any2Uint64(it)
	}
	return out
}

// SliceStr2Int32 [T string]
//  @Description: 字符串切片转成int32切片
//  @param in
//  @return []int64
func SliceStr2Int32[T string](in []T) []int32 {
	out := make([]int32, len(in))
	for i, it := range in {
		out[i] = Any2Int32(it)
	}
	return out
}

// SliceStr2Uint32 [T string]
//  @Description: 字符串切片转成uint32切片
//  @param in
//  @return []int64
func SliceStr2Uint32[T string](in []T) []uint32 {
	out := make([]uint32, len(in))
	for i, it := range in {
		out[i] = Any2Uint32(it)
	}
	return out
}

// SliceStr2Float64 [T string]
//  @Description: 字符串切片转成float64切片
//  @param in
//  @return []float64
func SliceStr2Float64[T string](in []T) []float64 {
	out := make([]float64, len(in))
	for i, it := range in {
		out[i] = Any2Float64(it)
	}
	return out
}

// SliceStr2Float32 [T string]
//  @Description: 字符串切片转成float32切片
//  @param in
//  @return []float64
func SliceStr2Float32[T string](in []T) []float32 {
	out := make([]float32, len(in))
	for i, it := range in {
		out[i] = Any2Float32(it)
	}
	return out
}

// SliceStr2Bool [T string]
//  @Description: 字符串切片转成bool切片
//  @param in
//  @return []bool
func SliceStr2Bool[T string](in []T) []bool {
	out := make([]bool, len(in))
	for i, it := range in {
		out[i] = Any2Bool(it)
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

// SliceUnion [T comparable]
//  @Description: 切片合并，且保证唯一
//  @param ins
//  @return []T
func SliceUnion[T comparable](ins ...[]T) []T {
	if len(ins) < 1 {
		return nil
	}
	count := 0
	for _, in := range ins {
		count += len(in)
	}
	tmp := make(map[T]struct{}, count)
	for _, in := range ins {
		for _, it := range in {
			tmp[it] = struct{}{}
		}
	}
	//
	return MapKeys(tmp)
}

// SliceIntersect [T comparable]
//  @Description:  切片交集。返回在两个切片中均存在的列表
//  @param first
//  @param second
//  @return []T
func SliceIntersect[T comparable](first, second []T) []T {
	f := Slice2Map(first)
	s := Slice2Map(second)
	//
	return MapIntersect(f, s)
}

// SliceExcept [T comparable]
//  @Description: 切片差集。返回在第一个切片中存在，第二个切片中不存在的列表
//  @param first
//  @param second
//  @return []T
func SliceExcept[T comparable](first, second []T) []T {
	f := Slice2Map(first)
	s := Slice2Map(second)
	//
	return MapExcept(f, s)
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
