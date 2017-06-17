package main

import "testing"

func Test_mul(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0test", args{0, 0}, 0},
		{"positive=1", args{1, 1}, 1},
		{"positive=2", args{2, 2}, 4},
		{"negative=1", args{-2, 2}, -4},
		{"negative=2", args{-1, -1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mul(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0test", args{0, 0}, 0},
		{"positive=1", args{1, 1}, 2},
		{"positive=2", args{2, 2}, 4},
		{"negative=1", args{-2, 2}, 0},
		{"negative=2", args{-1, -1}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
