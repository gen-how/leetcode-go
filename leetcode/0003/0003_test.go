package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Case 1",
			args: args{s: "abcabcbb"},
			want: 3, // "abc"
		},
		{
			name: "Case 2",
			args: args{s: "bbbbb"},
			want: 1, // "b"
		},
		{
			name: "Case 3",
			args: args{s: "pwwkew"},
			want: 3, // "wke"
		},
		{
			name: "Case 4",
			args: args{s: "abcad"},
			want: 4, // "bcad"
		},
		{
			name: "Case 5",
			args: args{s: "abba"},
			want: 2, // "ab"
		},
		{
			name: "Case 6",
			args: args{s: " "},
			want: 1, // " "
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
