package gwat

import "net/http"

type Client struct {
	ClientID   string
	HttpClient *http.Client
}

func NewClient(clientID string) *Client {
	return &Client{
		ClientID:   clientID,
		HttpClient: &http.Client{},
	}
}
