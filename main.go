package main

import (
	"fmt"
	"strings"
)

func renderGrid(count int, dataGrid [][]string, cellHeight int, cellWidth int) {
	// Initialize grid
	rows := count*cellHeight + count + 1
	grid := make([][]string, rows)
	for i := range rows {
		grid[i] = make([]string, count)
	}

	// First row
	grid[0][0] = "┌" + strings.Repeat("─", cellWidth) + "┬"
	for i := range cellHeight {
		grid[i+1][0] = "│" + strings.Repeat(" ", cellWidth) + "│"
	}
	grid[cellHeight+1][0] = "├" + strings.Repeat("─", cellWidth) + "┼"

	for i := 1; i < count-1; i++ {
		grid[0][i] = strings.Repeat("─", cellWidth) + "┬"
		for j := range cellHeight {
			grid[j+1][i] = strings.Repeat(" ", cellWidth) + "│"
		}
		grid[cellHeight+1][i] = strings.Repeat("─", cellWidth) + "┼"
	}

	grid[0][count-1] = strings.Repeat("─", cellWidth) + "┐\n"
	for i := range cellHeight {
		grid[i+1][count-1] = strings.Repeat(" ", cellWidth) + "│\n"
	}
	grid[cellHeight+1][count-1] = strings.Repeat("─", cellWidth) + "┤\n"

	// Middle rows
	// dataGridRowCounter := 0
	for i := cellHeight + 3; i < rows-cellHeight; i += cellHeight + 1 {
		for j := range cellHeight {
			grid[i-1+j][0] = "│" + strings.Repeat(" ", cellWidth) + "│"
		}
		grid[cellHeight-1+i][0] = "├" + strings.Repeat("─", cellWidth) + "┼"

		for j := 1; j < count-1; j++ {
			for k := range cellHeight {
				grid[i-1+k][j] = strings.Repeat(" ", cellWidth) + "│"
			}
			grid[cellHeight-1+i][j] = strings.Repeat("─", cellWidth) + "┼"
		}

		for j := range cellHeight {
			grid[i-1+j][count-1] = strings.Repeat(" ", cellWidth) + "│\n"
		}
		grid[cellHeight-1+i][count-1] = strings.Repeat("─", cellWidth) + "┤\n"

		// dataGridRowCounter++
	}

	// // Last row
	for i := range cellHeight {
		grid[rows-i-2][0] = "│" + strings.Repeat(" ", cellWidth) + "│"
	}
	grid[rows-1][0] = "└" + strings.Repeat("─", cellWidth) + "┴"

	for i := 1; i < count-1; i++ {
		for j := range cellHeight {
			grid[rows-j-2][i] = strings.Repeat(" ", cellWidth) + "│"
		}
		grid[rows-1][i] = strings.Repeat("─", cellWidth) + "┴"
	}

	for i := range cellHeight {
		grid[rows-i-2][count-1] = strings.Repeat(" ", cellWidth) + "│\n"
	}
	grid[rows-1][count-1] = strings.Repeat("─", cellWidth) + "┘\n"

	// Print cells
	for i := range rows {
		for j := range count {
			fmt.Print(grid[i][j])
		}
	}
}

func main() {
	const (
		cellHeight = 3
		cellWidth  = 9
		size       = 5
	)
	dataGrid := make([][]string, size)
	for i := range size {
		dataGrid[i] = make([]string, size)
		for j := range size {
			dataGrid[i][j] = fmt.Sprintf("%d", j+1)
		}
	}

	renderGrid(size, dataGrid, cellHeight, cellWidth)
}
