package main

import "math"

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func maxPathSum(root *treeNode) int {
	res := math.MinInt32
	maxGain(root, &res)
	return res
}

func maxGain(node *treeNode, m *int) int {
	if node == nil {
		return 0
	}

	lg := maxGain(node.Left, m)
	rg := maxGain(node.Right, m)

	notRoot := max(max(lg, rg)+node.Val, node.Val)
	asRoot := max(notRoot, lg+rg+node.Val)

	*m = max(*m, asRoot)
	return notRoot
}
