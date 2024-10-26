package main

import (
	"fmt"
	"strings"
)

func createMatrix(rows, cols int) [][]string {
	matrix := make([][]string, rows)
	for i := range matrix {
		matrix[i] = make([]string, cols)
	}
	return matrix
}

func getCellCharacter(i, j, rows, cols int) string {
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

func isCorner(i, j, rows, cols int) bool {
	return (i == 0 && j == 0) || (i == 0 && j == cols-1) ||
		(i == rows-1 && j == 0) || (i == rows-1 && j == cols-1)
}

func isVerticalBorder(i, j, rows, cols int) bool {
	return (i > 0 && i < rows-1) && (j == 0 || j == cols-1)
}

func isHorizontalBorder(i, j, rows, cols int) bool {
	return (j > 0 && j < cols-1) && (i == 0 || i == rows-1)
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(strings.Join(row, ""))
	}
}

func QuadA(cols, rows int) {
	if cols <= 0 || rows <= 0 {
		return
	}

	matrix := createMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = getCellCharacter(i, j, rows, cols)
		}
	}

	printMatrix(matrix)
}

func main() {
	QuadA(5, 3)
	fmt.Println("************************")
	QuadA(5, 1)
	fmt.Println("************************")
	QuadA(1, 1)
	fmt.Println("************************")
	QuadA(1, 5)
	fmt.Println("************************")
	QuadA(0, 0)
	fmt.Println("************************")
	QuadA(3, -2)
}
