impl Solution {
    pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {        
        let mut nums = nums;
        nums.sort();
        let mut res = vec![];

        for i in 0..nums.len() {
            if i > 0 && nums[i] == nums[i - 1] {
                continue;
            }

            if nums[i] > 0 {
                break;
            }

            let target = -nums[i];

            let (mut l, mut r) = (i + 1, nums.len() - 1);
            while l < r {
                if l == i {
                    l += 1;
                    continue;
                }

                if r == i {
                    r -= 1;
                    continue;
                }

                let sum = nums[l] + nums[r];

                if sum > target {
                    r -= 1;
                    continue;
                }

                if sum < target {
                    l += 1;
                    continue;
                }

                if sum == target {
                    res.push(vec![nums[i], nums[l], nums[r]]);

                    while l < r && nums[l] == nums[l + 1] {
                        l += 1;
                    }

                    while l < r && nums[r] == nums[r - 1] {
                        r -= 1;
                    }

                    l += 1;
                    r -= 1;
                }
            }
        }

        res
    }
}
