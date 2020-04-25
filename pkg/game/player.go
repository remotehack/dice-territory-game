package game

import (
	"errors"
)

type Player struct {
	ID    string
	Name  string
	Score uint
}

func NewPlayer(ID string, name string) (Player, error) {
	if "" == name {
		return Player{}, errors.New("can't create player with empty name")
	}
	return Player{ID: ID, Name: name, Score: 0}, nil
}
