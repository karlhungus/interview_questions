package main

func rotate(matrix [][]int) {
	n := len(matrix)
	//Transpose (flip across diagonal)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	//Swap across centre row
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[j][i], matrix[j][n-i-1] = matrix[j][n-i-1], matrix[j][i]
		}
	}
}
