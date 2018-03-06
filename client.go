package gwat

import "net/http"

type Client struct {
	ClientID   string
	httpClient *http.Client
}

func NewClient(clientID string) *Client {
	return &Client{
		ClientID:   clientID,
		httpClient: &http.Client{},
	}
}
