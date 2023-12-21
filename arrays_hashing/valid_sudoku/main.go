package main

func main() {

}

func isValidSudoku(board [][]byte) bool {

	for i := byte(0); i < 9; i++ {
		mapCol := make(map[string]byte)
		mapRow := make(map[string]byte)
		for j := byte(0); j < 9; j++ {
			xCol := board[i][j]
			_, ok := mapCol[string(xCol)]
			if string(xCol) != "." {
				if !ok {
					mapCol[string(xCol)]++
				} else {
					return false
				}
			}

			xRow := board[j][i]
			_, ok = mapRow[string(xRow)]
			if string(xRow) != "." {
				if !ok {
					mapRow[string(xRow)]++
				} else {
					return false
				}
			}

		}
	}

	mapBox3x3 := make(map[string]byte)
	for i := byte(0); i < 9; i += 3 {
		for j := byte(0); j < 9; j += 3 {

			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					x := string(board[k][l])
					if x != "." {
						_, ok := mapBox3x3[x]
						if !ok {
							mapBox3x3[x]++
						} else {
							return false
						}
					}
				}
			}
			mapBox3x3 = make(map[string]byte)
		}
	}

	return true
}
