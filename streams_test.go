package twitchapi

import (
	"fmt"
	"testing"
)

func TestGetStreams(t *testing.T) {
	var tests = []struct {
		input   StreamQueryParameters
		wantErr bool
	}{
		{
			StreamQueryParameters{
				First:    20,
				Language: []string{"en"},
			},
			false,
		},
		{
			StreamQueryParameters{
				First: 200,
			},
			true,
		},
		{
			StreamQueryParameters{
				Language: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "90", "91", "92", "93", "94", "95", "96", "97", "98", "99", "100"},
			},
			true,
		},
	}

	for _, test := range tests {
		_, _, err := c.GetStreams(test.input)
		if err != nil && !test.wantErr {
			s := fmt.Sprintf("GetStreams returned error: %v", err)
			t.Errorf(s)
			return
		}
	}
}

func TestGetStreamsMetadata(t *testing.T) {
	var tests = []struct {
		input   StreamQueryParameters
		wantErr bool
	}{
		{StreamQueryParameters{}, false},
		{StreamQueryParameters{First: 200}, true},
	}

	for _, test := range tests {
		_, _, err := c.GetStreamsMetadata(test.input)
		if err != nil && !test.wantErr {
			s := fmt.Sprintf("GetStreamsMetadata returned error: %v", err)
			t.Errorf(s)
			return
		}
	}
}
