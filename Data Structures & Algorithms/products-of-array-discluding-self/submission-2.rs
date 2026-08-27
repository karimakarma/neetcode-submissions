impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut res = vec![1; nums.len()];

        for i in 1..nums.len() {
            res[i] = res[i - 1] * nums[i - 1];
        }

        let mut t = *nums.last().unwrap();
        for i in (0..nums.len() - 1).rev() {
            res[i] *= t;
            t *= nums[i];
        }

        res
    }
}
