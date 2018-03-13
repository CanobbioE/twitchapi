package twitchapi

import (
	"testing"
)

func TestCreateClip(t *testing.T) {
	var tests = []struct {
		input   string
		authTkn string
		wantErr bool
	}{
		{
			"44322889",
			"cfabdegwdoklmawdzdo98xt2fo512y",
			true,
		},
		{
			"44322889",
			"",
			true,
		},
		{
			"",
			"",
			true,
		},
	}

	for _, test := range tests {
		_, err := c.CreateClip(test.input, test.authTkn)
		if err != nil && !test.wantErr {
			t.Errorf("%v", err)
			return
		}
	}
}

func TestGetClip(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{
			"AwkwardHelplessSalamanderSwiftRage",
			"babymetal",
		},
	}
	for _, test := range tests {
		got, err := c.GetClip(test.input)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		if got[0].Title != test.want {
			t.Errorf("GetClip(%s) = %s (expected %s)", test.input, got[0].Title, test.want)
		}
	}
}
