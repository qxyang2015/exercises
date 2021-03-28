package hashmap

/*
242. 有效的字母异位词
https://leetcode-cn.com/problems/valid-anagram/

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
*/
func isAnagram(s string, t string) bool {
	if len([]rune(s)) != len([]rune(t)) {
		return false
	}
	hashMap := make(map[rune]int)
	for _, v := range []rune(s) {
		hashMap[v]++
	}
	for _, v := range []rune(t) {
		hashMap[v]--
		if hashMap[v] < 0 {
			return false
		}
	}
	return true
}
