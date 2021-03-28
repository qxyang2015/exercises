package tree

/*
589. N 叉树的前序遍历
https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

给定一个 N 叉树，返回其节点值的 前序遍历 。
N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
*/
func preorder(root *Node) (target []int) {
	var preorderTree func(root *Node)
	preorderTree = func(root *Node) {
		if root == nil {
			return
		}
		target = append(target, root.Val)
		for _, c := range root.Children {
			preorderTree(c)
		}
	}
	preorderTree(root)
	return
}
