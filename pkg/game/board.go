package game

import "fmt"

const (
	minBoardHeight = 12
	maxBoardHeight = 30
	minBoardWidth  = 8
	maxBoardWidth  = 24
)

type Board struct {
	Width  uint8
	Height uint8
	Pieces []Piece
}

// Returns a new instance of a board with given height and width.
func NewBoard(width uint8, height uint8) (Board, error) {
	if !isBetween(height, minBoardHeight, maxBoardHeight) {
		return Board{}, fmt.Errorf("board height is out of bounds. Got %d, need to be between %d and %d inclusive", height, minBoardHeight, maxBoardHeight)
	}
	if !isBetween(width, minBoardWidth, maxBoardWidth) {
		return Board{}, fmt.Errorf("board width is out of bounds. Got %d, need to be between %d and %d inclusive", width, minBoardWidth, maxBoardWidth)
	}
	return Board{Width: width, Height: height, Pieces: []Piece{}}, nil
}

// Utility function to check whether the board is the right size.
func isBetween(what uint8, lowerBoundary uint8, upperBoundary uint8) bool {
	return what >= lowerBoundary && what <= upperBoundary
}
