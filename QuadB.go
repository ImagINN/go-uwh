package main

func getCellCharacterB(i, j, rows, cols int) string {
	switch {
	case isTopLeftCorner(i, j):
		return "/"
	case isTopRightCorner(i, j, cols):
		return "\\"
	case isBottomLeftCorner(i, j, rows):
		return "\\"
	case isBottomRightCorner(i, j, rows, cols):
		return "/"
	case isVerticalBorder(i, j, rows, cols):
		return "*"
	case isHorizontalBorder(i, j, rows, cols):
		return "*"
	default:
		return " "
	}
}

func QuadB(cols, rows int) {
	if cols <= 0 || rows <= 0 {
		return
	}

	matrix := createMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = getCellCharacterB(i, j, rows, cols)
		}
	}

	printMatrix(matrix)
}
