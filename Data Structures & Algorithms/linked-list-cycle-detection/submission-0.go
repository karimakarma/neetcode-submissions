/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	seenNode := map[*ListNode]bool{}

	for head != nil {
		if _, ok := seenNode[head]; ok {
			return true
		}

		seenNode[head] = true
		head = head.Next
	}

	return false

}
