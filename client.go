package twitchapi

import "net/http"

// Client represents a Twitch client used to perform API calls.
type Client struct {
	ClientID   string
	HttpClient *http.Client
}

// NewClient returns a newly created TwitchClient with the specified clientID.
func NewClient(clientID string) *Client {
	return &Client{
		ClientID:   clientID,
		HttpClient: &http.Client{},
	}
}
