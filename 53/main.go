package main

import "fmt"

func maxSubArray(nums []int) int {

	
	//brainstorming
	// where do we start?
	// where do we end?
	// taking inspiration from #238, what if we had a "running sum" that goes forwards and backwards
	// once we find the start, we can just loop through the rest of the array sum to see if it gets bigger at all
	// this does not work

	// solution space is O(n^2) so a brute force approach is O(n^2)
	// we need at least an O(n) algo

	// we can try a divide and conquer approach
	// maximum subarray of array of length 1 is itself
	// when merging, the subarray is either both combined, or one or the other, depending on the sum
	// if the sum is negative, we choose the least negative one
	// if the sum is positive, it's both

	result := make([]int, len(nums))
	
	var sum int = 0
	for _, num := range result {
		fmt.Println(num)
		sum += num
	}

	fmt.Println("input", nums)
	fmt.Println("resulting subarray", result)
	fmt.Println("sum", sum)
	return sum
}

func main() {
	maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})
	maxSubArray([]int{5,4,-1,7,8})
}
