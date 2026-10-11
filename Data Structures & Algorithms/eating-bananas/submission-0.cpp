#include <algorithm>

class Solution {
public:
    int minEatingSpeed(vector<int>& piles, int h) {
        int lo = 1, hi = *std::max_element(piles.begin(), piles.end()), min = hi;

        while (lo <= hi) {
            int k = lo + (hi - lo) / 2;
            long long hours = 0;

            for (int p : piles) hours += (p + k - 1LL) / k;

            if (hours <= h) {
                min = k;

                hi = k - 1;
            } else {
                lo = k + 1;
            }
        }

        return min;
    }
};
