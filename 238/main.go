package main

import (
	"fmt"

	stack "github.com/owadg/LeetCode/collections"
)

// algo
// 2 stacks
// first stack contains successive products of everything
// starting from last element
// second stack contains products as we go
// algo: multiply successive terms on to stack 1
// now, pop top, multiply last element onto second stack
// repeat
func productExceptSelf(nums []int) []int {

	return []int{0}
}

func main() {
	var one, two, three, four, five int = 1, 2, 3, 4, 5
	var myStack stack.Stack[int] = *stack.NewStack(one)
	fmt.Println("stack created with val of", myStack.Peek(), " on top")

	myStack.Push(two)
	fmt.Println("pushed ", myStack.Peek(), " on top")

	myStack.Push(three)
	fmt.Println("pushed ", myStack.Peek(), " on top")

	myStack.Push(four)
	fmt.Println("pushed ", myStack.Peek(), " on top")

	myStack.Push(five)
	fmt.Println("pushed ", myStack.Peek(), " on top")
}
