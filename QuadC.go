package main

func getCellCharacterC(i, j, rows, cols int) string {
	switch {
	case isTopCorner(i, j, cols):
		return "A"
	case isBottomCorner(i, j, rows, cols):
		return "C"
	case isVerticalBorder(i, j, rows, cols):
		return "B"
	case isHorizontalBorder(i, j, rows, cols):
		return "B"
	default:
		return " "
	}
}

func QuadC(cols, rows int) {
	if cols <= 0 || rows <= 0 {
		return
	}

	matrix := createMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = getCellCharacterC(i, j, rows, cols)
		}
	}

	printMatrix(matrix)
}
