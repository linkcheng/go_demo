package mystr

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length <= 1 {
		return length
	}

	maxLen := 1
	left := 0
	window := make(map[byte]int)
	for right:=0; right<length; right++ {
		rightChar := s[right]
		if rightCharIndex, ok := window[rightChar]; ok && rightCharIndex > left{
			left = rightCharIndex
		} 
		maxLen = max(maxLen, right-left+1)
		// println("left=",left, "right=", right, "maxLength=", maxLen)
		window[rightChar] = right+1
	}

	return maxLen
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	size := len(strs)
	if size == 1 {
		return strs[0]
	}
	
	cmp := func (str1, str2 string) string {
		minLen := min(len(str1), len(str2))
		i := 0
		for ; i<minLen; i++ {
			if str1[i] != str2[i] { 
				break
			}
		}
		return str1[:i]
	}

	minStr := strs[0]
	for i:=1; i<size; i++ {
		minStr = cmp(minStr, strs[i])
	}
	return minStr
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

