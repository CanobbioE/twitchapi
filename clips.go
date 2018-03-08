package gwat

import (
	"errors"
)

// CreateClip creates a clip and returns the information associated to the new clip.
// CreateClip requires an authentication token (authTkn) with scope 'clips:edit
// and the id of the stream from which the clip will be made (broadcasterID).
func (c *Client) CreateClip(broadcasterID, authTkn string) ([]ClipInfo, error) {
	uri := BaseURL + ClipsEP

	if !isNil(broadcasterID) {
		uri += "?broadcaster_id=" + broadcasterID
	} else {
		return nil, errors.New("broadcasterID must be specified")
	}

	h := Header{}
	if !isNil(authTkn) {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	} else {
		return nil, errors.New("An authorization token is needed")
	}

	res, err := c.request("POST", uri, h)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	retClipInfo := clipInfoData{}
	if err := parseResult(res, &retClipInfo); err != nil {
		return nil, err
	}
	return retClipInfo.Data, nil
}

// GetClip gets information about a clip specified by an optional id.
func (c *Client) GetClip(id string) ([]Clip, error) {
	uri := BaseURL + ClipsEP

	if !isNil(id) {
		uri += "?id=" + id
	}

	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	res, err := c.request("GET", uri, h)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	retClip := clipData{}
	if err := parseResult(res, &retClip); err != nil {
		return nil, err
	}

	return retClip.Data, nil
}
