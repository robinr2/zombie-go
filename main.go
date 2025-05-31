package main

import "fmt"

// func renderGrid(size int) {
// 	fmt.Printf("┌───┬")
// 	fmt.Printf("───┬")
// 	fmt.Printf("───┬")
// 	fmt.Printf("───┬")
// 	fmt.Printf("───┐")
// 	fmt.Printf("\n")
// 	fmt.Printf("│   │")
// 	fmt.Printf("   │")
// 	fmt.Printf("   │")
// 	fmt.Printf("   │")
// 	fmt.Printf("   │")
// 	fmt.Printf("\n")
// 	fmt.Printf("├───┼")
// 	fmt.Printf("───┼")
// 	fmt.Printf("───┼")
// 	fmt.Printf("───┼")
// 	fmt.Printf("───┤")
// 	fmt.Printf("\n")
//
// 	for _ = range size {
// 		fmt.Printf("│   │")
// 		fmt.Printf("   │")
// 		fmt.Printf("   │")
// 		fmt.Printf("   │")
// 		fmt.Printf("   │")
// 		fmt.Printf("\n")
// 		fmt.Printf("├───┼")
// 		fmt.Printf("───┼")
// 		fmt.Printf("───┼")
// 		fmt.Printf("───┼")
// 		fmt.Printf("───┤")
// 		fmt.Printf("\n")
// 	}
//
// 	fmt.Printf("│   │")
// 	fmt.Printf("   │")
// 	fmt.Printf("   │")
// 	fmt.Printf("   │")
// 	fmt.Printf("   │")
// 	fmt.Printf("\n")
// 	fmt.Printf("└───┴")
// 	fmt.Printf("───┴")
// 	fmt.Printf("───┴")
// 	fmt.Printf("───┴")
// 	fmt.Printf("───┘")
// 	fmt.Printf("\n")
//
// }

func renderGrid(count int) {
	rows := count*2 + 1
	grid := make([][]string, rows)
	for i := range rows {
		grid[i] = make([]string, count)
		for j := range count {
			grid[i][j] = "x"
		}
	}

	grid[0][0] = "┌───┬"
	grid[1][0] = "│   │"
	grid[2][0] = "├───┼"

	for i := 1; i < count-1; i++ {
		grid[0][i] = "───┬"
		grid[1][i] = "   │"
		grid[2][i] = "───┼"
	}

	grid[0][count-1] = "───┐\n"
	grid[1][count-1] = "   │\n"
	grid[2][count-1] = "───┤\n"

	for i := 0; i < rows-3; i += 2 {
		grid[3+i][0] = "│   │"
		grid[4+i][0] = "├───┼"

		for j := 1; j < count-1; j++ {
			grid[3+i][j] = "   │"
			grid[4+i][j] = "───┼"
		}

		grid[3+i][count-1] = "   │\n"
		grid[4+i][count-1] = "───┤\n"
	}

	grid[rows-2][0] = "│   │"
	grid[rows-1][0] = "└───┴"

	for i := 1; i < count-1; i++ {
		grid[rows-2][i] = "   │"
		grid[rows-1][i] = "───┴"
	}

	grid[rows-2][count-1] = "   │\n"
	grid[rows-1][count-1] = "───┘\n"

	foo := 0
	for i := range rows {
		for j := range count {
			fmt.Print(grid[i][j])
			if grid[i][j] == "x" {
				foo++
			}
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(foo)
}

func main() {
	renderGrid(25)
}
