class Solution {
    public:
        int findMin(vector<int> &nums) {
            int lo = 0, 
                hi = nums.size() - 1, 
                res = nums[0];

            while (lo <= hi) {
                if (nums[lo] < nums[hi]) {
                    res = min(nums[lo], res); break;
                }

                int mid = lo + (hi - lo) / 2;
                res = min(nums[mid], res);

                if (nums[lo] <= nums[mid]) {
                    lo = mid + 1;
                } else {
                    hi = mid - 1;
                }
            }

            return res;
        }
};
