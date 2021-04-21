package tree

/*
515. 在每个树行中找最大值
https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/

您需要在二叉树的每一行中找到最大的值。
*/
func largestValues(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	p := []*TreeNode{root}
	for len(p) > 0 {
		var q []*TreeNode
		max := p[0].Val
		for i := 0; i < len(p); i++ {
			if p[i].Val > max {
				max = p[i].Val
			}
			if p[i].Left != nil {
				q = append(q, p[i].Left)
			}
			if p[i].Right != nil {
				q = append(q, p[i].Right)
			}
		}
		ans = append(ans, max)
		p = q
	}
	return
}
