package common

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Position struct {
	X int
	Y int
}

type Item struct {
	Name        string
	Description string
}

type Inventory struct {
	Length int
	Slots  []Item
}
