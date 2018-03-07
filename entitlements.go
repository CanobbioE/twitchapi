package gwat

import "errors"

// uploadURL represent the response returned by CreateEntitlementGrantsUploadURL.
type uploadURL struct {
	url string `json:"url"`
}

// CreateEntitlementGrantsUploadURL creates a URL where you can upload a manifest file granting entitlement to users.
// manifestID is the unique identifier of the manifest file to be uploaded. Must be 1-64 characters.
// entitleType is the type of entitlement granted. Only "bulk_drops_grant" is supported.
func (c *Client) CreateEntitlementGrantsUploadURL(manifestID, entitleType, authTkn string) (string, error) {
	uri := BaseURL + EntitlementsEP + UploadEP
	retUploadURL := uploadURL{}

	ml := len(manifestID)
	if ml > 64 || ml < 1 {
		return "", errors.New("Manifest ID's length must be between 1 and 64")
	}

	if entitleType != "bulk_drop_grant" {
		return "", errors.New("Only \"bulk_drop_grant\" supported as entitle type")
	}

	h := Header{
		Field: "Authorization:",
		Value: "Bearer " + authTkn,
	}

	res, err := c.request("POST", uri, h)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if err := parseResult(res, &retUploadURL); err != nil {
		return "", err
	}

	return retUploadURL.url, nil
}
