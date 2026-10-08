class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int m = matrix.at(0).size(), n = matrix.size();
        int lo = 0, hi = n;

        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;

            if (matrix.at(mid).at(m - 1) == target) return true;
            if (matrix.at(mid).at(m - 1) > target) hi = mid;
            else lo = mid + 1;
        }

        if (lo >= n) return false;  // target exceeds all elements   

        vector<int> line = matrix.at(lo);
        lo = 0, hi = m;
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;

            if (line.at(mid) == target) return true;

            if (line.at(mid) > target) hi = mid;
            else lo = mid + 1;
        }

        return false;
    }
};
