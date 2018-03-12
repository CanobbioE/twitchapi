package twitchapi

import (
	"testing"
)

var c = NewClient("gxcprgf6w76j8wz42g327j4b98411i")

func TestApiCall(t *testing.T) {
	var tests = []struct {
		method string
		uri    string
		header Header
		want   string
	}{
		{
			"GET",
			"https://api.twitch.tv/helix/clips?id=AwkwardHelplessSalamanderSwiftRage",
			Header{
				Field: "Client-ID",
				Value: "gxcprgf6w76j8wz42g327j4b98411i",
			},
			"67955580",
		},
	}

	for _, test := range tests {
		got, err := c.apiCall(test.method, test.uri, test.header)
		if err != nil {
			t.Errorf("apiCall(%s) returned error: %v", test.uri, err)
		}
		defer got.Body.Close()
		ret := clipData{}
		if err := parseResult(got, &ret); err != nil {
			t.Errorf("parseResult() returned error: %v", err)
		}
		if ret.Data[0].BroadcasterID != test.want {
			t.Errorf("apiCall() returned %v", ret.Data[0])
		}
	}
}
