/**
 * @Author HatsuneMona
 * @Date  2020-11-14 19:05
 * @Description //TODO
 **/
package Q7

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "正数",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "负数",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "正数溢出",
			args: args{x: 2147483647},
			want: 0,
		},
		{
			name: "负数溢出",
			args: args{x: -2147483648},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
