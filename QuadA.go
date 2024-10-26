package main

func getCellCharacterA(i, j, rows, cols int) string {
	switch {
	case isCorner(i, j, rows, cols):
		return "o"
	case isVerticalBorder(i, j, rows, cols):
		return "|"
	case isHorizontalBorder(i, j, rows, cols):
		return "-"
	default:
		return " "
	}
}

func QuadA(cols, rows int) {
	if cols <= 0 || rows <= 0 {
		return
	}

	matrix := createMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = getCellCharacterA(i, j, rows, cols)
		}
	}

	printMatrix(matrix)
}
