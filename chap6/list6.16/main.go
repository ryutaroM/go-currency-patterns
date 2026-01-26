package main

import (
	"fmt"
	"math/rand"

	listing6_10 "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter6/listing6.10"
)

const matrixSize = 4

func rowMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int, row int, barrier *listing6_10.Barrier) {
	for {
		barrier.Wait()
		for col := 0; col < matrixSize; col++ {
			sum := 0
			for i := 0; i < matrixSize; i++ {
				sum += matrixA[row][i] * matrixB[i][col]
			}
			result[row][col] = sum
		}
		barrier.Wait()
	}
}

func generateMRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

func main() {
	var matrixA, matrixB, result [matrixSize][matrixSize]int
	barrier := listing6_10.NewBarrier(matrixSize + 1)
	for row := 0; row < matrixSize; row++ {
		go rowMultiply(&matrixA, &matrixB, &result, row, barrier)
	}

	for i := 0; i < 4; i++ {
		generateMRandMatrix(&matrixA)
		generateMRandMatrix(&matrixB)
		barrier.Wait()
		barrier.Wait()
		for i := 0; i < matrixSize; i++ {
			fmt.Println(matrixA[i], matrixB[i], result[i])
		}
		fmt.Println()
	}
}
