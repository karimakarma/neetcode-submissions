impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut group: HashMap<[i32; 26], Vec<String>> = HashMap::new();

        for s in strs {
            let mut key = [0; 26];

            for c in s.bytes() {
                key[(c - b'a') as usize] += 1;
            }
            group.entry(key).or_default().push(s);

        }

        group.into_values().collect()
    }
}