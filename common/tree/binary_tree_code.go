package tree

import (
	"strconv"
)

/*
297. 二叉树的序列化与反序列化
https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
*/

type Codec struct {
	str string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var codeSerialize func(node *TreeNode)
	codeSerialize = func(node *TreeNode) {
		if node == nil {
			c.str += "null,"
			return
		}
		c.str += strconv.Itoa(node.Val) + ","
		codeSerialize(node.Left)
		codeSerialize(node.Right)
	}
	codeSerialize(root)
	return c.str
}

//// Deserializes your encoded data to tree.
//func (c *Codec) deserialize(data string) *TreeNode {
//	strList := strings.Split(data, ",")
//}
