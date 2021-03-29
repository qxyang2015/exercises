package tree

/*
429. N 叉树的层序遍历
https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/

给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
*/
func levelOrder(root *Node) (target [][]int) {
	if root == nil {
		return
	}
	var dfs func(root *Node, level int)
	dfs = func(root *Node, level int) {
		if root == nil {
			return
		}
		if len(target) <= level {
			target = append(target, []int{})
		}
		target[level] = append(target[level], root.Val)
		for _, c := range root.Children {
			dfs(c, level+1)
		}
	}
	dfs(root, 0)
	return
}
