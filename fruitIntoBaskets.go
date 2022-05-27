package main

func totalTrees(trees []int) int {
	m := make(map[int]int)
	lo, hi, res := 0, 0, 0

	for hi < len(trees) {
		m[trees[hi]]++

		if len(m) > 2 {
			m[trees[lo]]--
			if m[trees[lo]] == 0 {
				delete(m, trees[lo])
			}
			lo++
		}

		if hi-lo+1 > res {
			res = hi - lo + 1
		}
		hi++
	}

	return res
}
