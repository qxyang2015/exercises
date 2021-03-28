package tree

/*
590. N 叉树的后序遍历
https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/

给定一个 N 叉树，返回其节点值的 后序遍历 。
N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
*/

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) (target []int) {
	var postorderTree func(root *Node)
	postorderTree = func(root *Node) {
		if root == nil {
			return
		}
		for _, c := range root.Children {
			postorderTree(c)
		}
		target = append(target, root.Val)
	}
	postorderTree(root)
	return
}
