impl Solution {
    pub fn count_students(mut students: Vec<i32>, mut sandwiches: Vec<i32>) -> i32 {
        let mut res = sandwiches.len() as i32;

        let mut i: usize = 0;
        for sw in sandwiches {
            let mut c = 0;
            while c < students.len() && sw != students[i] {
                i += 1;
                i %= students.len();
                c += 1;
            }

            if sw == students[i] {
                students[i] = -1;
                res -= 1;
            } else {
                break;
            }
        }

        res
    }
}
