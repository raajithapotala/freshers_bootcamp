package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Rows     int
	Columns  int
	Elements [][]int
}

func getrow(m Matrix) int {
	return m.Rows
}

func getcol(m Matrix) int {
	return m.Columns
}

func setelement(m Matrix, i int, j int, element int) {
	m.Elements[i][j] = element
}

func (matrix Matrix) AddMatrix(addMatrix Matrix) {
	if (len(matrix.Elements) != len(addMatrix.Elements)) || (len(matrix.Elements[0]) != len(addMatrix.Elements[0])) {
		panic("Unsupported matrix addition")
	}
	for i := 0; i < len(matrix.Elements); i++ {
		for j := 0; j < len(matrix.Elements[0]); j++ {
			matrix.Elements[i][j] = matrix.Elements[i][j] + addMatrix.Elements[i][j]
		}
	}
}

func printmatrix(m Matrix) {
	data, mat := json.MarshalIndent(m, "", "  ")
	if mat != nil {
		fmt.Println(mat)
	}
	fmt.Println(string(data))
}
func main() {
	m1 := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	n1 := [][]int{{0, 0, 0}, {1, 1, 1}, {2, 2, 2}}
	m := Matrix{3, 3, m1}
	n := Matrix{2, 3, n1}
	fmt.Println("No of rows = ", getrow(m))
	fmt.Println("No of columns = ", getcol(m))
	setelement(m, 1, 1, 5)
	fmt.Println(m.Elements)
	printmatrix(m)
	m.AddMatrix(n)
	fmt.Println(m.Elements)
	printmatrix(m)
}
