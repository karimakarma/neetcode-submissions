class Solution {
public:
    int search(vector<int>& nums, int target) {
        int lo = 0, hi = nums.size();

        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            int cur = nums[mid];

            if (cur == target) return mid;

            if (cur < target) {
                lo = mid + 1;
            } else {
                hi = mid;
            }
        }

        return -1;
    }
};
