impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, mut m: i32, nums2: &mut Vec<i32>, mut n: i32) {
        m -= 1;
        n -= 1;
        let mut k = m + n + 1;

        while n >= 0 {
            if m >= 0 && nums1[m as usize] > nums2[n as usize] {
                nums1[k as usize] = nums1[m as usize];
                m -= 1;
            } else {
                nums1[k as usize] = nums2[n as usize];
                n -= 1;
            }

            k -= 1;
        }
    }
}
