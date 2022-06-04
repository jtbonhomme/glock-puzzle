package board

import (
	"github.com/jtbonhomme/glock-puzzle/internal/cell"
)

type Board struct {
	size  int
	cells []cell.Cell
}

func New(s int) *Board {
	b := &Board{
		size:  s,
		cells: make([]cell.Cell, s*s),
	}
	return b
}

func (b *Board) Draw(cells []int) {

	for _, c := range cells {
		if c >= b.size*b.size {
			continue
		}
		x := c % b.size
		y := (c - c%b.size) / b.size
		b.cells[x+y*b.size].Fill()
	}
}

func (b *Board) String() string {
	var s string
	for y := 0; y < b.size; y++ {
		for x := 0; x < b.size; x++ {
			switch b.cells[x+y*b.size].IsFree() {
			case true:
				s += "[ ] "
			case false:
				s += "[X] "
			}
		}
		s += "\n"
	}
	return s
}
