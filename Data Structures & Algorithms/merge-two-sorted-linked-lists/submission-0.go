/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	i := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			i.Next = list1
			list1 = list1.Next
		} else {
			i.Next = list2
			list2 = list2.Next
		}
		i = i.Next
	}

	if list1 != nil {
		i.Next = list1
	} else {
		i.Next = list2
	}

	return dummy.Next
}
