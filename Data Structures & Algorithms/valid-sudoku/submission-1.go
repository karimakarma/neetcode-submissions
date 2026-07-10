func isValidSudoku(board [][]byte) bool {
	var rows, columns, grids [9][9]bool

	for y := range 9 {
		for x := range 9 {
			if board[y][x] == '.' {
				continue
			}
			val := board[y][x] - '1'
			gridIndex := (y/3)*3 + x/3
			if rows[y][val] || columns[x][val] || grids[gridIndex][val] {
				return false
			}
			rows[y][val] = true
			columns[x][val] = true
			grids[gridIndex][val] = true
		}
	}

	return true
}
