/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	i := head

	for i != nil {
		next := i.Next
		i.Next, prev, i = prev, i, next
	}

	return prev

}
