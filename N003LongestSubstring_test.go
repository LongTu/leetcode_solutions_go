package main

import (
	"fmt"
	"testing"
)

func TestN003LongestSubstring_lengthOfLongestSubstring(t *testing.T) {
	var n003 N003LongestSubstring
	fmt.Println("N003LongestSubstring:", n003.lengthOfLongestSubstring("abcabcbb"), n003.lengthOfLongestSubstring("eee"), n003.lengthOfLongestSubstring("dvdf"))
}