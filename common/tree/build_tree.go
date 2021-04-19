package tree

/*
105. 从前序与中序遍历序列构造二叉树
https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

根据一棵树的前序遍历与中序遍历构造二叉树。
注意:你可以假设树中没有重复的元素。
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	idx := 0
	for ; idx < len(inorder); idx++ {
		if preorder[0] == inorder[idx] {
			break
		}
	}
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
