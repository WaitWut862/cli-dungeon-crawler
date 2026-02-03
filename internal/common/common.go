package common

type Position struct {
	X int
	Y int
}

type Direction int

const (
	North = Direction(iota)
	East
	South
	West
)

type Item struct {
	Name        string
	Description string
	Durability  int
}

type Inventory struct {
	Slots    []*Item
	MaxSlots int
}

func AddItem(inv *Inventory, item *Item) bool {
	for i := 0; i < inv.MaxSlots; i++ {
		if inv.Slots[i] == nil {
			inv.Slots[i] = item
			return true
		}
	}
	return false
}

func RemoveItem(inv *Inventory, index int) *Item {
	if index < 0 || index > inv.MaxSlots {
		return nil
	}

	item := inv.Slots[index]
	inv.Slots[index] = nil
	return item
}

func MakeInventory(size int) *Inventory {
	return &Inventory{
		Slots:    make([]*Item, size),
		MaxSlots: size,
	}
}

type Statuses struct {
	Poisoned  bool
	Boosted   bool
	Weakened  bool
	Enraged   bool
	Fortified bool
}
