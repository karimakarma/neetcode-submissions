impl Solution {
    pub fn daily_temperatures(temperatures: Vec<i32>) -> Vec<i32> {
        let mut stack = Vec::<usize>::new();
        let mut res = vec![0; temperatures.len()];

        for i in 0..temperatures.len() {
            while !stack.is_empty() && temperatures[i] > temperatures[*stack.last().unwrap()] {
                let pop = stack.pop().unwrap();
                res[pop] = (i - pop) as i32;
            }

            stack.push(i);
        }

        res
    }
}
