class Solution {
    public:
        ListNode* removeNthFromEnd(ListNode* head, int n) {
            ListNode *slow = head, *fast = head;

            for (int i = 0; i < n && fast; i++) {
                fast = fast->next;
            }

            if (fast) while (fast->next) {
                fast = fast->next;
                slow = slow->next;
            }

            if (!slow->next) head = head->next;
            else if (!fast) head = head->next;
            else slow->next = slow->next->next;
            

            return head;
        }
};

