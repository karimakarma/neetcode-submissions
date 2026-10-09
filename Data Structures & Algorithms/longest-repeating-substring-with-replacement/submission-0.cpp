class Solution {
public:
    int characterReplacement(string s, int k) {
        int seen[26];
        for (int i = 0; i < 26; i++) {
            seen[i] = 0;
        }

        int l = 0, r = 0;
        int frq = 0;
        int res = 0;
        for (; r < s.size(); r++) {
            int cl = int(s[l]) - int('A');
            int cr = int(s[r]) - int('A');

            seen[cr]++;

            if (seen[cr] > frq) {
                frq = seen[cr];
            }

            int len = r - l + 1;

            if (len > res && len - frq <= k) {
                res = len;
            } else {
                l++;
                seen[cl]--;
            }
        }

        return res;
    }
};
