package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var errLog = log.New(os.Stderr, "Error: ", log.Lshortfile)

func scenicScoreNorth(grid [][]int, rowNum int, colNum int) int {
	if rowNum == 0 {
		return 0
	}

	treeHeight := grid[rowNum][colNum]

	for i := rowNum - 1; i >= 0; i-- {
		if grid[i][colNum] >= treeHeight {
			// Return the number of rows from given tree to the
			// first one that blocks the view.
			return rowNum - i
		}
	}

	return rowNum
}

func scenicScoreSouth(grid [][]int, rowNum int, colNum int) int {
	nRows := len(grid)

	if rowNum == nRows-1 {
		return 0
	}

	treeHeight := grid[rowNum][colNum]

	for i := rowNum + 1; i < nRows; i++ {
		if grid[i][colNum] >= treeHeight {
			// Return the number of rows from given tree to the
			// first one that blocks the view.
			return i - rowNum
		}
	}

	return nRows - rowNum - 1
}

func scenicScoreWest(grid [][]int, rowNum int, colNum int) int {
	if colNum == 0 {
		return 0
	}

	treeHeight := grid[rowNum][colNum]

	for i := colNum - 1; i >= 0; i-- {
		if grid[rowNum][i] >= treeHeight {
			// Return the number of rows from given tree to the
			// first one that blocks the view.
			return colNum - i
		}
	}

	return colNum
}

func scenicScoreEast(grid [][]int, rowNum int, colNum int) int {
	nCols := len(grid[0])

	if colNum == nCols-1 {
		return 0
	}

	treeHeight := grid[rowNum][colNum]

	for i := colNum + 1; i < nCols; i++ {
		if grid[rowNum][i] >= treeHeight {
			// Return the number of rows from given tree to the
			// first one that blocks the view.
			return i - colNum
		}
	}

	return nCols - colNum - 1
}

func calcScenicScore(grid [][]int, rowNum int, colNum int) int {
	northScore := scenicScoreNorth(grid, rowNum, colNum)
	southScore := scenicScoreSouth(grid, rowNum, colNum)
	eastScore := scenicScoreEast(grid, rowNum, colNum)
	westScore := scenicScoreWest(grid, rowNum, colNum)

	return northScore * southScore * eastScore * westScore
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
	maxScenicScore := 0

	// Calculate scenic scores for every tree.
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if score := calcScenicScore(grid, i, j); score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}

	fmt.Println(maxScenicScore)
}
