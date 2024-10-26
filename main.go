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

func isCorner(i, j, rows, cols int) bool {
	return (i == 0 && j == 0) || (i == 0 && j == cols-1) ||
		(i == rows-1 && j == 0) || (i == rows-1 && j == cols-1)
}

func isTopCorner(i, j, cols int) bool {
	return (i == 0 && j == 0) || (i == 0 && j == cols-1)
}

func isBottomCorner(i, j, rows, cols int) bool {
	return (i == rows-1 && j == 0) || (i == rows-1 && j == cols-1)
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

func main() {
	testQuad("QuadA", QuadA)
	testQuad("QuadB", QuadB)
	testQuad("QuadC", QuadC)
	testQuad("QuadD", QuadD)
	testQuad("QuadE", QuadE)
}
