/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next	// save a copy of the Next nodes of intigers
		head.Next = prev	// make the pointer points to the previous node
		prev = head			// make the pointer the previous node
		head = next			// make the pointer the next node
	}

	return prev
}