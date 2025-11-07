package main

import(
	"fmt"
	"io"
	"os"

	escapes "github.com/snugfox/ansi-escapes"
)

type Player struct {
	inventory Inventory
	health int
	position Position
	statuses Statuses
	facing Direction
}

type Direction int

const(
	north = Direction(iota)
	east
	south
	west
)

type Statuses struct {
	poisoned bool
	boosted bool
	weakened bool
	enraged bool
	fortified bool
}

type Position struct {
	x int
	y int
}

type Inventory struct {
	
}

type World struct {
	tick int
	tileMap map[Position]Tile
	entities map[Position][]Entity
}

type Tile struct {
	position Position
	groundType string
}


type Entity struct {
	position Position
	name string
	obstructing bool
}

func main() {
	p := &Player{}
	p.facing = north
	p.health = 100
	var i string
	
	w := &World{}
	w.makeWorld()
	
	fmt.Println(w)
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
	fmt.Print(escapes.EraseScreen)
	fmt.Println("position ", p.position, ", facing ", f, ", health", p.health, "tick ", w.tick)
}


func renderStart(w *World, p *Player) {
	f := p.facingString()
	fmt.Print(escapes.EraseScreen)
	fmt.Println("Enter 'help' or 'h' to see a detailed list of all available moves")
	fmt.Println("position ", p.position, ", facing ", f, ", health", p.health, "tick ", w.tick)
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



func (p *Player) turnLeft() {
	switch p.facing {
		case north:
			p.facing = west
		case east:
			p.facing = north
		case south:
			p.facing = east
		case west:
			p.facing = south
	}		
}


func (p *Player) turnRight() {
	switch p.facing {
		case north:
			p.facing = east
		case east:
			p.facing = south
		case south:
			p.facing = west
		case west:
			p.facing = north
	}
}


func (p *Player) move() {
	switch p.facing {
		case north:
			p.position.y ++
		case east:
			p.position.x ++
		case south:
			p.position.y --
		case west:
			p.position.x --
	}
}


func (p *Player) facingString() string {
	f := p.facing
	switch f {
		case north:
			return "North"
		case east:
			return "East"
		case south:
			return "South"
		case west:
			return "West"
		default:
			return "Direction not resolved"
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
	fmt.Print(escapes.EraseScreen)
	fmt.Println(string(data))
}

