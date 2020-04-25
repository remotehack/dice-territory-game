package game

import "fmt"

const (
	minBoardHeight = 12
	maxBoardHeight = 30
	minBoardWidth  = 8
	maxBoardWidth  = 24
)

type Board struct {
	Width  int8
	Height int8
	Pieces []Piece
}

func NewBoard(width int8, height int8) (Board, error) {
	if !isBetween(height, minBoardHeight, maxBoardHeight) {
		return Board{}, fmt.Errorf("board height is out of bounds. Got %d, need to be between %d and %d inclusive", height, minBoardHeight, maxBoardHeight)
	}
	if !isBetween(width, minBoardWidth, maxBoardWidth) {
		return Board{}, fmt.Errorf("board width is out of bounds. Got %d, need to be between %d and %d inclusive", width, minBoardWidth, maxBoardWidth)
	}
	return Board{Width: width, Height: height, Pieces: []Piece{}}, nil
}

func isBetween(what int8, lowerBoundary int8, upperBoundary int8) bool {
	return what >= lowerBoundary && what <= upperBoundary
}
