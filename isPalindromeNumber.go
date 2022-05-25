package main

func IsPalindromeNum(x int) bool {
	// Cannot be palindome if negative or
	// if the ones spot is 0 and x itself isn't 0
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}

	var a []int
	// append ones, tens, hundreds spot etc
	for i := x; i != 0; i /= 10 {
		a = append(a, i%10)
	}

	// compare arr i to arr[len(arr) - 1]
	j := int8(len(a) - 1)
	for _, v := range a {
		if v != a[j] {
			return false
		}
		j--
	}
	return true
}
