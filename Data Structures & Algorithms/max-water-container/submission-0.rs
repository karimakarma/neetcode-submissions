impl Solution {
    pub fn max_area(heights: Vec<i32>) -> i32 {
        let mut max = 0;
        let (mut l, mut r) = (0, heights.len() - 1);

        while l < r {
            let width = (r - l) as i32;
            let shortest = if heights[l] < heights[r] { l } else { r };

            let area = width * heights[shortest];

            if area > max {
                max = area;
            }

            if shortest == l {
                l += 1;
            } else {
                r -= 1;
            }
        }

        max
    }
}
