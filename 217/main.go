package ContainsDuplicate

import "fmt"

func containsDuplicate(nums []int) bool {
	var unique map[int]bool = make(map[int]bool)
	for _,i := range nums {
		val := unique[i]
		if val {
			return true
		} else {
			unique[i] = true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1,1,1,3,3,4,3,2,4,2}))
}