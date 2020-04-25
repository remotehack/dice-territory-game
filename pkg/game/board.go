package game

import (
	"errors"
	"fmt"
)

const (
	minBoardHeight = 12
	maxBoardHeight = 30
	minBoardWidth  = 8
	maxBoardWidth  = 24
)

type Board struct {
	Width   uint8
	Height  uint8
	Corners []Coordinate
	Pieces  []Piece
}

// Returns a new instance of a board with given height and width.
func NewBoard(width uint8, height uint8) (Board, error) {
	if !isBetween(height, minBoardHeight, maxBoardHeight) {
		return Board{}, fmt.Errorf("board height is out of bounds. Got %d, need to be between %d and %d inclusive", height, minBoardHeight, maxBoardHeight)
	}
	if !isBetween(width, minBoardWidth, maxBoardWidth) {
		return Board{}, fmt.Errorf("board width is out of bounds. Got %d, need to be between %d and %d inclusive", width, minBoardWidth, maxBoardWidth)
	}

	var c []Coordinate

	// Top left corner (start for player 1)
	c = append(c, Coordinate{X: 1, Y: 1})

	// Bottom right corner (start for player 2)
	c = append(c, Coordinate{X: width, Y: height})

	return Board{Width: width, Height: height, Corners: c, Pieces: []Piece{}}, nil
}

// Utility function to check whether the board is the right size.
func isBetween(what uint8, lowerBoundary uint8, upperBoundary uint8) bool {
	return what >= lowerBoundary && what <= upperBoundary
}

func (b *Board) canPlacePiece(p Piece) bool {
	// If any of the coordinates of the new piece intersect ANY pieces, it can't be placed.
	for _, piece := range b.Pieces {
		for _, c := range p.Coordinates {
			if piece.IsCoordinateWithin(c) {
				return false
			}
		}
	}

	// If any of the coordinates of the new piece is adjacent to our OWN pieces, it _can_ be placed.
	for _, piece := range b.Pieces {
		if p.Player != piece.Player {
			continue
		}
		for _, c := range p.Coordinates {
			if piece.IsAdjacent(c) {
				return true
			}
		}
	}

	return false
}

func (b *Board) PlacePiece(p Piece) (*Board, error) {
	if !b.canPlacePiece(p) {
		return b, errors.New("can't place piece there")
	}

	b.Pieces = append(b.Pieces, p)
	return b, nil
}
