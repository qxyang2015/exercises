package array

/*
88. 合并两个有序数组
https://leetcode-cn.com/problems/merge-sorted-array/

给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx1, idx2, idx := m-1, n-1, m+n-1
	for idx1 >= 0 && idx2 >= 0 {
		if nums1[idx1] > nums2[idx2] {
			nums1[idx] = nums1[idx1]
			idx1--
		} else {
			nums1[idx] = nums2[idx2]
			idx2--
		}
		idx--
	}
	for idx1 >= 0 {
		nums1[idx] = nums1[idx1]
		idx--
		idx1--
	}
	for idx2 >= 0 {
		nums1[idx] = nums2[idx2]
		idx--
		idx2--
	}
}
