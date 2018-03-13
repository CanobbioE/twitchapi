package twitchapi

import (
	"fmt"
	"testing"
)

func TestGetGamesById(t *testing.T) {
	var tests = []struct {
		input GameQueryParameters
		want  string
	}{
		{
			GameQueryParameters{
				IDs: []string{"493057"},
			},
			"PLAYERUNKNOWN'S BATTLEGROUNDS",
		},
	}

	for _, test := range tests {
		got, err := c.GetGames(test.input)
		if err != nil {
			s := fmt.Sprintf("%v", err)
			t.Errorf(s)
			return
		}
		if got[0].Name != test.want {
			t.Errorf("GetGames(%v) = %s (expected %s)", test.input.IDs, got[0].Name)
		}
	}
}

func TestGetGamesByName(t *testing.T) {
	var tests = []struct {
		input GameQueryParameters
		want  string
	}{
		{
			GameQueryParameters{
				Names: []string{"PLAYERUNKNOWN'S BATTLEGROUNDS"},
			},
			"493057",
		},
	}

	for _, test := range tests {
		got, err := c.GetGames(test.input)
		if err != nil {
			s := fmt.Sprintf("GetGames(%v) = %v", test.input.Names, err)
			t.Errorf(s)
			return
		}
		if got[0].ID != test.want {
			t.Errorf("GetGames(%v) = %s (expected %s)", test.input.Names, got[0].ID)
		}
	}
}

func TestTopGames(t *testing.T) {
	var tests = []struct {
		input   TopGameQueryParameters
		wantErr bool
	}{
		{
			TopGameQueryParameters{
				First: 20,
			},
			false,
		},
		{
			TopGameQueryParameters{
				First: 200,
			},
			true,
		},
		{
			TopGameQueryParameters{},
			false,
		},
	}
	for _, test := range tests {
		_, err := c.GetTopGames(test.input)
		if err != nil && !test.wantErr {
			t.Errorf("%v", err)
		}
	}
}
