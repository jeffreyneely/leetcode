package main

import "sort"

func threeSum(arr []int) [][]int {
	r := make([][]int, 0)
	if len(arr) < 3 {
		return r
	}

	sort.Ints(arr)

	for i := 0; i < len(arr) && arr[i] <= 0; i++ {
		if i == 0 || arr[i] != arr[i-1] {
			left := i + 1
			right := len(arr)
			for left < right {
				sum := arr[i] + arr[left] + arr[right]
				if sum == 0 {
					r = append(r, []int{arr[i], arr[left], arr[right]})
					left++
					right--

					for left < right && arr[left] == arr[left-1] {
						left++
					}

					for left < right && arr[right] == arr[right+1] {
						right--
					}
				} else if sum > 0 {
					right--
				} else if sum < 0 {
					left++
				}
			}
		}
	}

	return r
}
