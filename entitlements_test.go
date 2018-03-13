package twitchapi

import "testing"

func TestCreateEntitlementGrantsUploadURL(t *testing.T) {
	var tests = []struct {
		input   EntitlementURLQueryParameters
		authTkn string
		wantErr bool
	}{
		{
			EntitlementURLQueryParameters{
				ManifestID: "123456789010",
				Type:       "bulk_drops_grant",
			},
			"cfabdegwdoklmawdzdo98xt2fo512y",
			true,
		},
		{
			EntitlementURLQueryParameters{
				ManifestID: "123456789010",
				Type:       "fake_type",
			},
			"cfabdegwdoklmawdzdo98xt2fo512y",
			true,
		},
		{
			EntitlementURLQueryParameters{
				ManifestID: "123456789010",
				Type:       "bulk_drops_grant",
			},
			"",
			true,
		},
		{
			EntitlementURLQueryParameters{
				ManifestID: "",
				Type:       "bulk_drops_grant",
			},
			"",
			true,
		},
	}

	for _, test := range tests {
		_, err := c.CreateEntitlementGrantsUploadURL(test.input, test.authTkn)
		if err != nil && !test.wantErr {
			t.Errorf("%v", err)
		}
	}
}
