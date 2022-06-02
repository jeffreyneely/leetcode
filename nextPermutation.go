package main

func nextPermutation(nums []int) {
	ln := len(nums)
	if ln < 1 || nums == nil {
		return
	}

	lo := ln - 2
	for lo >= 0 && nums[lo] >= nums[lo+1] {
		lo--
	}

	if lo >= 0 {
		hi := ln - 1
		for nums[hi] <= nums[lo] {
			hi--
		}
		swap(&nums, lo, hi)
	}
	reverse(&nums, lo+1)
}

func reverse(arr *[]int, i int) {
	j := len(*arr)
	for i < j {
		swap(arr, i, j)
		i++
		j--
	}
}

func swap(arr *[]int, i, j int) {
	tmp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = tmp
}
