package game_test

import (
	"reflect"
	"testing"

	"javorszky/dice-territory-game/v2/pkg/game"
)

func TestNewPiece(t *testing.T) {
	type args struct {
		player  string
		originX uint8
		originY uint8
		width   uint8
		height  uint8
	}
	tests := []struct {
		name    string
		args    args
		want    game.Piece
		wantErr bool
	}{
		{
			name: "successfully create piece",
			args: args{
				player:  "playerid-1",
				originX: 4,
				originY: 6,
				width:   4,
				height:  3,
			},
			want: game.Piece{
				Player: "playerid-1",
				Origin: game.Coordinate{X: 4, Y: 6},
				AdjacentFields: []game.Coordinate{
					{X: 4, Y: 5},
					{X: 4, Y: 9},
					{X: 5, Y: 5},
					{X: 5, Y: 9},
					{X: 6, Y: 5},
					{X: 6, Y: 9},
					{X: 7, Y: 5},
					{X: 7, Y: 9},
					{X: 3, Y: 6},
					{X: 8, Y: 6},
					{X: 3, Y: 7},
					{X: 8, Y: 7},
					{X: 3, Y: 8},
					{X: 8, Y: 8},
				},
				Corners: []game.Coordinate{
					{X: 4, Y: 6},
					{X: 7, Y: 6},
					{X: 7, Y: 8},
					{X: 4, Y: 8},
				},
				Coordinates: []game.Coordinate{
					{X: 4, Y: 6},
					{X: 4, Y: 7},
					{X: 4, Y: 8},
					{X: 5, Y: 6},
					{X: 5, Y: 7},
					{X: 5, Y: 8},
					{X: 6, Y: 6},
					{X: 6, Y: 7},
					{X: 6, Y: 8},
					{X: 7, Y: 6},
					{X: 7, Y: 7},
					{X: 7, Y: 8},
				},
				Width:  4,
				Height: 3,
			},
			wantErr: false,
		},
		{
			name: "can't create piece with too short width",
			args: args{
				player:  "playerid-1",
				originX: 3,
				originY: 3,
				width:   0,
				height:  6,
			},
			want:    game.Piece{},
			wantErr: true,
		},
		{
			name: "can't create piece with too short height",
			args: args{
				player:  "playerid-1",
				originX: 3,
				originY: 3,
				width:   3,
				height:  0,
			},
			want:    game.Piece{},
			wantErr: true,
		},
		{
			name: "can't create piece with too long height",
			args: args{
				player:  "playerid-1",
				originX: 3,
				originY: 3,
				width:   3,
				height:  10,
			},
			want:    game.Piece{},
			wantErr: true,
		},
		{
			name: "can't create piece with too long width",
			args: args{
				player:  "playerid-1",
				originX: 3,
				originY: 3,
				width:   10,
				height:  3,
			},
			want:    game.Piece{},
			wantErr: true,
		},
		{
			name: "successfully create 1x1 piece",
			args: args{
				player:  "playerid-1",
				originX: 4,
				originY: 6,
				width:   1,
				height:  1,
			},
			want: game.Piece{
				Player: "playerid-1",
				Origin: game.Coordinate{X: 4, Y: 6},
				AdjacentFields: []game.Coordinate{
					{X: 4, Y: 5},
					{X: 4, Y: 7},
					{X: 3, Y: 6},
					{X: 5, Y: 6},
				},
				Corners: []game.Coordinate{
					{X: 4, Y: 6},
					{X: 4, Y: 6},
					{X: 4, Y: 6},
					{X: 4, Y: 6},
				},
				Coordinates: []game.Coordinate{
					{X: 4, Y: 6},
				},
				Width:  1,
				Height: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := game.NewPiece(tt.args.player, tt.args.originX, tt.args.originY, tt.args.width, tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPiece() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPiece() got = \n%#v, want \n%#v", got, tt.want)
			}
		})
	}
}

