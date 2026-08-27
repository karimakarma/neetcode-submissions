impl Solution {
    pub fn trap(height: Vec<i32>) -> i32 {
        let mut res = 0;

        let (mut maxl, mut maxr) = (height.first().unwrap(), height.last().unwrap());

        let (mut l, mut r) = (0, height.len() - 1);

        while l < r {
            if maxl < maxr {
                l += 1;
                if &height[l] > maxl {
                    maxl = &height[l];
                }
                res += maxl - height[l];
                continue;
            }

            r -= 1;
            if &height[r] > maxr {
                maxr = &height[r];
            }
            res += maxr - height[r];
        }

        res

    }
}
