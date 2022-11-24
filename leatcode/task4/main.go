package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 4, 6}
	fmt.Println(SearchInsert(nums, 2))
}

func SearchInsert(nums []int, target int) int {
	var midl int
	bottom := 0
	top := len(nums) - 1

	for top >= bottom {
		midl = (top + bottom) / 2

		if nums[midl] == target {
			return midl
		}
		if nums[midl] > target {
			top = midl - 1
		} else {
			bottom = midl + 1
		}
	}
	if nums[midl] < target && nums[top] > target {
		return top
	} else {
		return bottom
	}
}
