impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut group: HashMap<[i32; 26], Vec<String>> = HashMap::new();

        for s in strs {
            let mut key = [0; 26];

            for c in s.bytes() {
                key[(c - b'a') as usize] += 1;
            }

            if group.contains_key(&key) {
                group.get_mut(&key).unwrap().push(s);
                continue;
            }

            group.insert(key, vec![s]);
        }

        group.into_values().collect()
    }
}