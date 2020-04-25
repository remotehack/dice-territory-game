package game_test

import (
	"reflect"
	"testing"

	"javorszky/dice-territory-game/v2/pkg/game"
)

func TestNew(t *testing.T) {
	type args struct {
		Session     string
		BoardHeight uint8
		BoardWidth  uint8
		PlayerOne   string
		PlayerTwo   string
	}
	tests := []struct {
		name    string
		args    args
		want    game.Game
		wantErr bool
	}{
		{
			name: "New Game with correct size",
			args: args{
				Session:     "1",
				BoardHeight: 12,
				BoardWidth:  16,
				PlayerOne:   "playerOne",
				PlayerTwo:   "playerTwo",
			},
			want: game.Game{
				Session: "1",
				Board: game.Board{
					Width:  12,
					Height: 16,
					Pieces: []game.Piece{},
					Corners: []game.Coordinate{
						{
							X: 1,
							Y: 1,
						},
						{
							X: 12,
							Y: 16,
						},
					},
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
			},
			wantErr: false,
		},
		{
			name: "New Game error with too small width",
			args: args{

				Session:     "1",
				BoardHeight: 12,
				BoardWidth:  4,
				PlayerOne:   "iii",
				PlayerTwo:   "ccc",
			},
			want:    game.Game{},
			wantErr: true,
		},
		{
			name: "New Game error with too large width",
			args: args{

				Session:     "1",
				BoardHeight: 12,
				BoardWidth:  127,
				PlayerOne:   "iii",
				PlayerTwo:   "ccc",
			},
			want:    game.Game{},
			wantErr: true,
		},
		{
			name: "New Game error with too small height",
			args: args{

				Session:     "1",
				BoardHeight: 4,
				BoardWidth:  16,
				PlayerOne:   "iii",
				PlayerTwo:   "ccc",
			},
			want:    game.Game{},
			wantErr: true,
		},
		{
			name: "New Game error with too large height",
			args: args{

				Session:     "1",
				BoardHeight: 127,
				BoardWidth:  16,
				PlayerOne:   "iii",
				PlayerTwo:   "ccc",
			},
			want:    game.Game{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := game.New(tt.args.Session, tt.args.BoardHeight, tt.args.BoardWidth, tt.args.PlayerOne, tt.args.PlayerTwo)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
