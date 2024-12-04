package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
)

func check(e error) {
	if e != nil {
		println("Error")
		panic(e)
	}
}

func invalidArguments() {
	println("Invalid arguments")
	println("Usage: aoc04 <input>")
}

func scanRow(row []rune) int {
	windowSize := 4
	re := regexp.MustCompile(`XMAS`)
	re2 := regexp.MustCompile(`SAMX`)

	counter := 0

	for i := 0; i <= len(row)-windowSize; i++ {
		slice := string(row[i : i+windowSize])
		xmasCount := len(re.FindAll([]byte(slice), -1))
		samxCount := len(re2.FindAll([]byte(slice), -1))

		counter += xmasCount
		counter += samxCount
	}

	return counter
}

// flips the matrix along the main diagonal
func transpose(matrix [][]rune) {
	count := len(matrix)

	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}

func reverseColumns(matrix [][]rune) {
	slices.Reverse(matrix)
}

func rotate(matrix [][]rune) {
	transpose(matrix)

	for i := 0; i < len(matrix); i++ {
		slices.Reverse(matrix[i])
	}
}

func getUpperDiagonals(matrix [][]rune) [][]rune {
	width := len(matrix[0])
	var diagonals [][]rune

	for i := 0; i < width; i++ {
		var line []rune
		for j := 0; j < width; j++ {
			max := int(math.Min(float64(width-1), float64(i+j)))
			line = append(line, matrix[j][max])
			if max == width-1 {
				break
			}
		}

		diagonals = append(diagonals, line)
	}

	return diagonals
}

func getDiagonals(matrix [][]rune) [][]rune {
	upper := getUpperDiagonals(matrix)
	transpose(matrix)
	lower := getUpperDiagonals(matrix)
	return slices.Concat(upper, lower[1:]) // watch out for doubling the main diagonal!
}

func main() {
	var args = os.Args[1:]

	if len(args) < 1 {
		invalidArguments()
		return
	}

	data, err := os.ReadFile(args[0])
	check(err)

	// parse file into a matrix
	var matrix [][]rune
	for _, row := range strings.Split(string(data), "\n") {
		var cells []string
		for _, cell := range strings.Split(row, "") {
			cells = append(cells, cell)
		}
		matrix = append(matrix, []rune(row))
	}

	counter := 0

	if len(args) == 2 && args[1] == "X" {
		// looking for X-MAS

		// [i,j][_,_][i,j+2]
		// [_,_][i+1,j+1][_,_]
		// [i+2,j][_,_][i+2,j+2]

		for i := 0; i < len(matrix)-2; i++ {
			for j := 0; j < len(matrix[i])-2; j++ {
				if matrix[i+1][j+1] == 'A' &&
					((matrix[i][j] == 'M' && matrix[i+2][j+2] == 'S') || (matrix[i][j] == 'S' && matrix[i+2][j+2] == 'M')) &&
					((matrix[i][j+2] == 'M' && matrix[i+2][j] == 'S') || (matrix[i][j+2] == 'S' && matrix[i+2][j] == 'M')) {
					counter += 1
				}
			}
		}
	} else {
		// looking for XMAS
		// scan it rows
		for _, row := range matrix {
			counter += scanRow(row)
		}

		// scan diagonals in one direction
		diagonals := getDiagonals(matrix)
		for _, row := range diagonals {
			counter += scanRow(row)
		}

		// reversing columns and scanning for diagonals in other direction
		reverseColumns(matrix)
		revDiagonals := getDiagonals(matrix)

		for _, row := range revDiagonals {
			counter += scanRow(row)
		}

		// rotating 90 degreese to scan the columns
		rotate(matrix)
		for _, row := range matrix {
			counter += scanRow(row)
		}
	}

	fmt.Println("Found ", counter)
}
