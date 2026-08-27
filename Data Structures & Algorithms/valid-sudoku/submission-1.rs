impl Solution {
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        let mut seen_rows = vec![vec![false; 10]; 9];
        let mut seen_cols = vec![vec![false; 10]; 9];
        let mut seen_box = vec![vec![false; 10]; 9];

        for (i, row) in board.iter().enumerate() {
            for (j, ch) in row.iter().enumerate() {
                if *ch == '.' {
                    continue;
                }

                let digit = ch.to_digit(10).unwrap() as usize;

                let box_i = (i / 3) * 3 + j / 3;

                if seen_box[box_i][digit] || seen_rows[i][digit] || seen_cols[j][digit] {
                    return false;
                }

                seen_rows[i][digit] = true;
                seen_cols[j][digit] = true;
                seen_box[box_i][digit] = true;
            }
        }

        true
    }
}
