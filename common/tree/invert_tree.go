package tree

/*
226. 翻转二叉树
https://leetcode-cn.com/problems/invert-binary-tree/

翻转一棵二叉树。
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	right := invertTree(root.Right)
	left := invertTree(root.Left)
	root.Right, root.Left = left, right
	return root
}
