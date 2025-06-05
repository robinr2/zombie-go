package grid

import (
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
		dataGridPosition := dataGrid[0][0][i]
		grid[i+1][0] = "│" + dataGridPosition + strings.Repeat(" ", cellWidth-len(dataGridPosition)) + "│"
	}
	grid[cellHeight+1][0] = "├" + strings.Repeat("─", cellWidth) + "┼"

	for i := 1; i < count-1; i++ {
		grid[0][i] = strings.Repeat("─", cellWidth) + "┬"
		for j := range cellHeight {
			dataGridPosition := dataGrid[0][i][j]
			grid[j+1][i] = dataGridPosition + strings.Repeat(" ", cellWidth-len(dataGridPosition)) + "│"
		}
		grid[cellHeight+1][i] = strings.Repeat("─", cellWidth) + "┼"
	}

	grid[0][count-1] = strings.Repeat("─", cellWidth) + "┐\n"
	for i := range cellHeight {
		dataGridPosition := dataGrid[0][len(dataGrid[0])-1][i]
		grid[i+1][count-1] = dataGridPosition + strings.Repeat(" ", cellWidth-len(dataGridPosition)) + "│\n"
	}
	grid[cellHeight+1][count-1] = strings.Repeat("─", cellWidth) + "┤\n"

	// Middle rows
	dataGridRowCounter := 1
	for i := cellHeight + 3; i < rows-cellHeight; i += cellHeight + 1 {
		for j := range cellHeight {
			dataGridPosition := dataGrid[dataGridRowCounter][0][j]
			grid[i-1+j][0] = "│" + dataGridPosition + strings.Repeat(" ", cellWidth-len(dataGridPosition)) + "│"
		}
		grid[cellHeight-1+i][0] = "├" + strings.Repeat("─", cellWidth) + "┼"

		for j := 1; j < count-1; j++ {
			for k := range cellHeight {
				dataGridPosition := dataGrid[dataGridRowCounter][j][k]
				grid[i-1+k][j] = dataGridPosition + strings.Repeat(" ", cellWidth-len(dataGridPosition)) + "│"
			}
			grid[cellHeight-1+i][j] = strings.Repeat("─", cellWidth) + "┼"
		}

		for j := range cellHeight {
			dataGridPosition := dataGrid[dataGridRowCounter][len(dataGrid[dataGridRowCounter])-1][j]
			grid[i-1+j][count-1] = dataGridPosition + strings.Repeat(" ", cellWidth-len(dataGridPosition)) + "│\n"
		}
		grid[cellHeight-1+i][count-1] = strings.Repeat("─", cellWidth) + "┤\n"

		dataGridRowCounter++
	}

	// Last row
	for i := range cellHeight {
		dataGridPosition := dataGrid[len(dataGrid)-1][0][cellHeight-1-i]
		grid[rows-i-2][0] = "│" + dataGridPosition + strings.Repeat(" ", cellWidth-len(dataGridPosition)) + "│"
	}
	grid[rows-1][0] = "└" + strings.Repeat("─", cellWidth) + "┴"

	for i := 1; i < count-1; i++ {
		for j := range cellHeight {
			dataGridPosition := dataGrid[len(dataGrid)-1][i][cellHeight-1-j]
			grid[rows-j-2][i] = dataGridPosition + strings.Repeat(" ", cellWidth-len(dataGridPosition)) + "│"
		}
		grid[rows-1][i] = strings.Repeat("─", cellWidth) + "┴"
	}

	for i := range cellHeight {
		dataGridPosition := dataGrid[len(dataGrid)-1][len(dataGrid[len(dataGrid)-1])-1][cellHeight-1-i]
		grid[rows-i-2][count-1] = dataGridPosition + strings.Repeat(" ", cellWidth-len(dataGridPosition)) + "│\n"
	}
	grid[rows-1][count-1] = strings.Repeat("─", cellWidth) + "┘\n"

	return grid
}

func PrintGrid(grid [][]string) string {
	str := ""
	for i := range len(grid) {
		for j := range len(grid[i]) {
			str += grid[i][j]
		}
	}
	return str
}
