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
		name    string
		args    args
		want    game.Piece
		wantErr bool
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
			wantErr: false,
		},
		{
			name: "can't create piece with too short width",
			args: args{
				player: game.Player{
					ID:    "playerid-1",
					Name:  "Bob",
					Score: 0,
				},
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
				player: game.Player{
					ID:    "playerid-1",
					Name:  "Bob",
					Score: 0,
				},
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
				player: game.Player{
					ID:    "playerid-1",
					Name:  "Bob",
					Score: 0,
				},
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
				player: game.Player{
					ID:    "playerid-1",
					Name:  "Bob",
					Score: 0,
				},
				originX: 3,
				originY: 3,
				width:   10,
				height:  3,
			},
			want:    game.Piece{},
			wantErr: true,
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
				t.Errorf("NewPiece() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPiece_IsAdjacent(t *testing.T) {

	type args struct {
		p game.Piece
		c game.Coordinate
	}

	p := game.Piece{
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
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found adjacent coordinate on top",
			args: args{
				p: p,
				c: game.Coordinate{X: 6, Y: 5},
			},
			want: true,
		},
		{
			name: "found adjacent coordinate on bottom",
			args: args{
				p: p,
				c: game.Coordinate{X: 6, Y: 9},
			},
			want: true,
		},
		{
			name: "found adjacent coordinate on left",
			args: args{
				p: p,
				c: game.Coordinate{X: 3, Y: 7},
			},
			want: true,
		},
		{
			name: "found adjacent coordinate on right",
			args: args{
				p: p,
				c: game.Coordinate{X: 8, Y: 7},
			},
			want: true,
		},
		{
			name: "not found adjacent coordinate in corner",
			args: args{
				p: p,
				c: game.Coordinate{X: 3, Y: 5},
			},
			want: false,
		},
		{
			name: "not found adjacent coordinate on top: too far",
			args: args{
				p: p,
				c: game.Coordinate{X: 6, Y: 4},
			},
			want: false,
		},
		{
			name: "not found adjacent coordinate on bottom: too far",
			args: args{
				p: p,
				c: game.Coordinate{X: 6, Y: 10},
			},
			want: false,
		},
		{
			name: "not found adjacent coordinate on left: too far",
			args: args{
				p: p,
				c: game.Coordinate{X: 2, Y: 7},
			},
			want: false,
		},
		{
			name: "not found adjacent coordinate on right: too far",
			args: args{
				p: p,
				c: game.Coordinate{X: 9, Y: 7},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.p.IsAdjacent(tt.args.c); got != tt.want {
				t.Errorf("IsAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
