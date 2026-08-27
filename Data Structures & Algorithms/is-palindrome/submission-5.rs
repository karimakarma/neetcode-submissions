impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        let s = s.trim().chars().collect::<Vec<char>>();

        if s.len() <= 1 {
            return true;
        }
        
        let (mut l, mut r) = (0usize, s.len() - 1);
        while l < r {
            if s[l].is_whitespace() || !s[l].is_alphanumeric() {
                l += 1;
                continue;
            }
            

            if s[r].is_whitespace() || !s[r].is_alphanumeric() {
                r -= 1;
                continue;
            }

            if !s[l].eq_ignore_ascii_case(&s[r]) {
                return false;
            } else {
                l += 1;
                r -= 1;
            }
        }

        true
    }

}
