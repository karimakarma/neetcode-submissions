from collections import Counter
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t) : return False

        s_map, t_map = Counter(s), Counter(t)

        return s_map == t_map