package game_test

import (
	"reflect"
	"testing"

	"javorszky/dice-territory-game/v2/pkg/game"
)

func TestNewBoard(t *testing.T) {
	type args struct {
		width  uint8
		height uint8
	}
	tests := []struct {
		name    string
		args    args
		want    game.Board
		wantErr bool
	}{
		{
			name: "Valid sized board",
			args: args{
				width:  12,
				height: 16,
			},
			want: game.Board{
				Width:  12,
				Height: 16,
				Corners: []game.Coordinate{
					{X: 1, Y: 1},
					{X: 12, Y: 16},
				},
				Pieces: []game.Piece{},
			},
			wantErr: false,
		},
		{
			name: "Error, board too small width",
			args: args{
				width:  4,
				height: 16,
			},
			want:    game.Board{},
			wantErr: true,
		},
		{
			name: "Error, board too large width",
			args: args{
				width:  127,
				height: 16,
			},
			want:    game.Board{},
			wantErr: true,
		},
		{
			name: "Error, board too small height",
			args: args{
				width:  12,
				height: 4,
			},
			want:    game.Board{},
			wantErr: true,
		},
		{
			name: "Error, board too large height",
			args: args{
				width:  12,
				height: 127,
			},
			want:    game.Board{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := game.NewBoard(tt.args.width, tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_PlacePiece(t *testing.T) {
	b := game.Board{
		Width:  8,
		Height: 12,
		Corners: []game.Coordinate{
			{X: 1, Y: 1},
			{X: 8, Y: 1},
			{X: 8, Y: 12},
			{X: 1, Y: 12},
		},
		Pieces: []game.Piece{
			{
				Player: game.Player{
					ID:    "id-player-1",
					Name:  "Alice",
					Score: 0,
				},
				Origin: game.Coordinate{X: 1, Y: 1},
				AdjacentFields: []game.Coordinate{
					{X: 1, Y: 0},
					{X: 1, Y: 5},
					{X: 2, Y: 0},
					{X: 2, Y: 5},
					{X: 3, Y: 0},
					{X: 3, Y: 5},
					{X: 4, Y: 0},
					{X: 4, Y: 5},

					{X: 0, Y: 1},
					{X: 5, Y: 1},
					{X: 0, Y: 2},
					{X: 5, Y: 2},
					{X: 0, Y: 3},
					{X: 5, Y: 3},
					{X: 0, Y: 4},
					{X: 5, Y: 4},
				},
				Corners: []game.Coordinate{
					{X: 1, Y: 1},
					{X: 4, Y: 1},
					{X: 4, Y: 4},
					{X: 1, Y: 4},
				},
				Coordinates: []game.Coordinate{
					{X: 1, Y: 1},
					{X: 2, Y: 1},
					{X: 3, Y: 1},
					{X: 4, Y: 1},
					{X: 1, Y: 2},
					{X: 2, Y: 2},
					{X: 3, Y: 2},
					{X: 4, Y: 2},
					{X: 1, Y: 3},
					{X: 2, Y: 3},
					{X: 3, Y: 3},
					{X: 4, Y: 3},
					{X: 1, Y: 4},
					{X: 2, Y: 4},
					{X: 3, Y: 4},
					{X: 4, Y: 4},
				},
				Width:  4,
				Height: 4,
			},
			{
				Player: game.Player{
					ID:    "id-player-2",
					Name:  "Bob",
					Score: 0,
				},
				Origin: game.Coordinate{X: 7, Y: 11},
				AdjacentFields: []game.Coordinate{
					{X: 7, Y: 10},
					{X: 7, Y: 13},
					{X: 8, Y: 10},
					{X: 8, Y: 13},

					{X: 6, Y: 11},
					{X: 9, Y: 11},
					{X: 6, Y: 12},
					{X: 9, Y: 12},
				},
				Corners: []game.Coordinate{
					{X: 7, Y: 11},
					{X: 8, Y: 11},
					{X: 8, Y: 12},
					{X: 7, Y: 12},
				},
				Coordinates: []game.Coordinate{
					{X: 7, Y: 11},
					{X: 8, Y: 11},
					{X: 8, Y: 12},
					{X: 7, Y: 12},
				},
				Width:  2,
				Height: 2,
			},
		},
	}

	type args struct {
		p game.Piece
	}
	tests := []struct {
		name string
		b    game.Board
		args args
		want bool
	}{
		{
			name: "can place piece for player 1",
			b:    b,
			args: args{
				p: game.Piece{
					Player: game.Player{
						ID:    "id-player-1",
						Name:  "Alice",
						Score: 0,
					},
					Origin: game.Coordinate{X: 5, Y: 3},
					AdjacentFields: []game.Coordinate{
						{X: 5, Y: 2},
						{X: 5, Y: 4},
						{X: 6, Y: 2},
						{X: 6, Y: 4},
						{X: 7, Y: 2},
						{X: 7, Y: 4},
						{X: 8, Y: 2},
						{X: 8, Y: 4},
						{X: 4, Y: 3},
						{X: 9, Y: 3},
					},
					Corners: []game.Coordinate{
						{X: 5, Y: 3},
						{X: 8, Y: 3},
						{X: 8, Y: 3},
						{X: 5, Y: 3},
					},
					Coordinates: []game.Coordinate{
						{X: 5, Y: 3},
						{X: 6, Y: 3},
						{X: 7, Y: 3},
						{X: 8, Y: 3},
					},
					Width:  4,
					Height: 1,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.PlacePiece(tt.args.p); got != tt.want {
				t.Errorf("PlacePiece() = %v, want %v", got, tt.want)
			}
		})
	}
}
