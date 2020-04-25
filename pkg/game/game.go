package game

import "fmt"

type Game struct {
	Session string
	Board   Board
}

func New(session string, boardWidth uint8, boardHeight uint8, playerOne string, playerTwo string) (Game, error) {
	board, err := NewBoard(boardWidth, boardHeight, playerOne, playerTwo)
	if err != nil {
		return Game{}, fmt.Errorf("game.New(): NewBoard(): %w", err)
	}
	return Game{Session: session, Board: board}, nil
}
