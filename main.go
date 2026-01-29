package main

import (
	"fmt"
	"io"
	"os"

	"github.com/WaitWut862/cli-dungeon-crawler/internal/player"
)

type World struct {
	tick     int
	tileMap  map[Position]Tile
	entities map[Position][]Entity
}

type Tile struct {
	position   Position
	groundType string
}

type Entity struct {
	position    Position
	name        string
	obstructing bool
}

func main() {
	p := &Player{}
	p.facing = north
	p.health = 100
	var i string

	w := &World{}
	w.makeWorld()

	renderStart(w, p)

	for {
		fmt.Scanln(&i)

		switch i {
		case "m", "move", "i", "inspect", "p", "perform":
			readAndRun(i, p)
			w.updateTick()
			render(w, p)
		case "h", "help":
			readAndRun(i, p)
		default:
			readAndRun(i, p)
			render(w, p)
		}
	}
}

func render(w *World, p *Player) {
	f := p.facingString()
	fmt.Print("\033[H\033[2J")
	fmt.Printf("position %v | facing %s | health %v | tick %v ", p.position, f, p.health, w.tick)
}

func renderStart(w *World, p *Player) {
	render(w, p)
	fmt.Println()
	fmt.Println("Enter 'help' or 'h' to see a detailed list of all available moves")
}

func (w *World) makeWorld() {
	w.tileMap = make(map[Position]Tile)
	w.entities = make(map[Position][]Entity)
}

func (w *World) updateTick() {
	w.tick = w.tick + 1
}

func readAndRun(i string, p *Player) {
	switch i {
	case "h", "help":
		printHelp()

	case "m", "move":
		p.move()

	case "l", "left":
		p.turnLeft()

	case "r", "right":
		p.turnRight()
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
