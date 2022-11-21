package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))

}

func removeDuplicates(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if len(nums) == 0 || i == len(nums)-1 {
			if nums[len(nums)-1] != nums[len(nums)-2] {
				nums[index] = nums[i]
			}
			break
		}
		if nums[i] == nums[i+1] {
			continue
		} else {
			nums[index] = nums[i]
		}
	}

	return index + 1
}
