/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	nodeMap := map[*Node]*Node{}

	for i := head; i != nil; i = i.Next {
		nodeMap[i] = &Node{Val: i.Val}
	}

	for i := head; i != nil; i = i.Next {
		curr := nodeMap[i]

		curr.Next = nodeMap[i.Next]
		curr.Random = nodeMap[i.Random]
	}

	return nodeMap[head]
}
