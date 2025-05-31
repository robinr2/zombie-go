package main

import "fmt"

func renderGrid(count int, dataGrid [][]string) {
	rows := count*2 + 1
	grid := make([][]string, rows)
	for i := range rows {
		grid[i] = make([]string, count)
	}

	grid[0][0] = "┌───┬"
	grid[1][0] = fmt.Sprintf("│ %s │", dataGrid[0][0])
	grid[2][0] = "├───┼"

	for i := 1; i < count-1; i++ {
		grid[0][i] = "───┬"
		grid[1][i] = fmt.Sprintf(" %s │", dataGrid[0][i])
		grid[2][i] = "───┼"
	}

	grid[0][count-1] = "───┐\n"
	grid[1][count-1] = fmt.Sprintf(" %s │\n", dataGrid[0][count-1])
	grid[2][count-1] = "───┤\n"

	dataGridRowCounter := 0
	for i := 0; i < rows-3; i += 2 {
		grid[3+i][0] = fmt.Sprintf("│ %s │", dataGrid[dataGridRowCounter][0])
		grid[4+i][0] = "├───┼"

		for j := 1; j < count-1; j++ {
			grid[3+i][j] = fmt.Sprintf(" %s │", dataGrid[dataGridRowCounter][j])
			grid[4+i][j] = "───┼"
		}

		grid[3+i][count-1] = fmt.Sprintf(" %s │\n", dataGrid[dataGridRowCounter][count-1])
		grid[4+i][count-1] = "───┤\n"

		dataGridRowCounter++
	}

	grid[rows-2][0] = fmt.Sprintf("│ %s │", dataGrid[dataGridRowCounter-1][0])
	grid[rows-1][0] = "└───┴"

	for i := 1; i < count-1; i++ {
		grid[rows-2][i] = fmt.Sprintf(" %s │", dataGrid[dataGridRowCounter-1][i])
		grid[rows-1][i] = "───┴"
	}

	grid[rows-2][count-1] = fmt.Sprintf(" %s │\n", dataGrid[dataGridRowCounter-1][count-1])
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
	renderGrid(size, dataGrid)
}
