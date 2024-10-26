package main

func getCellCharacterD(i, j, rows, cols int) string {
	switch {
	case isLeftCorner(i, j, rows):
		return "A"
	case isRightCorner(i, j, rows, cols):
		return "C"
	case isVerticalBorder(i, j, rows, cols):
		return "B"
	case isHorizontalBorder(i, j, rows, cols):
		return "B"
	default:
		return " "
	}
}

func QuadD(cols, rows int) {
	if cols <= 0 || rows <= 0 {
		return
	}

	matrix := createMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = getCellCharacterD(i, j, rows, cols)
		}
	}

	printMatrix(matrix)
}
