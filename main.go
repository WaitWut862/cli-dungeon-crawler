package main

import (
	"fmt"
	"io"
	"os"

	"github.com/WaitWut862/cli-dungeon-crawler/internal/common"
	"github.com/WaitWut862/cli-dungeon-crawler/internal/mob"
	"github.com/WaitWut862/cli-dungeon-crawler/internal/resources"
	"github.com/WaitWut862/cli-dungeon-crawler/internal/world"
)

func main() {
	if err := resources.LoadItems(); err != nil {
		fmt.Println("Failed to load items:", err)
		os.Exit(1)
	}

	if err := resources.LoadMobs(); err != nil {
		fmt.Println("Failed to load mobs:", err)
		os.Exit(1)
	}

	p := mob.InitPlayer()
	var i string

	w := &world.World{}
	w.MakeWorld()

	var ptA common.Position = common.Position{X: -60, Y: -60}
	var ptB common.Position = common.Position{X: 60, Y: 60}
	w.GenerateChunk(ptA, ptB)

	renderStart(w, p)

	for {
		fmt.Scanln(&i)

		switch i {
		case "m", "move", "i", "inspect", "p", "perform":
			readAndRun(i, p)
			w.UpdateTick()
			render(w, p)
		case "h", "help":
			readAndRun(i, p)
		default:
			readAndRun(i, p)
			render(w, p)
		}
	}
}

func readAndRun(i string, p *mob.Player) {
	switch i {
	case "h", "help":
		printHelp()

	case "q", "quit":
		os.Exit(0)

	case "m", "move":
		p.Move()

	case "l", "left":
		p.TurnLeft()

	case "r", "right":
		p.TurnRight()
	}
}

func render(w *world.World, p *mob.Player) {
	f := p.FacingString()
	fmt.Print("\033[H\033[2J")
	for j := 60; j >= -60; j-- {
		for i := -60; i <= 60; i++ {
			pos := common.Position{X: i, Y: j}
			tile, exists := w.TileMap[pos]
			if (!exists || tile.GroundType == "") && p.Position != pos {
				print(" ")
			} else if p.Position == pos {
				switch p.Facing {
				case common.North:
					print("^")
				case common.East:
					print(">")
				case common.South:
					print("v")
				case common.West:
					print("<")
				}
			} else {
				print("#")
			}
		}
		println()
	}
	fmt.Printf("position %v | facing %s | health %v | tick %v\n", p.Position, f, p.Health, w.Tick)
}

func renderStart(w *world.World, p *mob.Player) {
	render(w, p)
	fmt.Println()
	fmt.Println("Enter 'help' or 'h' to see a detailed list of all available moves")
}

func printHelp() {
	file, err := os.Open("help.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Print("\033[H\033[2J")
	fmt.Println(string(data))
}
