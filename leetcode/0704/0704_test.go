package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Case 1",
			args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9},
			want: 4,
		},
		{
			name: "Case 2",
			args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2},
			want: -1,
		},
		{
			name: "Case 3",
			args: args{nums: []int{0, 1, 2}, target: 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search(%v, %v) = %v, want %v", tt.args.nums, tt.args.target, got, tt.want)
			}
		})
	}
}
