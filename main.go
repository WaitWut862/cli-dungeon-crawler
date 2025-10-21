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
	w := World{
		tileMap: make(map[Position]Tile),
		entities: make(map[Position][]Entity)
	}
}


func (w *World) updateTick() {
	w.tick = w.tick++
}
