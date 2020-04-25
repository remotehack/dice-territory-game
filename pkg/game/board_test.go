package game_test

import (
	"javorszky/dice-territory-game/v2/pkg/game"
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	type args struct {
		width  int8
		height int8
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
