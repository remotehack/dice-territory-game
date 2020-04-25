package game

import "fmt"

const (
	minBoardHeight = 12
	maxBoardHeight = 30
	minBoardWidth  = 8
	maxBoardWidth  = 24
)

type Game struct {
	Session     string
	BoardHeight int8
	BoardWidth  int8
	PlayerOne   string
	PlayerTwo   string
}

func New(session string, boardHeight int8, boardWidth int8, playerOne string, playerTwo string) (Game, error) {
	if !isBetween(boardHeight, minBoardHeight, maxBoardHeight) {
		return Game{}, fmt.Errorf("board height is out of bounds. Got %d, need to be between %d and %d inclusive", boardHeight, minBoardHeight, maxBoardHeight)
	}
	if !isBetween(boardWidth, minBoardWidth, maxBoardWidth) {
		return Game{}, fmt.Errorf("board width is out of bounds. Got %d, need to be between %d and %d inclusive", boardWidth, minBoardWidth, maxBoardWidth)
	}

	return Game{Session: session, BoardHeight: boardHeight, BoardWidth: boardWidth, PlayerOne: playerOne, PlayerTwo: playerTwo}, nil
}

func isBetween(what int8, lowerBoundary int8, upperBoundary int8) bool {
	return what >= lowerBoundary && what <= upperBoundary
}
