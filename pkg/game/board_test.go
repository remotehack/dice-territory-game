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
		p1     string
		p2     string
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
				p1:     "playerOne",
				p2:     "playerTwo",
			},
			want: game.Board{
				Width:  12,
				Height: 16,
				Corners: []game.Coordinate{
					{X: 1, Y: 1},
					{X: 12, Y: 16},
				},
				Pieces: []game.Piece{},
				Players: []game.Player{
					{
						ID:    "playerOne",
						Name:  "playerOne",
						Score: 0,
					},
					{
						ID:    "playerTwo",
						Name:  "playerTwo",
						Score: 0,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Error, board too small width",
			args: args{
				width:  4,
				height: 16,
				p1:     "playerOne",
				p2:     "playerTwo",
			},
			want:    game.Board{},
			wantErr: true,
		},
		{
			name: "Error, board too large width",
			args: args{
				width:  127,
				height: 16,
				p1:     "playerOne",
				p2:     "playerTwo",
			},
			want:    game.Board{},
			wantErr: true,
		},
		{
			name: "Error, board too small height",
			args: args{
				width:  12,
				height: 4,
				p1:     "playerOne",
				p2:     "playerTwo",
			},
			want:    game.Board{},
			wantErr: true,
		},
		{
			name: "Error, board too large height",
			args: args{
				width:  12,
				height: 127,
				p1:     "playerOne",
				p2:     "playerTwo",
			},
			want:    game.Board{},
			wantErr: true,
		},
		{
			name: "Error, player one has no name",
			args: args{
				width:  12,
				height: 16,
				p1:     "",
				p2:     "playerTwo",
			},
			want:    game.Board{},
			wantErr: true,
		},
		{
			name: "Error, player two has no name",
			args: args{
				width:  12,
				height: 16,
				p1:     "playerOne",
				p2:     "",
			},
			want:    game.Board{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := game.NewBoard(tt.args.width, tt.args.height, tt.args.p1, tt.args.p2)
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
	piece1 := game.Piece{
		Player: "id-player-1",
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
	}
	piece2 := game.Piece{
		Player: "id-player-2",
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
	}

	newPieceP1 := game.Piece{
		Player: "id-player-1",
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
	}

	newPieceP2 := game.Piece{
		Player: "id-player-2",
		Origin: game.Coordinate{X: 3, Y: 12},
		AdjacentFields: []game.Coordinate{
			{X: 3, Y: 11},
			{X: 3, Y: 13},
			{X: 4, Y: 11},
			{X: 4, Y: 13},
			{X: 5, Y: 11},
			{X: 5, Y: 13},
			{X: 6, Y: 11},
			{X: 6, Y: 13},

			{X: 2, Y: 12},
			{X: 7, Y: 12},
		},
		Corners: []game.Coordinate{
			{X: 3, Y: 12},
			{X: 7, Y: 12},
			{X: 7, Y: 12},
			{X: 3, Y: 12},
		},
		Coordinates: []game.Coordinate{
			{X: 3, Y: 12},
			{X: 4, Y: 12},
			{X: 5, Y: 12},
			{X: 6, Y: 12},
		},
		Width:  4,
		Height: 1,
	}

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
			piece1,
			piece2,
		},
	}

	type args struct {
		p game.Piece
	}
	tests := []struct {
		name    string
		b       game.Board
		args    args
		want    game.Board
		wantErr bool
	}{
		{
			name: "can place piece for player 1",
			b:    b,
			args: args{
				p: newPieceP1,
			},
			want: game.Board{
				Width:  8,
				Height: 12,
				Corners: []game.Coordinate{
					{X: 1, Y: 1},
					{X: 8, Y: 1},
					{X: 8, Y: 12},
					{X: 1, Y: 12},
				},
				Pieces: []game.Piece{
					piece1,
					piece2,
					newPieceP1,
				},
			},
			wantErr: false,
		},
		{
			name: "can place piece for player 2",
			b:    b,
			args: args{
				p: newPieceP2,
			},
			want: game.Board{
				Width:  8,
				Height: 12,
				Corners: []game.Coordinate{
					{X: 1, Y: 1},
					{X: 8, Y: 1},
					{X: 8, Y: 12},
					{X: 1, Y: 12},
				},
				Pieces: []game.Piece{
					piece1,
					piece2,
					newPieceP2,
				},
			},
			wantErr: false,
		},
		{
			name: "can not place piece for player 2 (intersects)",
			b:    b,
			args: args{
				p: game.Piece{
					Player: "id-player-2",
					Origin: game.Coordinate{X: 4, Y: 12},
					AdjacentFields: []game.Coordinate{
						{X: 4, Y: 11},
						{X: 4, Y: 13},
						{X: 5, Y: 11},
						{X: 5, Y: 13},
						{X: 6, Y: 11},
						{X: 6, Y: 13},
						{X: 7, Y: 11},
						{X: 7, Y: 13},

						{X: 3, Y: 12},
						{X: 8, Y: 12},
					},
					Corners: []game.Coordinate{
						{X: 4, Y: 12},
						{X: 8, Y: 12},
						{X: 8, Y: 12},
						{X: 4, Y: 12},
					},
					Coordinates: []game.Coordinate{
						{X: 4, Y: 12},
						{X: 5, Y: 12},
						{X: 6, Y: 12},
						{X: 7, Y: 12},
					},
					Width:  4,
					Height: 1,
				},
			},
			want:    b,
			wantErr: true,
		},
		{
			name: "can not place piece for player 2 (does not touch anything)",
			b:    b,
			args: args{
				p: game.Piece{
					Player: "id-player-2",
					Origin: game.Coordinate{X: 4, Y: 7},
					AdjacentFields: []game.Coordinate{
						{X: 4, Y: 6},
						{X: 4, Y: 8},
						{X: 5, Y: 6},
						{X: 5, Y: 8},
						{X: 6, Y: 6},
						{X: 6, Y: 8},
						{X: 7, Y: 6},
						{X: 7, Y: 8},

						{X: 3, Y: 7},
						{X: 8, Y: 7},
					},
					Corners: []game.Coordinate{
						{X: 4, Y: 7},
						{X: 8, Y: 7},
						{X: 8, Y: 7},
						{X: 4, Y: 7},
					},
					Coordinates: []game.Coordinate{
						{X: 4, Y: 7},
						{X: 5, Y: 7},
						{X: 6, Y: 7},
						{X: 7, Y: 7},
					},
					Width:  4,
					Height: 1,
				},
			},
			want:    b,
			wantErr: true,
		},
		{
			name: "can not place piece for player 2 (touches p1's, but not p2's)",
			b:    b,
			args: args{
				p: game.Piece{
					Player: "id-player-2",
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
			want:    b,
			wantErr: true,
		},
		{
			name: "can not place piece for player 2 (out of bounds of the board)",
			b:    b,
			args: args{
				p: game.Piece{
					Player: "id-player-2",
					Origin: game.Coordinate{X: 6, Y: 3},
					AdjacentFields: []game.Coordinate{
						{X: 6, Y: 2},
						{X: 6, Y: 4},
						{X: 7, Y: 2},
						{X: 7, Y: 4},
						{X: 8, Y: 2},
						{X: 8, Y: 4},
						{X: 9, Y: 2},
						{X: 9, Y: 4},
						{X: 5, Y: 3},
						{X: 10, Y: 3},
					},
					Corners: []game.Coordinate{
						{X: 6, Y: 3},
						{X: 9, Y: 3},
						{X: 9, Y: 3},
						{X: 6, Y: 3},
					},
					Coordinates: []game.Coordinate{
						{X: 6, Y: 3},
						{X: 7, Y: 3},
						{X: 8, Y: 3},
						{X: 9, Y: 3},
					},
					Width:  4,
					Height: 1,
				},
			},
			want:    b,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localBoard := b

			got, err := localBoard.PlacePiece(tt.args.p)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewPiece() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("NewPiece() got = \n%#v, want \n%#v", *got, tt.want)
			}
		})
	}
}

func TestBoard_IsCoordinateWithin(t *testing.T) {
	b := game.Board{
		Width:   8,
		Height:  12,
		Players: nil,
		Corners: []game.Coordinate{
			{
				X: 1,
				Y: 1,
			},
			{
				X: 8,
				Y: 12,
			},
		},
		Pieces: nil,
	}

	type args struct {
		cs []game.Coordinate
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "coordinates are within the board",
			args: args{
				cs: []game.Coordinate{
					{X: 1, Y: 1},
					{X: 1, Y: 2},
					{X: 1, Y: 3},
					{X: 1, Y: 4},
					{X: 1, Y: 5},
					{X: 1, Y: 6},
					{X: 1, Y: 7},
					{X: 1, Y: 8},
					{X: 1, Y: 9},
					{X: 1, Y: 10},
					{X: 1, Y: 11},
					{X: 1, Y: 12},
					{X: 2, Y: 1},
					{X: 2, Y: 2},
					{X: 2, Y: 3},
					{X: 2, Y: 4},
					{X: 2, Y: 5},
					{X: 2, Y: 6},
					{X: 2, Y: 7},
					{X: 2, Y: 8},
					{X: 2, Y: 9},
					{X: 2, Y: 10},
					{X: 2, Y: 11},
					{X: 2, Y: 12},
					{X: 3, Y: 1},
					{X: 3, Y: 2},
					{X: 3, Y: 3},
					{X: 3, Y: 4},
					{X: 3, Y: 5},
					{X: 3, Y: 6},
					{X: 3, Y: 7},
					{X: 3, Y: 8},
					{X: 3, Y: 9},
					{X: 3, Y: 10},
					{X: 3, Y: 11},
					{X: 3, Y: 12},
					{X: 4, Y: 1},
					{X: 4, Y: 2},
					{X: 4, Y: 3},
					{X: 4, Y: 4},
					{X: 4, Y: 5},
					{X: 4, Y: 6},
					{X: 4, Y: 7},
					{X: 4, Y: 8},
					{X: 4, Y: 9},
					{X: 4, Y: 10},
					{X: 4, Y: 11},
					{X: 4, Y: 12},
					{X: 5, Y: 1},
					{X: 5, Y: 2},
					{X: 5, Y: 3},
					{X: 5, Y: 4},
					{X: 5, Y: 5},
					{X: 5, Y: 6},
					{X: 5, Y: 7},
					{X: 5, Y: 8},
					{X: 5, Y: 9},
					{X: 5, Y: 10},
					{X: 5, Y: 11},
					{X: 5, Y: 12},
					{X: 6, Y: 1},
					{X: 6, Y: 2},
					{X: 6, Y: 3},
					{X: 6, Y: 4},
					{X: 6, Y: 5},
					{X: 6, Y: 6},
					{X: 6, Y: 7},
					{X: 6, Y: 8},
					{X: 6, Y: 9},
					{X: 6, Y: 10},
					{X: 6, Y: 11},
					{X: 6, Y: 12},
					{X: 7, Y: 1},
					{X: 7, Y: 2},
					{X: 7, Y: 3},
					{X: 7, Y: 4},
					{X: 7, Y: 5},
					{X: 7, Y: 6},
					{X: 7, Y: 7},
					{X: 7, Y: 8},
					{X: 7, Y: 9},
					{X: 7, Y: 10},
					{X: 7, Y: 11},
					{X: 7, Y: 12},
					{X: 8, Y: 1},
					{X: 8, Y: 2},
					{X: 8, Y: 3},
					{X: 8, Y: 4},
					{X: 8, Y: 5},
					{X: 8, Y: 6},
					{X: 8, Y: 7},
					{X: 8, Y: 8},
					{X: 8, Y: 9},
					{X: 8, Y: 10},
					{X: 8, Y: 11},
					{X: 8, Y: 12},
				},
			},
			want: true,
		},
		{
			name: "coordinates are outside the board",
			args: args{
				cs: []game.Coordinate{
					{X: 1, Y: 0},
					{X: 1, Y: 13},
					{X: 2, Y: 0},
					{X: 2, Y: 13},
					{X: 3, Y: 0},
					{X: 3, Y: 13},
					{X: 4, Y: 0},
					{X: 4, Y: 13},
					{X: 5, Y: 0},
					{X: 5, Y: 13},
					{X: 6, Y: 0},
					{X: 6, Y: 13},
					{X: 7, Y: 0},
					{X: 7, Y: 13},
					{X: 8, Y: 0},
					{X: 8, Y: 13},
					{X: 0, Y: 1},
					{X: 9, Y: 1},
					{X: 0, Y: 2},
					{X: 9, Y: 2},
					{X: 0, Y: 3},
					{X: 9, Y: 3},
					{X: 0, Y: 4},
					{X: 9, Y: 4},
					{X: 0, Y: 5},
					{X: 9, Y: 5},
					{X: 0, Y: 6},
					{X: 9, Y: 6},
					{X: 0, Y: 7},
					{X: 9, Y: 7},
					{X: 0, Y: 8},
					{X: 9, Y: 8},
					{X: 0, Y: 9},
					{X: 9, Y: 9},
					{X: 0, Y: 10},
					{X: 9, Y: 10},
					{X: 0, Y: 11},
					{X: 9, Y: 11},
					{X: 0, Y: 12},
					{X: 9, Y: 12},
					{X: 0, Y: 0},
					{X: 0, Y: 13},
					{X: 9, Y: 0},
					{X: 9, Y: 13},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localBoard := b

			for _, c := range tt.args.cs {
				if got := localBoard.IsCoordinateWithin(c); got != tt.want {
					t.Errorf("IsCoordinateWithin() = %v, want %v, coordinate %v", got, tt.want, c)
				}
			}
		})
	}
}
