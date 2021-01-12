/**
 * @Author HatsuneMona
 * @Date  2020-11-12 14:53
 * @Description //TODO
 **/
package Q206

import (
	"20.leecode/leetcode"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *leetcode.ListNode
	}
	tests := []struct {
		name string
		args args
		want *leetcode.ListNode
	}{
		// TODO: Add test cases.
		{
			name: "示例",
			args: args{&leetcode.ListNode{
				Val:  1,
				Next: &leetcode.ListNode{
					Val:  2,
					Next: &leetcode.ListNode{
						Val:  3,
						Next: &leetcode.ListNode{
							Val:  4,
							Next: &leetcode.ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			}},
			want: &leetcode.ListNode{
				Val:  5,
				Next: &leetcode.ListNode{
					Val:  4,
					Next: &leetcode.ListNode{
						Val:  3,
						Next: &leetcode.ListNode{
							Val:  2,
							Next: &leetcode.ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
