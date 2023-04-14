package main

import (
	"fmt"
)

// algo
// 2 stacks
// first stack contains successive products of everything
// starting from last element - aka last element is first on stack
// second stack contains products as we go
// algo: multiply successive terms on to stack 1
// now, pop top, multiply last element onto second stack
// repeat
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	lastIndex := len(nums) - 1

	backwardsSuccessiveProducts := make([]int, len(nums))
	backwardsSuccessiveProducts[lastIndex] = nums[lastIndex]
	for i := lastIndex - 1; i >= 0; i-- {
		backwardsSuccessiveProducts[i] = backwardsSuccessiveProducts[i+1] * nums[i]
	}

	forwardsSuccessiveProducts := make([]int, len(nums))
	forwardsSuccessiveProducts[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		forwardsSuccessiveProducts[i] = forwardsSuccessiveProducts[i-1] * nums[i]
	}

	result[0] = backwardsSuccessiveProducts[1]
	result[lastIndex] = forwardsSuccessiveProducts[lastIndex-1]
	for i := 1; i < lastIndex; i++ {
		result[i] = backwardsSuccessiveProducts[i+1] * forwardsSuccessiveProducts[i-1]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))

	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums2))
}
