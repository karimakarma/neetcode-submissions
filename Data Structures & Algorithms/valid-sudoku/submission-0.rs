
impl Solution {
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        // Each row must contain the digits 1-9 without duplicates.
        for i in 0..9 {
            let mut seen = [false; 10];
            for j in 0..9 {
                let cell = &board[i][j];

                if *cell == '.' {
                    continue;
                }

                let cell = cell.to_digit(10).unwrap() as usize;

                let mut c = &mut seen[cell];

                if *c {
                    return false;
                }

                *c = true
            }
        }

        // Each column must contain the digits 1-9 without duplicates.
        for i in 0..9 {
            let mut seen = [false; 10];
            for j in 0..9 {
                let cell = &board[j][i];

                if *cell == '.' {
                    continue;
                }

                let cell = cell.to_digit(10).unwrap() as usize;

                let mut c = &mut seen[cell];

                if *c {
                    return false;
                }

                *c = true
            }
        }

        // Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.
        let mut seen = vec![vec![false; 10]; 9];
        for i in 0..9 {
            for j in 0..9 {
                let cell = &board[i][j];

                if *cell == '.' {
                    continue;
                }

                let cell = cell.to_digit(10).unwrap() as usize;

                let idx: usize = (i / 3) * 3 + (j / 3);

                if seen[idx][cell] {
                    return false;
                }

                seen[idx][cell] = true;
            }
        }

        true
    }
}
