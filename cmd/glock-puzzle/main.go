package main

import (
	"flag"
	"fmt"

	"github.com/jtbonhomme/glock-puzzle/internal/board"
)

func main() {
	var boardSize = flag.Int("n", 3, "board size")
	flag.Parse()

	b := board.New(*boardSize)
	fmt.Println(b.String())
	fmt.Println(b.Graph())

	b.Draw(2)
	fmt.Println(b.String())
	fmt.Println(b.Graph())

	b.Draw(4)
	fmt.Println(b.String())
	fmt.Println(b.Graph())
}
