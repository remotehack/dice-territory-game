package game_test

import (
	"reflect"
	"testing"

	"javorszky/dice-territory-game/v2/pkg/game"
)

func TestNewPlayer(t *testing.T) {
	type args struct {
		ID   string
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    game.Player
		wantErr bool
	}{
		{
			name: "can create a player with a name",
			args: args{
				ID:   "playerid-1",
				name: "John",
			},
			want: game.Player{
				ID:    "playerid-1",
				Name:  "John",
				Score: 0,
			},
			wantErr: false,
		},
		{
			name: "can not create a player with an empty name",
			args: args{
				ID:   "playerid-1",
				name: "",
			},
			want:    game.Player{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := game.NewPlayer(tt.args.ID, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
