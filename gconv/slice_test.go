/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://code.dobest.com/research-go
 * @software: GoLand
 * @file: slice_test.go
 * @time: 2022/7/29 15:32
 * @project: generic-gconv
 */

package gconv

import (
	"testing"
)

func TestSlice2Interface(t *testing.T) {

	got := Slice2Any([]string{"1", "b", "c"})
	t.Logf("got string slice result: %v", got)

	got = Slice2Any([]int{1, 2, 3})
	t.Logf("got int slice result: %v", got)

	got = Slice2Any([]uint{1, 2, 3})
	t.Logf("got uint slice result: %v", got)

	got = Slice2Any([]float32{0.1, 0.2, 0.3})
	t.Logf("got float slice result: %v", got)

}

func TestNumSlice2StrSlice(t *testing.T) {

	got := SliceNum2Str([]uint{1, 2, 3})
	t.Logf("got string slice result: %v", got)
}

func TestSliceMinus(t *testing.T) {

	got1 := SliceMinus([]int{1, 2, 3}, []int{1, 6})
	t.Logf("got int slice minus result: %v", got1)

	got2 := SliceMinus([]string{"1", "2", "3"}, []string{"1", "6"})
	t.Logf("got str slice minus result: %v", got2)
}

func TestSliceMerge(t *testing.T) {

	got1 := SliceMerge([]int{1, 2, 3}, []int{1, 6})
	t.Logf("got int slice minus result: %v", got1)

	got2 := SliceMerge([]string{"1", "2", "3"}, []string{"1", "6"})
	t.Logf("got str slice minus result: %v", got2)
}
