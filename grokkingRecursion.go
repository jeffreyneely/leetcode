package main

import (
	"math"
	"math/rand"
	"time"
)

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

func QSort(arr *[]int, i, j int) {
	if i < j {
		pv := rndPartition(arr, i, j)
		QSort(arr, i, pv-1)
		QSort(arr, pv+1, j)
	}
}

func rndPartition(arr *[]int, i, j int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	rv := rnd.Int()
	pvt := i + rv%(j-i+1)
	swp(arr, j, pvt)
	return partition(arr, i, j)
}

func partition(arr *[]int, lo, hi int) int {
	idx := lo
	pv := hi

	for i := lo; i < hi; i++ {
		if (*arr)[i] < (*arr)[pv] {
			swp(arr, i, idx)
			idx++
		}
	}
	swp(arr, pv, idx)
	return idx
}

func swp(arr *[]int, i, j int) {
	tmp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = tmp
}
