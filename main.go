package main

import "fmt"

func renderGrid(count int, dataGrid [][]string, cellSize int) {
	rows := count*cellSize + count + 1
	grid := make([][]string, rows)
	for i := range rows {
		grid[i] = make([]string, count)
	}

	grid[0][0] = "┌───┬"
	for i := 1; i <= cellSize; i++ {
		grid[i][0] = fmt.Sprintf("│ %s │", dataGrid[0][0])
	}
	grid[cellSize+1][0] = "├───┼"

	for i := 1; i < count-1; i++ {
		grid[0][i] = "───┬"
		for j := 1; j <= cellSize; j++ {
			grid[j][i] = fmt.Sprintf(" %s │", dataGrid[0][i])
		}
		grid[cellSize+1][i] = "───┼"
	}

	grid[0][count-1] = "───┐\n"
	for i := 1; i <= cellSize; i++ {
		grid[i][count-1] = fmt.Sprintf(" %s │\n", dataGrid[0][count-1])
	}
	grid[cellSize+1][count-1] = "───┤\n"

	// dataGridRowCounter := 0
	for i := cellSize + 3; i < rows-cellSize; i += cellSize + 1 {
		for j := range cellSize {
			grid[i-1+j][0] = "│   │"
		}
		grid[cellSize-1+i][0] = "├───┼"

		for j := 1; j < count-1; j++ {
			for k := range cellSize {
				grid[i-1+k][j] = "   │"
			}
			grid[cellSize-1+i][j] = "───┼"
		}

		for j := range cellSize {
			grid[i-1+j][count-1] = "   │\n"
		}
		grid[cellSize-1+i][count-1] = "───┤\n"

		// dataGridRowCounter++
	}

	for i := range cellSize {
		grid[rows-i-2][0] = "│   │"
	}
	grid[rows-1][0] = "└───┴"

	for i := 1; i < count-1; i++ {
		for j := range cellSize {
			grid[rows-j-2][i] = "   │"
		}
		grid[rows-1][i] = "───┴"
	}

	for i := range cellSize {
		grid[rows-i-2][count-1] = "   │\n"
	}
	grid[rows-1][count-1] = "───┘\n"

	for i := range rows {
		for j := range count {
			fmt.Print(grid[i][j])
		}
	}
}

func main() {
	size := 5
	dataGrid := make([][]string, size)
	for i := range size {
		dataGrid[i] = make([]string, size)
		for j := range size {
			dataGrid[i][j] = fmt.Sprintf("%d", j+1)
		}
	}
	renderGrid(size, dataGrid, 3)
}
