/**
 * @Author HatsuneMona
 * @Date  2021-01-30 09:37
 **/
package main

import (
	. "20.leecode/leetcode"
	"20.leecode/tools"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "第1285个测试用例",
			args: args{
				l1: tools.CreateList([]int{0, 8, 6, 5, 6, 8, 3, 5, 7}),
				l2: tools.CreateList([]int{6, 7, 8, 0, 8, 5, 8, 9, 7}),
			},
			want: []int{6, 5, 5, 6, 4, 4, 2, 5, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tools.ListToArr(addTwoNumbers(tt.args.l1, tt.args.l2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
