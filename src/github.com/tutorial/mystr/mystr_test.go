package mystr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	println("=========test lengthOfLongestSubstring=========")
	tests := []struct{
		in string
		expected int
	}{
		{"abcabcbb", 3}, // "abc"
		{"bbbbb", 1},  // "abc"
		{"pwwkew", 3},  // "wke"
		{"abba", 2},  // "ab"
	}

	for _, tt := range tests {
        actual := lengthOfLongestSubstring(tt.in)
        assert.Equal(t, tt.expected, actual)
    }
}

func TestLongestCommonPrefix(t *testing.T) {
	println("=========test longestCommonPrefix=========")
	tests := []struct{
		in []string
		expected string
	}{
		{[]string{"flower","flow","flight"}, "fl"},
		{[]string{"dog","racecar","car"}, ""},
	}

	for _, tt := range(tests) {
		actual := longestCommonPrefix(tt.in)
		assert.Equal(t, tt.expected, actual)
	}
}
