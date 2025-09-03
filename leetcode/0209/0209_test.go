package main

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Case 1",
			args: args{target: 7, nums: []int{2, 3, 1, 2, 4, 3}},
			want: 2, // [4, 3]
		},
		{
			name: "Case 2",
			args: args{target: 4, nums: []int{1, 4, 4}},
			want: 1, // [4]
		},
		{
			name: "Case 3",
			args: args{target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}},
			want: 0, // []
		},
		{
			name: "Case 4",
			args: args{target: 1000000000, nums: func() []int {
				result := [100000]int{}
				for i := range result {
					result[i] = 10000
				}
				return result[:]
			}()},
			want: 100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen(%v, %v) = %v, want %v", tt.args.target, tt.args.nums, got, tt.want)
			}
		})
	}
}
