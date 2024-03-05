package main

import "fmt"

/**

Create a struct called 'Matrix'. The Matrix struct has the following information:
number of rows of matrix
number of columns of matrix
elements of matrix (You can use 2D vector)

The Matrix struct has methods for each of the following:
get the number of rows
get the number of columns
set the elements of the matrix at a given position (i,j)
adding two matrices. (2nd matrix can be taken as input to the method)
print matrix structure as json
You can assume that the dimensions are correct for the multiplication and addition.

Expectation: Use structs, slices, methods

*/

type Matrix struct {
	//number of rows of matrix
	rows int
	//number of rows of matrix
	columns int
	//elements of matrix (You can use 2D vector)
	elements [][]int
}

// initialize the matrix
func NewMatrix(rows, columns int) *Matrix {
	m := Matrix{rows, columns, make([][]int, rows)}
	for i := 0; i < rows; i++ {
		m.elements[i] = make([]int, columns)
	}
	return &m
}

// get the number of rows
func (m *Matrix) getRows() int {
	return m.rows
}

// get the number of columns
func (m *Matrix) getColumns() int {
	return m.columns
}

// get the elements of the matrix at a given position (i,j)
func (m *Matrix) getElement(i, j int) int {
	return m.elements[i][j]
}

// set the elements of the matrix at a given position (i,j)
func (m *Matrix) setElement(i, j, k int) {
	m.elements[i][j] = k
}

// adding two matrices. (2nd matrix can be taken as input to the method)
func (m *Matrix) addMatrix(n *Matrix) {
	if m.getRows() != n.getRows() || m.getColumns() != n.getColumns() {
		fmt.Println("Matrices are not of same size")
		return
	}
	for i := 0; i < m.getRows(); i++ {
		for j := 0; j < m.getColumns(); j++ {
			m.setElement(i, j, m.getElement(i, j)+n.getElement(i, j))
		}
	}
}

// print matrix structure as json
func (matrix *Matrix) printMatrixInJson() {
	fmt.Println("{")
	for _, row := range matrix.elements {
		for _, column := range row {
			fmt.Printf("%v ", column)
		}
		fmt.Println("")
	}
	fmt.Println("}")
}

func main() {
	fmt.Println("Running")
	m := NewMatrix(3, 3)
	fmt.Println("Initial m :", *m)
	k := 0
	for i := 0; i < m.getRows(); i++ {
		for j := 0; j < m.getColumns(); j++ {
			k = k + 1
			m.setElement(i, j, k)
		}
	}
	fmt.Println("Initial m :", *m)

	n := NewMatrix(3, 3)
	for i := 0; i < n.getRows(); i++ {
		for j := 0; j < n.getColumns(); j++ {
			k = k + 1
			n.setElement(i, j, k)
		}
	}
	fmt.Println("Initial n :", *n)

	m.addMatrix(n)
	fmt.Println("After adding n to m :", *m)

	fmt.Println("result json : ", m.printMatrixInJson)
}
