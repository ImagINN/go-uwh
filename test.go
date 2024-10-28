package main

import "fmt"

func test(x, y int, function func(int, int)) {
	function(x, y)
}

func testQuad(quadName string, quad func(int, int)) {
	fmt.Println("TEST CASES: " + quadName)
	fmt.Println()
	fmt.Println("Test Case 1:")
	fmt.Println()
	test(5, 3, quad)
	fmt.Println("------------------------")
	fmt.Println("Test Case 2:")
	fmt.Println()
	test(5, 1, quad)
	fmt.Println("------------------------")
	fmt.Println("Test Case 3:")
	fmt.Println()
	test(7, 16, quad)
	fmt.Println("------------------------")
	fmt.Println("Test Case 4:")
	fmt.Println()
	test(1, 5, quad)
	fmt.Println("------------------------")
	fmt.Println("Test Case 5:")
	fmt.Println()
	test(0, 0, quad)
	fmt.Println("------------------------")
	fmt.Println("Test Case 6:")
	fmt.Println()
	test(3, -2, quad)
	fmt.Println("******************************************")
}
