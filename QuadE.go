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

func isTopLeftCorner(i, j int) bool {
	return i == 0 && j == 0
}

func isTopRightCorner(i, j, cols int) bool {
	return i == 0 && j == cols-1
}

func isBottomLeftCorner(i, j, rows int) bool {
	return i == rows-1 && j == 0
}

func isBottomRightCorner(i, j, rows, cols int) bool {
	return i == rows-1 && j == cols-1
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

func QuadE(cols, rows int) {
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
	fmt.Println("TEST CASES:")
	fmt.Println("Test Case 1:")
	fmt.Println()
	QuadE(5, 3)
	fmt.Println("------------------------")
	fmt.Println("Test Case 2:")
	fmt.Println()
	QuadE(5, 1)
	fmt.Println("------------------------")
	fmt.Println("Test Case 3:")
	fmt.Println()
	QuadE(1, 1)
	fmt.Println("------------------------")
	fmt.Println("Test Case 4:")
	fmt.Println()
	QuadE(1, 5)
	fmt.Println("------------------------")
	fmt.Println("Test Case 5:")
	fmt.Println()
	QuadE(0, 0)
	fmt.Println("------------------------")
	fmt.Println("Test Case 6:")
	fmt.Println()
	QuadE(18, 8)
}
