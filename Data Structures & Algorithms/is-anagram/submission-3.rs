impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }

        let mut map = [0; 26];

        for c in s.bytes() {
            map[(c - b'a') as usize] += 1;
        }

        for c in t.bytes() {
            let cur = &mut map[(c - b'a') as usize];

            if cur == &0 {
                return false;
            }

            *cur -= 1;
        }

        for i in map {
            if i > 0 {
                return false;
            }
        }
        true
    }
}
