package ProductofArrayExcept

import (
	stack "LeetCode/collections/mystack"
)
//algo
// 2 stacks
// first stack contains successive products of everything
// starting from last element
// second stack contains products as we go
// algo: multiply successive terms on to stack 1
// now, pop top, multiply last element onto second stack
// repeat
func productExceptSelf(nums []int) []int {
	first := stack.New()
    return []int{0}
}