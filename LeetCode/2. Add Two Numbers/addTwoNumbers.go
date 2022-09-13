package main

import "fmt"

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
 Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Printf("2. Add Two Numbers\n")
	var l1 = []int{2, 4, 3}
	var l2 = []int{5, 6, 4}
	addTwoNumbers(l1, l2)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Printf("addTwoNumbers func \n")

}
