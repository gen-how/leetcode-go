package main

import (
	"testing"
)

// intsToList convert []int to ListNode
func intsToList(ints []int) *ListNode {
	if len(ints) == 0 {
		return nil
	}

	head := &ListNode{}
	tail := head
	for _, val := range ints {
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
	}
	return head.Next
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "Case 1",
			args: args{l1: intsToList([]int{2, 4, 3}), l2: intsToList([]int{5, 6, 4})},
			want: intsToList([]int{7, 0, 8}),
		},
		{
			name: "Case 2",
			args: args{l1: intsToList([]int{0}), l2: intsToList([]int{0})},
			want: intsToList([]int{0}),
		},
		{
			name: "Case 3",
			args: args{l1: intsToList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}), l2: intsToList([]int{5, 6, 4})},
			want: intsToList([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			for g, w := got, tt.want; (g != nil) || (w != nil); g, w = g.Next, w.Next {
				if g.Val != w.Val {
					t.Errorf("got %d, want %d", g.Val, w.Val)
				}
			}
		})
	}
}
