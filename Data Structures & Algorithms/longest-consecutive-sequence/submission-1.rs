impl Solution {
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        if nums.is_empty() {
            return 0;
        }

        let set: HashSet<i32> = nums.iter().cloned().collect();

        let mut longest = 0;

        for i in set.iter() {
            if !set.contains(&(i - 1)) {
                let mut len = 1;

                while set.contains(&(i + len)) {
                    len += 1;
                }

                if len > longest {
                    longest = len;
                }
            }
        }

        longest
    }
}
