package main

func nextPermutation(nums []int) {
	l := len(nums)
	if l < 1 {
		return
	}

	lo := l - 2
	for lo >= 0 && nums[lo] >= nums[lo+1] {
		lo--
	}

	if lo > 0 {
		hi := l - 1
		for nums[hi] <= nums[lo] {
			hi--
		}
		swap(&nums, lo, hi)
	}
	reverse(&nums, lo+1, l-1)
}

func swap(arr *[]int, i int, j int) {
	tmp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = tmp
}

func reverse(arr *[]int, i int, j int) {
	for i < j {
		swap(arr, i, j)
		i++
		j--
	}
}
