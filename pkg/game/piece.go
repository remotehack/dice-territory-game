package game

import (
	"fmt"
)

type Piece struct {
	Player         Player
	Origin         Coordinate
	AdjacentFields []Coordinate
	Corners        []Coordinate
	Width          uint8
	Height         uint8
}

func NewPiece(player Player, originX uint8, originY uint8, width uint8, height uint8) (Piece, error) {
	if width <= 0 || width > 6 || height <= 0 || height > 6 {
		return Piece{}, fmt.Errorf("can't create piece with these dimensions: width: %d, height: %d", width, height)
	}

	p := Piece{
		Player: player,
		Origin: Coordinate{
			X: originX,
			Y: originY,
		},
		Width:  width,
		Height: height,
	}

	p.AdjacentFields = p.getAdjacentFields()
	p.Corners = p.getCorners()

	return p, nil
}

type Coordinate struct {
	X uint8
	Y uint8
}

func (p Piece) getCorners() []Coordinate {
	var c []Coordinate

	// 0 -> X
	// |
	// v
	// Y

	// Origin coordinate is always the top left corner
	c = append(c, p.Origin)

	// Top right
	c = append(c, Coordinate{X: p.Origin.X + p.Width - 1, Y: p.Origin.Y})

	// Bottom right
	c = append(c, Coordinate{X: p.Origin.X + p.Width - 1, Y: p.Origin.Y + p.Height - 1})

	// Bottom left
	c = append(c, Coordinate{X: p.Origin.X, Y: p.Origin.Y + p.Height - 1})

	return c
}

func (p Piece) IsCoordinateWithin(c Coordinate) bool {
	// check for X
	if c.X >= p.Origin.X && c.X <= p.Origin.X+p.Width-1 && c.Y >= p.Origin.Y && c.Y <= p.Origin.Y+p.Height-1 {
		return true
	}

	return false
}

func (p Piece) IsAdjacent(c Coordinate) bool {
	for _, val := range p.AdjacentFields {
		if c == val {
			return true
		}
	}
	return false
}

func (p Piece) getAdjacentFields() []Coordinate {
	var a []Coordinate

	// horizontals
	for i := uint8(0); i < p.Width; i++ {
		a = append(a, Coordinate{
			X: p.Origin.X + i,
			Y: p.Origin.Y - 1,
		}, Coordinate{
			X: p.Origin.X + i,
			Y: p.Origin.Y + p.Height,
		})
	}

	// verticals
	for i := uint8(0); i < p.Height; i++ {
		a = append(a, Coordinate{
			X: p.Origin.X - 1,
			Y: p.Origin.Y + i,
		}, Coordinate{
			X: p.Origin.X + p.Width,
			Y: p.Origin.Y + i,
		})
	}

	return a
}
