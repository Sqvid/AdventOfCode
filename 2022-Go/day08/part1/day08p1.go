package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var errLog = log.New(os.Stderr, "Error: ", log.Lshortfile)

func isHiddenFromNorth(grid [][]int, rowNum int, colNum int) bool {
	treeHeight := grid[rowNum][colNum]

	for i := 0; i < rowNum; i++ {
		if grid[i][colNum] >= treeHeight {
			return true
		}
	}

	return false
}

func isHiddenFromSouth(grid [][]int, rowNum int, colNum int) bool {
	nCols := len(grid[0])

	treeHeight := grid[rowNum][colNum]

	for i := rowNum + 1; i < nCols; i++ {
		if grid[i][colNum] >= treeHeight {
			return true
		}
	}

	return false
}

func isHiddenFromWest(grid [][]int, rowNum int, colNum int) bool {
	treeHeight := grid[rowNum][colNum]

	for i := 0; i < colNum; i++ {
		if grid[rowNum][i] >= treeHeight {
			return true
		}
	}

	return false
}

func isHiddenFromEast(grid [][]int, rowNum int, colNum int) bool {
	nCols := len(grid[0])
	treeHeight := grid[rowNum][colNum]

	for i := colNum + 1; i < nCols; i++ {
		if grid[rowNum][i] >= treeHeight {
			return true
		}
	}

	return false

}

func isHidden(grid [][]int, rowNum int, colNum int) bool {
	nRows, nCols := len(grid), len(grid[0])

	if rowNum == 0 || rowNum == nRows-1 || colNum == 0 || colNum == nCols-1 {
		return false
	}

	if isHiddenFromNorth(grid, rowNum, colNum) &&
		isHiddenFromSouth(grid, rowNum, colNum) &&
		isHiddenFromWest(grid, rowNum, colNum) &&
		isHiddenFromEast(grid, rowNum, colNum) {
		return true
	}

	return false
}

func main() {
	input, err := os.Open("../input/input")
	if err != nil {
		errLog.Fatalln(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes)

	grid := make([][]int, 0, 100)
	grid = append(grid, []int{})

	var row, col int

	// Load data into grid.
	for scanner.Scan() {
		ch := scanner.Text()

		if ch == "\n" {
			grid = append(grid, []int{})
			row++
		} else {
			val, err := strconv.Atoi(ch)
			if err != nil {
				errLog.Fatalln(err)
			}

			grid[row] = append(grid[row], val)
			col++
		}
	}
	// Remove trailing empty row.
	grid = grid[:len(grid)-1]

	nRows, nCols := len(grid), len(grid[0])
	numVisible := 0

	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if !isHidden(grid, i, j) {
				numVisible++
			}
		}
	}

	fmt.Println(numVisible)
}
