impl Solution {
    pub fn encode(strs: Vec<String>) -> String {
        let mut res = String::new();

        for s in strs {
            res += &format!("{}#{s}", s.len());
        }

        res
    }

    pub fn decode(s: String) -> Vec<String> {
        let mut res = Vec::<String>::new();
        let mut chars = s.chars();

        while let Some(mut ch) = chars.next() {
            let mut n = String::new();

            while ch != '#' {
                n.push(ch);
                ch = chars.next().unwrap();
            }

            let len: usize = n.parse().unwrap();

            res.push(chars.by_ref().take(len).collect());
        }

        res
    }
}
