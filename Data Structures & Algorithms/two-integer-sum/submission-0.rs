use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut seen = HashMap::new();

        for (i, cur) in nums.iter().enumerate() {
            let x = target - cur;

            if seen.contains_key(&x) {
                return vec![seen[&x] as i32, i as i32];
            }

            seen.insert(cur, i);
        }

        vec![]
    }
}
