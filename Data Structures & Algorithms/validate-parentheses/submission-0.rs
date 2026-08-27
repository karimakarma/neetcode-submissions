impl Solution {
    pub fn is_valid(s: String) -> bool {
        let map = HashMap::from([('(', ')'), ('[', ']'), ('{', '}')]);

        let mut stack = Vec::<char>::new();

        for ch in s.chars() {
            if map.contains_key(&ch) {
                stack.push(*map.get(&ch).unwrap());
            } else if stack.pop().unwrap_or('.') != ch {
                return false;
            }
        }

        stack.is_empty()
    }
}
