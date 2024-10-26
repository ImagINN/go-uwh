package main

func getCellCharacterE(i, j, rows, cols int) string {
	switch {
	case isTopLeftCorner(i, j):
		return "A"
	case isTopRightCorner(i, j, cols):
		return "C"
	case isBottomLeftCorner(i, j, rows):
		return "C"
	case isBottomRightCorner(i, j, rows, cols):
		return "A"
	case isVerticalBorder(i, j, rows, cols):
		return "B"
	case isHorizontalBorder(i, j, rows, cols):
		return "B"
	default:
		return " "
	}
}

func QuadE(cols, rows int) {
	if cols <= 0 || rows <= 0 {
		return
	}

	matrix := createMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = getCellCharacterE(i, j, rows, cols)
		}
	}

	printMatrix(matrix)
}
