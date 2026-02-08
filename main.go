package main

import (
	"fmt"
	"io"
	"os"

	"github.com/WaitWut862/cli-dungeon-crawler/internal/camera"
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

	cam := camera.NewCamera(80, 40)

	renderStart(w, p, cam)

	for {
		fmt.Scanln(&i)

		switch i {
		case "m", "move", "i", "inspect", "p", "perform":
			readAndRun(i, p, cam)
			w.UpdateTick()
			render(w, p, cam)
		case "h", "help":
			readAndRun(i, p, cam)
		default:
			readAndRun(i, p, cam)
			render(w, p, cam)
		}
	}
}

func readAndRun(i string, p *mob.Player, cam *camera.Camera) {
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

	case "cw":
		cam.MoveNorth()
	case "cs":
		cam.MoveSouth()
	case "ca":
		cam.MoveWest()
	case "cd":
		cam.MoveEast()
	case "cc":
		cam.Position = p.Position
	}
}

func render(w *world.World, p *mob.Player, cam *camera.Camera) {
	f := p.FacingString()
	fmt.Print("\033[H\033[2J")

	halfW := cam.Width / 2
	halfH := cam.Height / 2

	for j := cam.Position.Y + halfH; j >= cam.Position.Y-halfH; j-- {
		for i := cam.Position.X - halfW; i <= cam.Position.X+halfW; i++ {
			pos := common.Position{X: i, Y: j}
			tile, exists := w.TileMap[pos]
			if p.Position == pos {
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
			} else if exists && tile.GroundType != "" {
				print("#")
			} else {
				print(" ")
			}
		}
		println()
	}
	fmt.Printf("position %v | facing %s | health %v | tick %v | camera (%d,%d)\n",
		p.Position, f, p.Health, w.Tick, cam.Position.X, cam.Position.Y)
}

func renderStart(w *world.World, p *mob.Player, cam *camera.Camera) {
	render(w, p, cam)
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
