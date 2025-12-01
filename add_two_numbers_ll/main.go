package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}

	current := dummyHead

	p := l1
	q := l2

	carry := 0

	for p != nil || q != nil || carry > 0 {
		val1 := 0
		if p != nil {
			val1 = p.Val
		}

		val2 := 0
		if q != nil {
			val2 = q.Val
		}

		sum := val1 + val2 + carry

		carry = sum / 10
		digit := sum % 10

		newNode := &ListNode{Val: digit}

		current.Next = newNode

		current = current.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	return dummyHead.Next
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

func printList(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Printf("%d -> ", temp.Val)
		temp = temp.Next
	}
	fmt.Println("nil")
}

func main() {
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 7})
	result1 := addTwoNumbers(l1, l2)
	printList(result1)

	l3 := createList([]int{9, 9, 9, 9, 9, 9, 9})
	l4 := createList([]int{9, 9, 9, 9})
	result2 := addTwoNumbers(l3, l4)
	printList(result2)
}
