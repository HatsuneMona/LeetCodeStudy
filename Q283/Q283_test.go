/**
 * @Author HatsuneMona
 * @Date  2020-11-19 19:42
 * @Description //TODO
 **/
package Q283

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "默认测试",
			args: args{nums: []int{0, 1, 0, 3, 12}},
		},
		{
			name: "只有一个0的错误",
			args: args{nums: []int{0}},
		},
		{
			name: "末尾是0的错误",
			args: args{nums: []int{3, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
