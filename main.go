package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"zombie/grid"

	"github.com/eiannone/keyboard"
)

func setup(size int, cellHeight int) [][][]string {
	dataGrid := make([][][]string, size)
	for i := range dataGrid {
		dataGrid[i] = make([][]string, size)
		for j := range dataGrid[i] {
			dataGrid[i][j] = make([]string, cellHeight)
			for k := range dataGrid[i][j] {
				randomNumber := rand.Intn(101)
				if randomNumber <= 40 {
					dataGrid[i][j][k] = "Zombie"
				} else {
					dataGrid[i][j][k] = "foo"
				}
			}
		}
	}

	return dataGrid
}

func main() {
	// TODO: make a function that generates the initial dataGrid, maybe use setup
	// This way, I can set a size
	rand.Seed(time.Now().UnixNano())
	dataGrid := setup(2, 1)
	fmt.Println(grid.PrintGrid(grid.RenderGridAuto(dataGrid)))

	err := keyboard.Open()
	if err != nil {
		log.Fatalf("Failed to open keyboard: %v", err)
	}
	defer keyboard.Close()

	keyEvents := make(chan keyboard.KeyEvent)

	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				log.Printf("Keyboard error: %v", err)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			keyEvents <- keyboard.KeyEvent{Rune: char, Key: key}
		}
	}()

	gameTick := time.NewTicker(1000 * time.Millisecond)
	defer gameTick.Stop()

	for {
		select {
		case event := <-keyEvents:
			if event.Key == keyboard.KeyEsc || event.Key == keyboard.KeyCtrlC {
				return
			}

			switch event.Rune {
			case 'w':
				log.Print("w")
			case 'a':
				log.Print("a")
			case 's':
				log.Print("s")
			case 'd':
				log.Print("d")
			}

		case <-gameTick.C:
			log.Print("tick")
		}
		log.Print("render")
		fmt.Println(grid.PrintGrid(grid.RenderGridAuto(dataGrid)))
	}
}
