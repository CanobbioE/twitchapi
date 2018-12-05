package twitchapi

import "errors"

// CreateEntitlementGrantsUploadURL creates a URL where you can upload a
// manifest file granting entitlement to users.
// It requires an application access token (authTkn)
func (c *Client) CreateEntitlementGrantsUploadURL(qp EntitlementURLQueryParameters, authTkn string) (string, error) {

	// check parameters boundries
	ml := len(qp.ManifestID)
	if ml > 64 || ml < 1 {
		return "", errors.New("Manifest ID's length must be between 1 and 64")
	}

	err := isValid("type", qp.Type, []string{"bulk_drops_grant"})
	if err != nil {
		return "", errors.New("Only \"bulk_drops_grant\" supported as entitle type")
	}

	// create the header
	h := Header{}
	if !isEmpty(authTkn) {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	} else {
		return "", errors.New("CreateEntitlementGrantsUploadURL: An authorization token is needed")
	}

	// perform API call
	uri := makeUri(BaseURL+EntitlementsEP+UploadEP, qp)
	res, err := c.apiCall("POST", uri, h)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// parse the result
	if res.Status != "200 OK" {
		return "", errors.New("CreateEntitlementGrantsUploadURL returned status:" + res.Status)
	}

	retUploadURL := uploadData{}
	if err := parseResult(res, &retUploadURL); err != nil {
		return "", err
	}

	return retUploadURL.Data[0].URL, nil
}
