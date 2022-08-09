package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	x := 999999999
	for idx := 0; idx < b.N; idx++ {
		reverse(x)
	}
}
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
			name: "x = 123",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "x = -123",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "x = 120",
			args: args{x: 120},
			want: 21,
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
