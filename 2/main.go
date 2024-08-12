package main

import "fmt"

type ListNode struct {
    Val int
 	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sum := l1.Val + l2.Val

    // base case and also a nice thing to make it easier to do other calculations (but more objects are created...)
    // but makes carrying easy and covers all cases
    if (l1.Next == nil && l2.Next == nil) {
        if (sum > 9){
            return &ListNode{sum - 10, &ListNode{1, nil}}
        } else {
            return &ListNode{sum, nil}
        }
    } else if (l1.Next == nil){
        l1.Next = &ListNode{0, nil}
    } else if (l2.Next == nil){
        l2.Next = &ListNode{0, nil}
    }

	// encodes the carrying logic
    if (sum > 9) {
            // carry the one (it doesn't say we can't modify the original linked list)
            // we know both are present, so it doesn't matter which one we use
            l1.Next.Val += 1
            return &ListNode{sum - 10, addTwoNumbers(l1.Next, l2.Next)}
    } else {
            return &ListNode{sum, addTwoNumbers(l1.Next, l2.Next)}
    }
}

func main() {
	fmt.Println(addTwoNumbers(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}))
}