use std::collections::HashMap;

impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        // step 1: getthe numbers and frequencies
        let mut map = HashMap::new();
        for n in nums.iter() {
            *map.entry(*n).or_insert(0) += 1;
        }

        // step 2: make a vec of vec containing the numbers with the frequencies as its index
        let mut freq: Vec<Vec<i32>> = vec![vec![]; nums.len() + 1]; // set to the length of the nums vec
        // because no numbers can appear more than the total numbers of elements

        for (n, f) in map {
            freq[f].push(n);
        }

        // step 3: pop the freq vec until k numbers is retrieved
        let mut res: Vec<i32> = Vec::new();
        while res.len() < k as usize {
            let mut vec_pop = freq.pop().unwrap();

            while let Some(n) = vec_pop.pop() {
                res.push(n);
            }
        }

        res
    }
}
