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

func isLeftCorner(i, j, rows int) bool {
	return (i == 0 && j == 0) || (i == rows-1 && j == 0)
}

func isRightCorner(i, j, rows, cols int) bool {
	return (i == 0 && j == cols-1) || (i == rows-1 && j == cols-1)
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

func QuadD(cols, rows int) {
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
	QuadD(5, 3)
	fmt.Println("------------------------")
	fmt.Println("Test Case 2:")
	fmt.Println()
	QuadD(5, 1)
	fmt.Println("------------------------")
	fmt.Println("Test Case 3:")
	fmt.Println()
	QuadD(1, 1)
	fmt.Println("------------------------")
	fmt.Println("Test Case 4:")
	fmt.Println()
	QuadD(1, 5)
	fmt.Println("------------------------")
	fmt.Println("Test Case 5:")
	fmt.Println()
	QuadD(0, 0)
	fmt.Println("------------------------")
	fmt.Println("Test Case 6:")
	fmt.Println()
	QuadD(3, -2)
}
