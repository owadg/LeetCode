package BestTimetoBuyandSellStock

// we consider a backtracking solution
// we have two choices, the day to sell, the day to buy
// there are n-1 possibilities to buy on day k, and there are n-k possibilities to sell on

// idk what kind of algo this is, but let's try starting with a base case of the last two elem
// then, let's iterate backward, keeping track of the largest and smallest elements
func maxProfit(prices []int) int {
	// init values
	smallest := prices[len(prices) - 1]
	largest := prices[len(prices) - 1]

	maxProf := largest - smallest

	// we loop back, and everytime we find a new minimum, we use that to find our max possible profit
	for i := len(prices)-2; i >= 0; i-- {
		if prices[i] > largest {
			largest = prices[i]
		} else if prices[i] < largest {
			smallest = prices[i]
			if largest - smallest > maxProf {
				maxProf = largest - smallest
			}
		}
	}
    return maxProf
}