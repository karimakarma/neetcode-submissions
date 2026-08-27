impl Solution {
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        let mut stack = Vec::<i32>::new();

        for t in tokens.iter() {
            match t.as_str() {
                "+" | "-" | "*" | "/" => {
                    let (b, a) = (stack.pop().unwrap(), stack.pop().unwrap());

                    match t.as_str() {
                        "+" => stack.push(a.saturating_add(b)),
                        "-" => stack.push(a.saturating_sub(b)),
                        "*" => stack.push(a.saturating_mul(b)),
                        "/" => stack.push(a.saturating_div(b)),
                        _ => (),
                    }
                }
                _ => stack.push(t.parse().unwrap()),
            }
        }

        stack.pop().unwrap()
    }
}
