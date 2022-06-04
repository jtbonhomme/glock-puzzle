package board

import (
	"fmt"

	"github.com/jtbonhomme/glock-puzzle/internal/cell"
)

type Board struct {
	size  int
	cells []*cell.Cell
	edges map[int][]int
}

func New(s int) *Board {
	b := &Board{
		size: s,
	}
	for i := 0; i < b.size*b.size; i++ {
		c := cell.New()
		b.AddCell(c)
		// x != 0
		if i%b.size != 0 {
			b.AddEdge((i - 1), i)
		}
		// y != 0
		if (i-i%b.size)/b.size != 0 {
			b.AddEdge(i-b.size, i)
		}
	}
	return b
}

// AddCell adds a cell to the board
func (b *Board) AddCell(c *cell.Cell) {
	b.cells = append(b.cells, c)
}

// AddEdge adds an edge to the board
func (b *Board) AddEdge(c1, c2 int) {
	if b.edges == nil {
		b.edges = make(map[int][]int)
	}
	b.edges[c1] = append(b.edges[c1], c2)
	b.edges[c2] = append(b.edges[c2], c1)
}

func contains(s []int, i int) int {
	for k, v := range s {
		if v == i {
			return k
		}
	}

	return -1
}

func (b *Board) Draw(c int) {
	if c >= b.size*b.size {
		return
	}
	b.cells[c].Fill()
	// near: get cells linked to cell 'c'
	near := b.edges[c]
	for _, n := range near {
		i := contains(b.edges[n], c)
		// for each near cell, remove c from cell's edges
		if i != -1 {
			a := b.edges[n]
			a[i] = a[len(a)-1] // Copy last element to index i.
			a[len(a)-1] = -1   // Erase last element (write zero value).
			a = a[:len(a)-1]   // Truncate slice.
			b.edges[n] = a
		}
	}
	delete(b.edges, c)
}

func (b *Board) String() string {
	var s string
	for y := 0; y < b.size; y++ {
		for x := 0; x < b.size; x++ {
			s += b.cells[x+y*b.size].String() + " "
		}
		s += "\n"
	}
	return s
}

func (b *Board) Graph() string {
	s := ""
	for i := 0; i < len(b.cells); i++ {
		s += fmt.Sprintf("%d -> ", i)
		near := b.edges[i]
		for j := 0; j < len(near); j++ {
			s += fmt.Sprintf("%d ", near[j])
		}
		s += "\n"
	}
	return s
}
