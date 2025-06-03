package main

import (
	"fmt"
	"strings"
)

func RenderGridAuto(dataGrid [][][]string) [][]string {
	size := len(dataGrid)
	maxCellHeight := 0
	maxCellWidth := 0
	for i := range len(dataGrid) {
		for j := range len(dataGrid[i]) {
			cellHeight := len(dataGrid[i][j])
			if cellHeight > maxCellHeight {
				maxCellHeight = cellHeight
			}
			for k := range len(dataGrid[i][j]) {
				cellWidth := len(dataGrid[i][j][k])
				if cellWidth > maxCellWidth {
					maxCellWidth = cellWidth
				}
			}
		}
	}
	grid := renderGrid(size, maxCellHeight, maxCellWidth, dataGrid)
	return grid
}

func renderGrid(count int, cellHeight int, cellWidth int, dataGrid [][][]string) [][]string {
	// Initialize grid
	rows := count*cellHeight + count + 1
	grid := make([][]string, rows)
	for i := range rows {
		grid[i] = make([]string, count)
	}

	// First row
	grid[0][0] = "┌" + strings.Repeat("─", cellWidth) + "┬"
	for i := range cellHeight {
		grid[i+1][0] = "│" + dataGrid[0][0][i] + "│"
	}
	grid[cellHeight+1][0] = "├" + strings.Repeat("─", cellWidth) + "┼"

	for i := 1; i < count-1; i++ {
		grid[0][i] = strings.Repeat("─", cellWidth) + "┬"
		for j := range cellHeight {
			grid[j+1][i] = dataGrid[0][i][j] + "│"
		}
		grid[cellHeight+1][i] = strings.Repeat("─", cellWidth) + "┼"
	}

	grid[0][count-1] = strings.Repeat("─", cellWidth) + "┐\n"
	for i := range cellHeight {
		grid[i+1][count-1] = dataGrid[0][len(dataGrid[0])-1][i] + "│\n"
	}
	grid[cellHeight+1][count-1] = strings.Repeat("─", cellWidth) + "┤\n"

	// Middle rows
	dataGridRowCounter := 1
	for i := cellHeight + 3; i < rows-cellHeight; i += cellHeight + 1 {
		for j := range cellHeight {
			grid[i-1+j][0] = "│" + dataGrid[dataGridRowCounter][0][j] + "│"
		}
		grid[cellHeight-1+i][0] = "├" + strings.Repeat("─", cellWidth) + "┼"

		for j := 1; j < count-1; j++ {
			for k := range cellHeight {
				grid[i-1+k][j] = dataGrid[dataGridRowCounter][j][k] + "│"
			}
			grid[cellHeight-1+i][j] = strings.Repeat("─", cellWidth) + "┼"
		}

		for j := range cellHeight {
			grid[i-1+j][count-1] = dataGrid[dataGridRowCounter][len(dataGrid[dataGridRowCounter])-1][j] + "│\n"
		}
		grid[cellHeight-1+i][count-1] = strings.Repeat("─", cellWidth) + "┤\n"

		dataGridRowCounter++
	}

	// Last row
	for i := range cellHeight {
		grid[rows-i-2][0] = "│" + dataGrid[len(dataGrid)-1][0][cellHeight-1-i] + "│"
	}
	grid[rows-1][0] = "└" + strings.Repeat("─", cellWidth) + "┴"

	for i := 1; i < count-1; i++ {
		for j := range cellHeight {
			grid[rows-j-2][i] = dataGrid[len(dataGrid)-1][i][cellHeight-1-j] + "│"
		}
		grid[rows-1][i] = strings.Repeat("─", cellWidth) + "┴"
	}

	for i := range cellHeight {
		grid[rows-i-2][count-1] = dataGrid[len(dataGrid)-1][len(dataGrid[len(dataGrid)-1])-1][cellHeight-1-i] + "│\n"
	}
	grid[rows-1][count-1] = strings.Repeat("─", cellWidth) + "┘\n"

	return grid
}

func printGrid(grid [][]string) string {
	str := ""
	for i := range len(grid) {
		for j := range len(grid[i]) {
			str += grid[i][j]
		}
	}
	return str
}

func main() {
	dataGrid := [][][]string{
		{{"123", "456"}, {"456", "123"}},
		{{"klm", "fjn"}, {"fjn", "klm"}},
	}
	fmt.Println(printGrid(RenderGridAuto(dataGrid)))
}
