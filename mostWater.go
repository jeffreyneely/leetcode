package main

func mostWater(height []int) int {
	a := 0               // max area
	l := 0               // left p
	r := len(height) - 1 // right p

	for l < r {
		w := r - l // width
		a = maxA(a, minH(l, r)*w)

		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return a
}

func minH(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxA(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
