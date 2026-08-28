impl Solution {
    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut head = head;

        let mut length = 0;

        let mut i = head.as_mut();

        while i.is_some() {
            i = i.unwrap().next.as_mut();
            length += 1;
        }

        let mut i = head.as_mut();

        if length <= 1 {
            return None;
        }

        if length - n == 0 {
            return head.unwrap().next;
        }

        for _ in 0..length - n - 1 {
            i = i.unwrap().next.as_mut();
        }

        let node = i.unwrap();

        node.next = node.next.take().unwrap().next;

        head
    }
}
