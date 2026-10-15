impl Solution {
    pub fn remove_nth_from_end(mut head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        if let Some(ref mut head) = head {
            let start = &raw mut **head;
            let mut slow = start;
            let mut fast = start;

            unsafe {
                for _ in 0..n {
                    if let Some(ref mut next) = (*fast).next {
                        fast = &raw mut **next;
                    } else { fast = std::ptr::null_mut() }
                }

                if !fast.is_null() { while let Some(ref mut next) = (*fast).next {
                    fast = &raw mut **next;
                    if let Some(ref mut next) = (*slow).next {
                        slow = &raw mut **next;
                    } 
                }}

                if fast.is_null() || (*slow).next.is_none() {
                    return head.next.take();
                } else {
                    if let Some(mut next) = (*slow).next.take() {
                        (*slow).next = next.next.take();
                    }
                }
            }

        }

        head
    }
}
