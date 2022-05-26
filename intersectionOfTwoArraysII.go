package main

import "sort"

func intersection(arr1 []int, arr2 []int) []int {
	res := make([]int, 0)
	i := 0
	j := 0

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			res = append(res, arr1[i])
			i++
			j++
		} else if arr1[i] > arr2[j] {
			j++
		} else if arr1[i] < arr2[j] {
			i++
		}
	}
	return res
}
