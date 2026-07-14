/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	fHalf := head
	middle, i := fHalf, fHalf
	for i != nil && i.Next != nil {
		middle = middle.Next
		i = i.Next.Next
	}
	sHalf := middle.Next
	middle.Next = nil

	var prev *ListNode
	for sHalf != nil {
		next := sHalf.Next
		sHalf.Next = prev
		prev = sHalf
		sHalf = next
	}

	i = fHalf
	for prev != nil {
		nextF := i.Next
		nextS := prev.Next
		i.Next = prev
		prev.Next = nextF
		i = nextF
		prev = nextS
	}
}
