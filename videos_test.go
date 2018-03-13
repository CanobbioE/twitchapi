package twitchapi

import "testing"

func TestGetVideos(t *testing.T) {
	var tests = []struct {
		input   VideoQueryParameters
		wantErr bool
	}{
		{
			VideoQueryParameters{
				GameID: "493057",
			},
			false,
		},
		{
			VideoQueryParameters{
				ID: "172982667",
			},
			false,
		},
		{
			VideoQueryParameters{
				First: 200,
			},
			true,
		},
	}

	for _, test := range tests {
		_, _, err := c.GetVideos(test.input)
		if err != nil && !test.wantErr {
			t.Errorf("%v", err)
		}
	}
}
