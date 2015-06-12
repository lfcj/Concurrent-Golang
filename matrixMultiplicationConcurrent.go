//Matrix Multiplication with Go

package main

import (
	"fmt"
	"sync"
)
// Implementation of a Barrier from Book with ISBN: 978-2-642-29968-1/Page 103
type Barrier struct {
	M, n     uint
	mutex, s sync.Mutex
}

func createBarrier(n uint) *Barrier {
	x := new(Barrier)
	x.M = n
	x.s.Lock()
	return x
}
func (x *Barrier) Wait() {
	x.mutex.Lock()
	x.n++
	if x.n == x.M {
		if x.n == 1 {
			x.mutex.Unlock()
		} else {
			x.n--
			x.s.Unlock()
		}
	} else {
		x.mutex.Unlock()
		x.s.Lock()
		x.n--
		if x.n == 0 {
			x.mutex.Unlock()
		} else {
			x.s.Unlock()
		}
	}
}
// End of implementation of Barrier. 

/* Brief 'cheatsheet' of matrix multiplication. 
		mat(m x n) * mat2(n x p) = C(m x p)
			mat = [[x,y,z]
				   [1,2,3]
			       [x,y,z]
			       [3,4,5]]
			mat2 = [[x,y,z,5,a]
				 	[1,2,3,6,b]
			     	[x,y,z,7,c]]
			mat(4x3)
			mat2(3x5)
			mat.length = 4 -> rows of mat
			mat[0].length = 3 -> columns of mat
	*/
var (
	result  [][]int
	barrier *Barrier
)

// Format: Multiplicate A with B. 
func multiplicate(matA [][]int, matB [][]int) {

	// Initialize result
	result = make([][]int, len(matA))
	for i := range result {
		result[i] = make([]int, len(matB[0]))
	}
	// Initialize barrier. +1 because of main thread multiplicate.
	barrier = createBarrier(uint(len(matA) * len(matB[0]))+1)
	// Do multiplications concurrently
	for row := range matA {
		for col := range matB[0] {
			go subMult(row, col, matA[row], matB)
		}
	}
	barrier.Wait()
}
func subMult(currRow int, currCol int, rowA []int, matB [][]int) {

	sum := 0
	for i := range rowA {
		number := matB[i][currCol]
		sum += rowA[i] * number
	}
	result[currRow][currCol] = sum
	barrier.Wait()
}

func printMatrix(matrix [][]int) {
	for i := range matrix {
		fmt.Printf("%d\n ", matrix[i])
	}
}

func main() {
	matA := [][]int{{1, 1}, {0, 1}}
	matB := [][]int{{1, 0}, {1, 1}}
	fmt.Print("Test1\nThe desired result is:\n ")
	printMatrix([][]int{{2, 1},
	 	{1, 1}})
	multiplicate(matA, matB)
	fmt.Print("Result:\n ")
	printMatrix(result)

	matA1 := [][]int{{2, 3, -3}, {-1, 6, -4}, {0, 4, 1}, {-2, 5, 7}}
	matB1 := [][]int{{2, 8, 3, -3}, {9, 6, 10, -4}, {11, 4, -1, 0}}
	fmt.Print("Test2\nThe desired result is:\n ")
	multiplicate(matA1, matB1)
	printMatrix([][]int{{-2, 22, 39, -18},
		{8, 12, 61, -21},
		{47, 28, 39, -16},
		{118, 42, 37, -14}})
	fmt.Print("Result:\n ")
	printMatrix(result)
}
