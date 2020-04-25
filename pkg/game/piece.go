package game

import "go/types"

type Piece struct {
	Player Player
	Origin types.Tuple
	Width  int8
	Height int8
}
