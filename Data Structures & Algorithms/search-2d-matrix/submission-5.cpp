class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int n = matrix.at(0).size(), m = matrix.size();
        int lo = 0, hi = m * n - 1;
    
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;

            if (matrix.at(mid / n).at(mid % n) == target) return true;

            if (matrix.at(mid / n).at(mid % n) > target) hi = mid - 1;
            else lo = mid + 1;
        }

        return false;
    }
};
