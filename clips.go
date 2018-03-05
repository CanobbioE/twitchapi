package gwat

import (
	"encoding/json"
	"io/ioutil"
)

// Clip represent a clip as described by the twitch API documentation
type Clip struct {
	ID            string `json:"id"`
	URL           string `json:"url"`
	EmbedUrl      string `json:"embed_url"`
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

// CreateClip creates a clip and returns id and edit URL for the new clip.
// CreateClip requires an authentication token (authTkn) with scope 'clips:edit'
func (c *Client) CreateClip(broadcasterID string) {
}

// GetClip gets information abput a clip specified by id (optional).
func (c *Client) GetClip(id string) Clip {
	uri := BaseURL + ClipsEP

	if id != nil {
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	retClip := Clip{}
	json.Unmarshal(body, &retClip)

	return retClip, nil
}
