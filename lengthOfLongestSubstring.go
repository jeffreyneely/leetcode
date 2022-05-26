package main

func lengthOfLongestSubstring(s string) int {
	var e = len(s)
	dict := make(map[byte]int)
	res := -1
	d := -1

	for i := 0; i < e; i++ {
		if v, ok := dict[s[i]]; !ok {
			// if the value is not already in the dictionary,
			// add the index of the value; iterate our window.
			dict[s[i]] = i
			d++
		} else {
			d = min(i-v-1, d+1)
			dict[s[i]] = i
		}
		res = max(res, d)
	}
	return res + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
