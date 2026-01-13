package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, len(nums) * 2 + 1)
	fmt.Println("Result:", nums)
}

func rotate(nums []int, k int)  {
	tmpNums := make([]int, len(nums))
	copy(tmpNums, nums)
	length := len(nums)
	offset := k % length
	position := 0;
	for i := offset; i < length; i++ {
		nums[i] = tmpNums[position]
		position++
	}
	for i := range offset {
		nums[i] = tmpNums[position]
		position++
	}
}
