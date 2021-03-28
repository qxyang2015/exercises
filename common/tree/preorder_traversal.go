package tree

/*
144. 二叉树的前序遍历
https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
*/

func preorderTraversal(root *TreeNode) (target []int) {
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		target = append(target, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)
	return
}
