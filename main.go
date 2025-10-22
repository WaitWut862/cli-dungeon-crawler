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
	
	var i string
	
	w := World{
		tileMap: make(map[Position]Tile),
		entities: make(map[Position][]Entity),
	}
	fmt.Println(w)

	for {
		fmt.Scanln(&i)

		if i == "m" || i == "l" || i == "r" || i == "i" || i == "p" {
			w.updateTick()
			fmt.Println("Tick was updated: ", w.tick)
			readAndRun(i)
		} else {
			fmt.Println("Tick was not updated")
			readAndRun(i)
		}
	}
}


func (w *World) updateTick() {
	w.tick = w.tick + 1
}


func readAndRun(i string) {
	switch {
		case i == "s":
			fmt.Println("soup")
		case i == "o":
			fmt.Println("owala koala")
		case i == "m":
			fmt.Println("the hero has moved")
	}
}
