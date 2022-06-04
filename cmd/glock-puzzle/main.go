package main

import (
	"flag"
	"fmt"

	"github.com/jtbonhomme/glock-puzzle/internal/board"
)

func main() {
	var boardSize = flag.Int("n", 9, "board size")
	flag.Parse()

	b := board.New(*boardSize)
	b.Draw([]int{2, 3, 4, 5})
	fmt.Println(b.String())
}
