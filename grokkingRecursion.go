package main

import "math"

func MaxValue(arr []int) int {
	m := math.MinInt32
	findMax(arr, &m)
	return m
}

func findMax(arr []int, m *int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	*m = max(arr[0], *m)
	return findMax(arr[1:], m)
}

func CountElem(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return 1 + CountElem(arr[1:])
}

func SumElem(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + SumElem(arr[1:])
}
