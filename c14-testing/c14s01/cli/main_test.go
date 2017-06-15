package main

import "testing"

func Test_mul(t *testing.T) {

	type args struct {
		a int
		b int
	}

	tests := []struct {
		args
		expect int
	}{
		{args{0, 0}, 0},
		{args{1, 1}, 1},
		{args{2, 2}, 4},
		{args{-2, 2}, -4},
		{args{-1, -1}, 1},
	}

	for _, tt := range tests {
		got := mul(tt.args.a, tt.args.b)
		if got != tt.expect {
			t.Errorf("got: %v, expected: %v\n", got, tt.expect)
		}
	}

}
