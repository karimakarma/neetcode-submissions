class Solution {
public:
    int findDuplicate(vector<int>& nums) {
        vector<bool> seen = {};

        for (int i = 0; i <= nums.size(); i++) {
            seen.push_back(false);
        }

        for (int i = 0; i < nums.size(); i++) {
            int cur = nums.at(i);

            if (seen.at(cur)) return cur;

            seen[cur] = true;
        }

        return 0;
    }
};

