package tree

import "math"

/*
111. 二叉树的最小深度
https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。
*/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	//叶子节点是指没有子节点的节点，注意找最小的时候的结束条件
	minD := math.MaxInt32
	if root.Right != nil {
		minD = IntMin(minDepth(root.Right), minD)
	}
	if root.Left != nil {
		minD = IntMin(minDepth(root.Left), minD)
	}
	return minD + 1
}

func IntMin(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
