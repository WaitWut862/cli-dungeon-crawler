package main

import(
	"fmt"
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
	var i string
	
	w := World{
		tileMap: make(map[Position]Tile),
		entities: make(map[Position][]Entity),
	}
	fmt.Println(w)

	for {
		fmt.Scanln(&i)

		switch i {
			case "m", "l", "r", "i", "p":
				readAndRun(i, p)
				w.updateTick()
				fmt.Println("Tick was updated: ", w.tick)
			default:
				readAndRun(i, p)
				fmt.Println("Tick was not updated:", w.tick)
		}
	}
}


func (w *World) updateTick() {
	w.tick = w.tick + 1
}


func readAndRun(i string, p *Player) {
	switch i {
		case "m":
			p.move()
			fmt.Println("the hero has moved", p.position)
			
		case "l":
			p.turnLeft()
			fmt.Println(p.facing)

		case "r":
			p.turnRight()
			fmt.Println(p.facing)
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
