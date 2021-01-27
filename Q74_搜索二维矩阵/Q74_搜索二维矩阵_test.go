/**
 * @Author HatsuneMona
 * @Date  2021-01-27 10:10
 * @Description //TODO
 **/
package main

import "testing"

func TestBinarySearch(t *testing.T) {
	mat := []int{1, 3, 5, 7, 9, 10, 12, 14, 16, 18, 20}
	type args struct {
		matrix  []int
		target  int
		a       int
		b       int
		precise bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "精确查找，存在",
			args: args{
				matrix:  mat,
				target:  7,
				a:       0,
				b:       len(mat),
				precise: true,
			},
			want: 3,
		}, {
			name: "精确查找，不存在",
			args: args{
				matrix:  mat,
				target:  15,
				a:       0,
				b:       len(mat),
				precise: true,
			},
			want: -1,
		},
		{
			name: "非精确查找，存在",
			args: args{
				matrix:  mat,
				target:  8,
				a:       0,
				b:       len(mat),
				precise: false,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.matrix, tt.args.target, tt.args.a, tt.args.b, tt.args.precise); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例一",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
				target: 3,
			},
			want: true,
		},
		{
			name: "示例二",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
				target: 13,
			},
			want: false,
		},
		{
			name: "出错",
			args: args{
				matrix: [][]int{{1, 3}},
				target: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
