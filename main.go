package main

import (
	"fmt"
	"io"
	"os"

	"github.com/WaitWut862/cli-dungeon-crawler/internal/player"
	"github.com/WaitWut862/cli-dungeon-crawler/internal/world"
)

func main() {
	p := &player.Player{}
	p.Facing = player.North
	p.Health = 100
	var i string

	w := &world.World{}
	w.MakeWorld()

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

func render(w *world.World, p *player.Player) {
	f := p.FacingString()
	fmt.Print("\033[H\033[2J")
	fmt.Printf("position %v | facing %s | health %v | tick %v ", p.Position, f, p.Health, w.Tick)
}

func renderStart(w *world.World, p *player.Player) {
	render(w, p)
	fmt.Println()
	fmt.Println("Enter 'help' or 'h' to see a detailed list of all available moves")
}


func readAndRun(i string, p *player.Player) {
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
