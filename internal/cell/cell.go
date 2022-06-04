package cell

type Cell struct {
	x, y    int
	full    bool
	visited bool
	n       *Cell
	s       *Cell
	e       *Cell
	w       *Cell
}

func New(x, y int) *Cell {
	return &Cell{
		x: x,
		y: y,
	}
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

func (c *Cell) Link(direction string, cell *Cell) {
	switch direction {
	case "N":
		c.n = cell
	case "S":
		c.s = cell
	case "E":
		c.e = cell
	case "W":
		c.w = cell
	}
}
