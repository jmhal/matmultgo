package matmultgo

import "time"

// Creates a row major matrix
func createMatrix(order int) []float64 {
	var matrix []float64
	for i := 0; i < order*order; i++ {
		matrix = append(matrix, 1)
	}
	return matrix
}

// Receives two row major matrices and multiply them
// Returns the product matrix and the elapsed time
func matMultByInput(order int, matrix1 []float64, matrix2 []float64) ([]float64, time.Duration) {
	matrix3 := createMatrix(order)

	start := time.Now()
	var temp float64
	for i := 0; i < order; i++ {
		for j := 0; j < order; j++ {
			temp = 0
			for k := 0; k < order; k++ {
				temp += matrix1[i*order+k] * matrix2[k*order+j]
			}
			matrix3 = append(matrix3, temp)
		}
	}
	elapsed := time.Since(start)

	return matrix3, elapsed
}

// For bechmark purposes, multiply matrices filled with ones
func matMultByOrder(order int) ([]float64, time.Duration) {
	matrix1 := createMatrix(order)
	matrix2 := createMatrix(order)
	matrix3 := createMatrix(order)

	start := time.Now()
	var temp float64
	for i := 0; i < order; i++ {
		for j := 0; j < order; j++ {
			temp = 0
			for k := 0; k < order; k++ {
				temp += matrix1[i*order+k] * matrix2[k*order+j]
			}
			matrix3 = append(matrix3, temp)
		}
	}
	elapsed := time.Since(start)

	return matrix3, elapsed
}

