package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	if head == nil {
		return 0
	}
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	reversedHead := reverse(slow)
	maxSum := 0
	for ; head != nil && reversedHead != nil; head, reversedHead = head.Next, reversedHead.Next {
		maxSum = max(head.Val+reversedHead.Val, maxSum)
	}
	return maxSum
}

func reverse(head *ListNode) *ListNode {
	var prev, curr *ListNode = nil, head

	for curr != nil {
		prev, curr, curr.Next = curr, curr.Next, prev
	}
	return prev
}
