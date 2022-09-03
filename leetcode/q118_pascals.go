package leetcode

import "fmt"

// Given an integer numRows,
// return the first numRows of Pascal's triangle.

func GeneratePascalTriangle(numRows int) [][]int {
	triangle := [][]int{{1}}

	for i := 2; i <= numRows; i++ {
		row := make([]int, i)
		row[0] = 1
		row[len(row)-1] = 1

		parent := triangle[i-2]

		// We only need to loop from second element to the "middle"
		// since each row is a palindrome
		limit := len(parent)/2 + 1
		for j := 1; j < limit; j++ {
			n := parent[j-1] + parent[j]
			row[j] = n
			row[len(row)-1-j] = n
		}

		triangle = append(triangle, row)
	}

	return triangle
}

func PrettyTriangle(numRows int) {
	pt := GeneratePascalTriangle(numRows)

	for _, row := range pt {
		fmt.Println(row)
	}
}
