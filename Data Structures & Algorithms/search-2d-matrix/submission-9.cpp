class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int n = matrix.at(0).size(), m = matrix.size();
        int lo = 0, hi = m * n - 1;
    
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            int c = matrix.at(mid / n).at(mid % n);

            if (c == target) return true;

            if (c > target) hi = mid - 1;
            else lo = mid + 1;
        }

        return false;
    }
};
