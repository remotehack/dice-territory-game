package game_test

import (
	"javorszky/dice-territory-game/v2/pkg/game"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		Session     string
		BoardHeight int8
		BoardWidth  int8
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
				PlayerOne:   "iii",
				PlayerTwo:   "ccc",
			},
			want: game.Game{
				Session:     "1",
				BoardHeight: 12,
				BoardWidth:  16,
				PlayerOne:   "iii",
				PlayerTwo:   "ccc",
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
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
