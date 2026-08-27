impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = head;
        let mut prev = None;
        let mut i = head.take();

        // loop until i is None
        while let Some(mut node) = i {
            // save the next node
            let next = node.next.take();

            // set the node "i" to point to "prev" (initially None)
            node.next = prev.take();

            // set "prev" to point to "i"
            prev = Some(node);

            // set "i" to point to the saved next node
            i = next;
        }

        prev
    }
}
