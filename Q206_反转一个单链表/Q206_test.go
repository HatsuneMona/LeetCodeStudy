/**
 * @Author HatsuneMona
 * @Date  2020-11-12 14:53
 * @Description //TODO
 **/
package Q206_反转一个单链表

import (
	"20.leecode/leetcodeType"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *leetcodeType.ListNode
	}
	tests := []struct {
		name string
		args args
		want *leetcodeType.ListNode
	}{
		// TODO: Add test cases.
		{
			name: "示例",
			args: args{&leetcodeType.ListNode{
				Val: 1,
				Next: &leetcodeType.ListNode{
					Val: 2,
					Next: &leetcodeType.ListNode{
						Val: 3,
						Next: &leetcodeType.ListNode{
							Val: 4,
							Next: &leetcodeType.ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			}},
			want: &leetcodeType.ListNode{
				Val: 5,
				Next: &leetcodeType.ListNode{
					Val: 4,
					Next: &leetcodeType.ListNode{
						Val: 3,
						Next: &leetcodeType.ListNode{
							Val: 2,
							Next: &leetcodeType.ListNode{
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
