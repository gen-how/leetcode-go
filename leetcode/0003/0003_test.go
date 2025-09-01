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
			name: "Case \"abcabcbb\"",
			args: args{s: "abcabcbb"},
			want: 3, // "abc"
		},
		{
			name: "Case \"bbbbb\"",
			args: args{s: "bbbbb"},
			want: 1, // "b"
		},
		{
			name: "Case \"pwwkew\"",
			args: args{s: "pwwkew"},
			want: 3, // "wke"
		},
		{
			name: "Case \"abcad\"",
			args: args{s: "abcad"},
			want: 4, // "bcad"
		},
		{
			name: "Case \"abba\"",
			args: args{s: "abba"},
			want: 2, // "ab"
		},
		{
			name: "Case \" \"",
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
