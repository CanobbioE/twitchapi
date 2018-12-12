package twitchapi

import (
	"errors"
)

// CreateClip creates a clip and returns the information associated to the new clip.
// CreateClip requires an authentication token (authTkn) with scope 'clips:edit
// and the id of the stream from which the clip will be made (broadcasterID).
func (c *Client) CreateClip(broadcasterID, authTkn string) ([]ClipInfo, error) {
	retClipInfo := clipInfoData{}
	uri := BaseURL + ClipsEP

	if err := checkRequiredFields("CreateClip", "all", broadcasterID, authTkn); err != nil {
		return nil, err
	}
	uri += "?broadcaster_id=" + broadcasterID

	// creating the header
	h := Header{
		Field: "Authorization",
		Value: "Bearer " + authTkn,
	}

	// perform API call
	res, err := c.apiCall("POST", uri, h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// parse the response
	if res.Status != "200 OK" {
		return nil, errors.New("CreateClip returned status: " + res.Status)
	}
	if err := parseResult(res, &retClipInfo); err != nil {
		return nil, err
	}

	return retClipInfo.Data, nil
}

// GetClip gets information about one or more clips specified by an id (broadcaster, game, clip).
func (c *Client) GetClip(qp ClipQueryParameter) ([]Clip, Cursor, error) {
	retClip := clipData{}

	// checking required fields
	err := checkRequiredFields("GetClip", "any", qp.BroadcasterID, qp.GameID, qp.ID)
	if err != nil {
		e := errors.New("GetClip: at least one id must be specified")
		return []Clip{}, Cursor{}, e
	}

	if !isEmpty(qp.First) && (qp.First > 100 || qp.First < 0) {
		err := errors.New("GetClip: first parameter must be between 0 and 100")
		return []Clip{}, Cursor{}, err
	}

	// creating the header
	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	// perform API call
	uri := makeUri(BaseURL+ClipsEP, qp)
	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return []Clip{}, Cursor{}, err
	}
	defer res.Body.Close()

	// parse result
	if res.Status != "200 OK" {
		err := errors.New("CreateClip returned status: " + res.Status)
		return []Clip{}, Cursor{}, err
	}
	if err := parseResult(res, &retClip); err != nil {
		return []Clip{}, Cursor{}, err
	}

	return retClip.Data, retClip.Cursor, nil
}
