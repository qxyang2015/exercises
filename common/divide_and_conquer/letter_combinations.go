package divide_and_conquer

/*
17. 电话号码的字母组合
https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) (ans []string) {
	if len(digits) <= 0 {
		return ans
	}
	var dfs func(digits string, idx int, combination string)
	dfs = func(digits string, idx int, combination string) {
		if idx >= len(digits) {
			ans = append(ans, combination)
			return
		}
		letters := phoneMap[string(digits[idx])]
		for _, c := range letters {
			dfs(digits, idx+1, combination+string(c))
		}
	}
	dfs(digits, 0, "")
	return
}
