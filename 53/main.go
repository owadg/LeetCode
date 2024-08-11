package MaximumSumOfSubarray

func maxSubArray(nums []int) int {
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
