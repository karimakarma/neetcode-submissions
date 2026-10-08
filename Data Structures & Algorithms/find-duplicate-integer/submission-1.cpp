class Solution {
public:
    int findDuplicate(vector<int>& nums) {
        int slow = nums[0], fast = nums[0];

        do {
            slow = nums[slow];
            fast = nums[nums[fast]];
        } while (slow != fast);

        int res = nums[0];

        while (res != slow) {
            slow = nums[slow];
            res = nums[res];
        }

        return res;
    }
};
