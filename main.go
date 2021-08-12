package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println("done!")
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	pos := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[pos] {
			nums[pos+1] = nums[i]
			pos++
		}
	}
	return pos
}
