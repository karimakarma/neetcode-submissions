/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for i := head; i != nil; i = i.Next.Next {
		next := i.Next
		clone := &Node{Val: i.Val, Next: next}
		i.Next = clone
	}

	for i := head; i != nil; i = i.Next.Next {
		if i.Random != nil {
			i.Next.Random = i.Random.Next
		}
	}

	cloned := head.Next
	j := cloned
	for i := head; i != nil; i = i.Next {
		i.Next = i.Next.Next

		if j.Next != nil {
			j.Next = j.Next.Next
		} else {
			j.Next = nil
		}

		j = j.Next
	}

	return cloned
}
