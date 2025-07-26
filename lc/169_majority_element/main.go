package main

import (
	"fmt"
)

func main() {
	println("Hello")
	var nums = []int{2, 2, 1, 1, 1, 2, 2}
	var res = majorityElement(nums)
	fmt.Println("Result:", res)
}
func majorityElement(nums []int) int {
	var m = make(map[int]int)
	var b = 0

	if len(nums) == 0 {
		return nums[0]
	}

	for _, n := range nums {
		m[n] += 1
		if m[b] < m[n] {
			b = n
		}
	}
	return b
}
