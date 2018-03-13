package twitchapi

import (
	"testing"
)

func TestGetUsers(t *testing.T) {
	var tests = []struct {
		input   UserQueryParameters
		authTkn string
		wantErr bool
		want    string
	}{
		{
			UserQueryParameters{},
			"",
			true,
			"",
		},
		{
			UserQueryParameters{
				IDs: []string{"44322889"},
			},
			"cfabdegwdoklmawdzdo98xt2fo512y",
			true,
			"dallas",
		},
	}

	for _, test := range tests {
		got, err := c.GetUsers(test.input, test.authTkn)
		if err != nil {
			if !test.wantErr {
				t.Errorf("%v", err)
			}
		} else {
			if got[0].DisplayName != test.want {
				t.Errorf("GetUsers() = %s (expected %s)", got[0].DisplayName, test.want)
			}
		}
	}
}

func TestGetUsersFollow(t *testing.T) {
	var tests = []struct {
		input   FollowQueryParameters
		wantErr bool
	}{
		{
			FollowQueryParameters{
				ToID: "23161357",
			},
			false,
		},
		{
			FollowQueryParameters{},
			true,
		},
	}

	for _, test := range tests {
		_, _, err := c.GetUserFollows(test.input)
		if err != nil {
			if !test.wantErr {
				t.Errorf("%v", err)
			}
		}
	}
}

func TestUpdateUser(t *testing.T) {
	var tests = []struct {
		input   string
		authTkn string
		wantErr bool
	}{
		{
			"Description",
			"123123",
			true,
		},
	}

	for _, test := range tests {
		_, err := c.UpdateUser(test.input, test.authTkn)
		if err != nil {
			if !test.wantErr {
				t.Errorf("%v", err)
			}
		}
	}

}
