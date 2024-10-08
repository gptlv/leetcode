package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	columns := [9][9]bool{}
	subboxes := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '1'
			boxIndex := (i/3)*3 + j/3

			if rows[i][num] || columns[j][num] || subboxes[boxIndex][num] {
				return false
			}

			rows[i][num] = true
			columns[j][num] = true
			subboxes[boxIndex][num] = true
		}
	}

	return true
}
