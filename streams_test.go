package twitchapi

import (
	"fmt"
	"testing"
)

func TestGetStreams(t *testing.T) {
	var tests = []struct {
		input StreamQueryParameters
	}{
		{
			StreamQueryParameters{
				First: 20,
			},
		},
	}

	for _, test := range tests {
		got, _, err := c.GetStreams(test.input)
		fmt.Println(got[0].ID)
		if err != nil {
			s := fmt.Sprintf("GetStreams returned error: %v", err)
			t.Errorf(s)
			return
		}
	}
}

func TestGetStreamsMetadata(t *testing.T) {
	var tests = []struct {
		input StreamQueryParameters
	}{
		{StreamQueryParameters{}},
	}

	for _, test := range tests {
		got, _, err := c.GetStreamsMetadata(test.input)
		fmt.Println(got[0].GameID, err)
		if err != nil {
			s := fmt.Sprintf("GetStreamsMetadata returned error: %v", err)
			t.Errorf(s)
			return
		}
	}
}
