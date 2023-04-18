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
	// I dont think this strategy works, because it assumes they can always be added

	//maybe do above strat, but with an extra bool saying if it can be merged or if it has to be taken by itself
	// algo
	// recurse into base case
	// get sum
	// merge - if greater is on

	//okay lets just think of an iterative algo first

	// okay so I spoiled myself with Kadane's algorithm
	// lets justify it in my head
	// lets say we have [0,-3,1,1]
	// at the nth index, is it true that the largest maximum subarray is either to the left, or at this index?
	// well, yes, of course.
	//


	
	var globalMaxSum int = nums[0]

	localMaxSum := 0
	for i := range nums {
		// these statements give the maximum subarray at the ith index
		if localMaxSum > 0 {
			localMaxSum += nums[i]
		} else {
			localMaxSum = nums[i]
		}

		// these statements record the largest sum
		if globalMaxSum < localMaxSum {
			globalMaxSum = localMaxSum
		}
	}

	return globalMaxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(maxSubArray([]int{5,4,-1,7,8}))
	fmt.Println(maxSubArray([]int{0,-3,1,1}))
}
