package stack

/*
20. 有效的括号
https://leetcode-cn.com/problems/valid-parentheses/

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
    左括号必须用相同类型的右括号闭合。
    左括号必须以正确的顺序闭合。
*/

/*
1、string range 为rune, str[i]为byte
2、如果当前元素属于右括号：1、如果栈为空、2、或者栈顶与当前元素不匹配 则不是有效括号
*/

var pairs = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if findVal, ok := pairs[s[i]]; ok {
			if len(stack) == 0 || findVal != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
