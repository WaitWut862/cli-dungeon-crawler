package main

import (
	"fmt"
	"io"
	"os"

	c "cli-dungeon-crawler/internal/common"
	m "cli-dungeon-crawler/internal/mobs"
	"cli-dungeon-crawler/internal/world"
)

func main() {
	p := m.NewMob("player")
	var i string

	w := new(world.World)
	w.MakeWorld()
	w.GenerateChunk(c.Pos(-20, -20), c.Pos(20, 20), "stone")

	renderStart()

	for {
		fmt.Scanln(&i)

		switch i {
		case "m", "move", "i", "inspect", "p", "perform":
			readAndRun(i, p, w)
			w.UpdateTick()
			render(w, p)
		case "h", "help", "d", "debug":
			readAndRun(i, p, w)
		default:
			readAndRun(i, p, w)
			render(w, p)
		}
	}
}

func render(w *world.World, p *m.Mob) {
	f := p.FacingString()

	groundType := ""

	if tile, ok := w.TileMap[p.Position]; ok {
		groundType = tile.GroundType
	}

	lookingAt := "nothing"
	if entities, ok := w.Entities[p.Front()]; ok && len(entities) > 0 {
		lookingAt = entities[0].Name
	}

	fmt.Print("\033[H\033[2J")
	fmt.Printf("positition = %v | facing = %s | floor = %s | looking at =  %s | health = %d | tick = %d", p.Position, f, groundType, lookingAt, p.Health, w.Tick)
}

func renderStart() {
	fmt.Println("Enter 'help' or 'h' to see a detailed list of all available moves")
}

func readAndRun(i string, p *m.Mob, w *world.World) {
	switch i {
	case "h", "help":
		printHelp()

	case "m", "move":
		p.Move()

	case "l", "left":
		p.TurnLeft()

	case "r", "right":
		p.TurnRight()
	case "q", "quit":
		fmt.Println("Terminated")
		os.Exit(0)
	case "d", "debug":
		fmt.Printf("%+v", w.TileMap)
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
