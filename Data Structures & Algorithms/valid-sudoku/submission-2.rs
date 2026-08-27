impl Solution {
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        let mut seen_rows = vec![HashSet::new(); 9];
        let mut seen_cols = vec![HashSet::new(); 9];
        let mut seen_box = vec![HashSet::new(); 9];

        for (i, row) in board.iter().enumerate() {
            for (j, ch) in row.iter().enumerate() {
                if *ch == '.' {
                    continue;
                }

                let digit = ch.to_digit(10).unwrap() as usize;

                let box_i = (i / 3) * 3 + j / 3;

                if seen_box[box_i].contains(&digit)
                    || seen_rows[i].contains(&digit)
                    || seen_cols[j].contains(&digit)
                {
                    return false;
                }

                seen_rows[i].insert(digit);
                seen_cols[j].insert(digit);
                seen_box[box_i].insert(digit);
            }
        }

        true
    }
}
