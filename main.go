package main

import (
	"fmt"
	"io"
	"os"

	c "cli-dungeon-crawler/internal/common"
	m "cli-dungeon-crawler/internal/mobs"
	"cli-dungeon-crawler/internal/world"
)

type direction = c.Direction
type position = c.Position

func main() {
	p := new(m.Mob)
	p.Facing = c.North
	p.Health = 100
	var i string

	w := new(world.World)
	w.MakeWorld()

	fmt.Println(w)
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

func render(w *world.World, p *m.Mob) {
	f := p.FacingString()
	fmt.Print("\033[H\033[2J")
	fmt.Println("position ", p.Position, ", facing ", f, ", health", p.Health, "tick ", w.Tick)
}

func renderStart(w *world.World, p *m.Mob) {
	f := p.FacingString()
	fmt.Print("\033[H\033[2J")
	fmt.Println("position ", p.Position, ", facing ", f, ", health", p.Health, "tick ", w.Tick)
	fmt.Println("Enter 'help' or 'h' to see a detailed list of all available moves")
}

func readAndRun(i string, p *m.Mob) {
	switch i {
	case "h", "help":
		printHelp()

	case "m", "move":
		p.Move()

	case "l", "left":
		p.TurnLeft()

	case "r", "right":
		p.TurnRight()
	}
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