func TestPiece_IsAdjacent(t *testing.T) {

	type args struct {
		p  game.Piece
		cs []game.Coordinate
	}

	p := game.Piece{
		Player: "playerid-1",
		Origin: game.Coordinate{X: 4, Y: 6},
		AdjacentFields: []game.Coordinate{
			{X: 4, Y: 5},
			{X: 4, Y: 9},
			{X: 5, Y: 5},
			{X: 5, Y: 9},
			{X: 6, Y: 5},
			{X: 6, Y: 9},
			{X: 7, Y: 5},
			{X: 7, Y: 9},
			{X: 3, Y: 6},
			{X: 8, Y: 6},
			{X: 3, Y: 7},
			{X: 8, Y: 7},
			{X: 3, Y: 8},
			{X: 8, Y: 8},
		},
		Corners: []game.Coordinate{
			{X: 4, Y: 6},
			{X: 7, Y: 6},
			{X: 7, Y: 8},
			{X: 4, Y: 8},
		},
		Coordinates: []game.Coordinate{
			{X: 4, Y: 6},
			{X: 5, Y: 6},
			{X: 6, Y: 6},
			{X: 7, Y: 6},

			{X: 4, Y: 7},
			{X: 5, Y: 7},
			{X: 6, Y: 7},
			{X: 7, Y: 7},

			{X: 4, Y: 8},
			{X: 5, Y: 8},
			{X: 6, Y: 8},
			{X: 7, Y: 8},
		},
		Width:  4,
		Height: 3,
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "points found adjacent to piece",
			args: args{
				p: p,
				cs: []game.Coordinate{
					{X: 4, Y: 5},
					{X: 4, Y: 9},
					{X: 5, Y: 5},
					{X: 5, Y: 9},
					{X: 6, Y: 5},
					{X: 6, Y: 9},
					{X: 7, Y: 5},
					{X: 7, Y: 9},
					{X: 3, Y: 6},
					{X: 8, Y: 6},
					{X: 3, Y: 7},
					{X: 8, Y: 7},
					{X: 3, Y: 8},
					{X: 8, Y: 8},
				},
			},
			want: true,
		},
		{
			name: "points found not adjacent (inside)",
			args: args{
				p: p,
				cs: []game.Coordinate{
					{X: 4, Y: 6},
					{X: 5, Y: 6},
					{X: 6, Y: 6},
					{X: 7, Y: 6},

					{X: 4, Y: 7},
					{X: 5, Y: 7},
					{X: 6, Y: 7},
					{X: 7, Y: 7},

					{X: 4, Y: 8},
					{X: 5, Y: 8},
					{X: 6, Y: 8},
					{X: 7, Y: 8},
				},
			},
			want: false,
		},
		{
			name: "points found not adjacent (outside, far, corner)",
			args: args{
				p: p,
				cs: []game.Coordinate{
					{X: 4, Y: 4},
					{X: 4, Y: 10},
					{X: 5, Y: 4},
					{X: 5, Y: 10},
					{X: 6, Y: 4},
					{X: 6, Y: 10},
					{X: 7, Y: 4},
					{X: 7, Y: 10},
					{X: 2, Y: 6},
					{X: 9, Y: 6},
					{X: 2, Y: 7},
					{X: 9, Y: 7},
					{X: 2, Y: 8},
					{X: 9, Y: 8},

					{X: 3, Y: 5},
					{X: 8, Y: 5},
					{X: 8, Y: 9},
					{X: 3, Y: 9},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, c := range tt.args.cs {
				if got := tt.args.p.IsAdjacent(c); got != tt.want {
					t.Errorf("IsAdjacent() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestPiece_IsCoordinateWithin(t *testing.T) {
	p := game.Piece{
		Player: "playerid-1",
		Origin: game.Coordinate{X: 4, Y: 6},
		AdjacentFields: []game.Coordinate{
			{X: 4, Y: 5},
			{X: 4, Y: 9},
			{X: 5, Y: 5},
			{X: 5, Y: 9},
			{X: 6, Y: 5},
			{X: 6, Y: 9},
			{X: 7, Y: 5},
			{X: 7, Y: 9},
			{X: 3, Y: 6},
			{X: 8, Y: 6},
			{X: 3, Y: 7},
			{X: 8, Y: 7},
			{X: 3, Y: 8},
			{X: 8, Y: 8},
		},
		Corners: []game.Coordinate{
			{X: 4, Y: 6},
			{X: 7, Y: 6},
			{X: 7, Y: 8},
			{X: 4, Y: 8},
		},
		Coordinates: []game.Coordinate{
			{X: 4, Y: 6},
			{X: 5, Y: 6},
			{X: 6, Y: 6},
			{X: 7, Y: 6},

			{X: 4, Y: 7},
			{X: 5, Y: 7},
			{X: 6, Y: 7},
			{X: 7, Y: 7},

			{X: 4, Y: 8},
			{X: 5, Y: 8},
			{X: 6, Y: 8},
			{X: 7, Y: 8},
		},
		Width:  4,
		Height: 3,
	}

	type args struct {
		p  game.Piece
		cs []game.Coordinate
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "coordinates are within the piece",
			args: args{
				p: p,
				cs: []game.Coordinate{
					{X: 4, Y: 6},
					{X: 5, Y: 6},
					{X: 6, Y: 6},
					{X: 7, Y: 6},

					{X: 4, Y: 7},
					{X: 5, Y: 7},
					{X: 6, Y: 7},
					{X: 7, Y: 7},

					{X: 4, Y: 8},
					{X: 5, Y: 8},
					{X: 6, Y: 8},
					{X: 7, Y: 8},
				},
			},
			want: true,
		},
		{
			name: "coordinates are not within the piece",
			args: args{
				p: p,
				cs: []game.Coordinate{
					{X: 4, Y: 5},
					{X: 4, Y: 9},
					{X: 5, Y: 5},
					{X: 5, Y: 9},
					{X: 6, Y: 5},
					{X: 6, Y: 9},
					{X: 7, Y: 5},
					{X: 7, Y: 9},
					{X: 3, Y: 6},
					{X: 8, Y: 6},
					{X: 3, Y: 7},
					{X: 8, Y: 7},
					{X: 3, Y: 8},
					{X: 8, Y: 8},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, c := range tt.args.cs {
				if got := tt.args.p.IsCoordinateWithin(c); got != tt.want {
					t.Errorf("IsCoordinateWithin() = %v, want %v,\ncoordinate: %#v\n\n", got, tt.want, c)
				}
			}
		})
	}
}
