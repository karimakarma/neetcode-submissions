/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
		slow = slow.Next
	}

	return false
}
