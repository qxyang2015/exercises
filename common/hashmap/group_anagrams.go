package hashmap

/*
49. 字母异位词分组
https://leetcode-cn.com/problems/group-anagrams/

给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
*/

func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		val := [26]int{}
		for _, b := range str {
			val[b-'a']++
		}
		mp[val] = append(mp[val], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
