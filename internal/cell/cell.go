package cell

type Cell struct {
	full    bool
	visited bool
}

func New() *Cell {
	return &Cell{}
}

func (c *Cell) IsFree() bool {
	return !c.full
}

func (c *Cell) Fill() {
	c.full = true
}

func (c *Cell) Free() {
	c.full = false
}

func (c *Cell) Visit() {
	c.visited = true
}

func (c *Cell) IsVisited() bool {
	return !c.visited
}

func (c *Cell) String() string {
	if c.IsFree() {
		return "[ ]"
	}
	return "[X]"
}
