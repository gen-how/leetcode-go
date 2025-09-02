package main

import (
	"reflect"
	"slices"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantArr []int
	}{
		// TODO: Add test cases.
		{
			name:    "Case 1",
			args:    args{nums: []int{3, 2, 2, 3}, val: 3},
			wantLen: 2,
			wantArr: []int{2, 2},
		},
		{
			name:    "Case 2",
			args:    args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			wantLen: 5,
			wantArr: []int{0, 0, 1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.wantLen {
				t.Errorf("removeElement(%v, %v) = %v, want %v", tt.args.nums, tt.args.val, got, tt.wantLen)
			}
			new_nums := tt.args.nums[:tt.wantLen]
			slices.Sort(new_nums)
			if !reflect.DeepEqual(new_nums, tt.wantArr) {
				t.Errorf("nums = %v, want %v", new_nums, tt.wantArr)
			}

		})
	}
}
