impl Solution {
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let (mut l, mut r) = (0, numbers.len() - 1);

        while l < r {
            if numbers[l] + numbers[r] < target {
                l += 1;
            } else if numbers[l] + numbers[r] > target {
                r -= 1;
            } else {
                return vec![l as i32 + 1, r as i32 + 1];
            }
        }

        vec![]
    }
}
