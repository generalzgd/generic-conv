/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://code.dobest.com/research-go
 * @software: GoLand
 * @file: map_test.go
 * @time: 2022/7/29 16:27
 * @project: generic-gconv
 */

package gconv

import (
	"testing"
)

func TestMap2KvPairs(t *testing.T) {

	got := MapKvPairs(map[string]int{"a": 1})
	t.Logf("got map kv pairs: %v", got)

	got = MapKvPairs(map[int]float32{100: 1.0})
	t.Logf("got map kv pairs: %v", got)
}

func TestGetMapKeys(t *testing.T) {
	got := MapKeys(map[int]float32{100: 1.0})
	t.Logf("got map keys: %v", got)
}

func TestMapMerge(t *testing.T) {

	got := MapMerge(map[string]int{"a": 1}, map[string]int{"b": 2})
	t.Logf("MapMerge() = %v", got)

}

func TestMapVal2Int(t *testing.T) {
	got := MapVal2Int(map[string]string{"a": "120"})
	t.Logf("MapVal2Int() = %v", got)

	got2 := MapVal2Int(map[int]string{100: "120"})
	t.Logf("MapVal2Int() = %v", got2)
}

func TestMapVal2Str(t *testing.T) {
	got := MapVal2Str(map[int]int{100: 120})
	t.Logf("MapVal2Str() = %v", got)
}

func TestMapVal2Any(t *testing.T) {
	var got map[int]interface{}
	got = MapVal2Any(map[int]int{100: 120})
	t.Logf("MapVal2Any() = %v", got)
}

func TestStructToMap(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestStructToMap",
			args: args{
				in: struct {
					A int
					B *int
					C string
				}{
					A: 1,
					B: func() *int {
						v := 2
						return &v
					}(),
					C: "asdf",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapFromStruct(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("StructToMap() got = %v", got)
		})
	}
}
