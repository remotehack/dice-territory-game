package game_test

import (
	"reflect"
	"testing"

	"javorszky/dice-territory-game/v2/pkg/game"
)

func TestNewPiece(t *testing.T) {
	type args struct {
		player  game.Player
		originX uint8
		originY uint8
		width   uint8
		height  uint8
	}
	tests := []struct {
		name string
		args args
		want game.Piece
	}{
		{
			name: "successfully create piece",
			args: args{
				player: game.Player{
					ID:    "playerid-1",
					Name:  "John",
					Score: 0,
				},
				originX: 4,
				originY: 6,
				width:   4,
				height:  3,
			},
			want: game.Piece{
				Player: game.Player{
					ID:    "playerid-1",
					Name:  "John",
					Score: 0,
				},
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
				Width:  4,
				Height: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := game.NewPiece(tt.args.player, tt.args.originX, tt.args.originY, tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPiece() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

//
//func TestPiece_IsAdjacent(t *testing.T) {
//	type fields struct {
//		Player         game.Player
//		Origin         game.Coordinate
//		AdjacentFields []Coordinate
//		Corners        []Coordinate
//		Width          uint8
//		Height         uint8
//	}
//	type args struct {
//		c game.Coordinate
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := game.Piece{
//				Player:         tt.fields.Player,
//				Origin:         tt.fields.Origin,
//				AdjacentFields: tt.fields.AdjacentFields,
//				Corners:        tt.fields.Corners,
//				Width:          tt.fields.Width,
//				Height:         tt.fields.Height,
//			}
//			if got := p.IsAdjacent(tt.args.c); got != tt.want {
//				t.Errorf("IsAdjacent() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestPiece_IsCoordinateWithin(t *testing.T) {
//	type fields struct {
//		Player         game.Player
//		Origin         game.Coordinate
//		AdjacentFields []Coordinate
//		Corners        []Coordinate
//		Width          uint8
//		Height         uint8
//	}
//	type args struct {
//		c game.Coordinate
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := game.Piece{
//				Player:         tt.fields.Player,
//				Origin:         tt.fields.Origin,
//				AdjacentFields: tt.fields.AdjacentFields,
//				Corners:        tt.fields.Corners,
//				Width:          tt.fields.Width,
//				Height:         tt.fields.Height,
//			}
//			if got := p.IsCoordinateWithin(tt.args.c); got != tt.want {
//				t.Errorf("IsCoordinateWithin() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestPiece_getAdjacentFields(t *testing.T) {
//	type fields struct {
//		Player         game.Player
//		Origin         game.Coordinate
//		AdjacentFields []Coordinate
//		Corners        []Coordinate
//		Width          uint8
//		Height         uint8
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   []Coordinate
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := game.Piece{
//				Player:         tt.fields.Player,
//				Origin:         tt.fields.Origin,
//				AdjacentFields: tt.fields.AdjacentFields,
//				Corners:        tt.fields.Corners,
//				Width:          tt.fields.Width,
//				Height:         tt.fields.Height,
//			}
//			if got := p.getAdjacentFields(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("getAdjacentFields() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestPiece_getCorners(t *testing.T) {
//	type fields struct {
//		Player         game.Player
//		Origin         game.Coordinate
//		AdjacentFields []Coordinate
//		Corners        []Coordinate
//		Width          uint8
//		Height         uint8
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   []Coordinate
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := game.Piece{
//				Player:         tt.fields.Player,
//				Origin:         tt.fields.Origin,
//				AdjacentFields: tt.fields.AdjacentFields,
//				Corners:        tt.fields.Corners,
//				Width:          tt.fields.Width,
//				Height:         tt.fields.Height,
//			}
//			if got := p.getCorners(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("getCorners() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
