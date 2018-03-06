package gwat

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Clip represent a clip as described by the twitch API documentation.
type Clip struct {
	ID            string `json:"id"`
	URL           string `json:"url"`
	EmbedURL      string `json:"embed_url"`
	BroadcasterID string `json:"broadcaster_id"`
	CreatorID     string `json:"creator_id"`
	VideoID       string `json:"video_id"`
	GameID        string `json:"game_id"`
	Language      string `json:"language"`
	Title         string `json:"title"`
	ViewCount     int    `json:"view_count"`
	CreatedAt     string `json:"created_at"`
	ThumbnailURL  string `json:"thumbnail_url"`
}

// DataClip represent the response generated by CreateClip as described by the twitch API documentation.
type DataClip struct {
	ID      string `json:"id"`
	editURL string `json:"edit_url"`
}

// CreateClip creates a clip and returns the information associated to the new clip.
// CreateClip requires an authentication token (authTkn) with scope 'clips:edit
// and the id of the stream from which the clip will be made (broadcasterID).
func (c *Client) CreateClip(broadcasterID, authTkn string) (DataClip, error) {
	retDataClip := DataClip{}
	uri := BaseURL + ClipsEP

	if broadcasterID != "" {
		uri += "?broadcaster_id=" + broadcasterID
	} else {
		return retDataClip, errors.New("broadcasterID must be specified")
	}

	h := Header{}
	if authTkn != "" {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	} else {
		return retDataClip, errors.New("An authorization token is needed")
	}

	res, err := c.request("POST", uri, h)
	if err != nil {
		return retDataClip, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return retDataClip, err
	}

	json.Unmarshal(body, &retDataClip)
	return retDataClip, nil
}

// GetClip gets information about a clip specified by an optional id.
func (c *Client) GetClip(id string) (Clip, error) {
	retClip := Clip{}
	uri := BaseURL + ClipsEP

	if id != "" {
		uri += "?id=" + id
	}

	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	res, err := c.request("GET", uri, h)
	if err != nil {
		return retClip, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return retClip, err
	}
	json.Unmarshal(body, &retClip)

	return retClip, nil
}
