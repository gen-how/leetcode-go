package leetcode

func lengthOfLongestSubstring(s string) int {
	// This assumes that `s` contains only ASCII characters.
	charIndices := [256]int{0}
	maxLen := 0
	l := 0
	for r := 0; r < len(s); r++ {
		// Treats index value 0 as not present.
		if idx := charIndices[s[r]]; idx-1 >= l {
			l = idx
		}
		charIndices[s[r]] = r + 1
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}

/* Uses sliding window and hash map which stores the index of characters.

func lengthOfLongestSubstring(s string) int {
	charIndices := make(map[byte]int)
	maxLen := 0
	l := 0
	for r := 0; r < len(s); r++ {
		if idx, exists := charIndices[s[r]]; exists && idx >= l {
			l = idx + 1
		}
		charIndices[s[r]] = r
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}
*/

/* Uses sliding window and hash map which stores the frequency of characters.

func lengthOfLongestSubstring(s string) int {
	charFreq := make(map[byte]int)
	maxLen := 0
	l := 0
	for r := 0; r < len(s); r++ {
		for cnt, exists := charFreq[s[r]]; exists && cnt > 0; cnt = charFreq[s[r]] {
			charFreq[s[l]]--
			l++
		}
		charFreq[s[r]]++
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}
*/

/* Uses sliding window and Set

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(key T) {
	s[key] = struct{}{}
}

func (s Set[T]) Has(key T) bool {
	_, exists := s[key]
	return exists
}

func (s Set[T]) Remove(key T) {
	delete(s, key)
}

func lengthOfLongestSubstring(s string) int {
	charSet := make(Set[byte])
	maxLen := 0
	l := 0
	for r := 0; r < len(s); r++ {
		for charSet.Has(s[r]) {
			charSet.Remove(s[l])
			l++
		}
		charSet.Add(s[r])
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}
*/
