package twitchapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Header represent a simplified http.Header.
type Header struct {
	Field string // field name
	Value string // field value
}

// apiCall performs a http request and returns the response.
func (c *Client) apiCall(method, uri string, h Header) (*http.Response, error) {
	client := c.HttpClient

	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(h.Field, h.Value)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// parseResult reads a json from src into dst
func parseResult(src *http.Response, dst interface{}) error {
	body, err := ioutil.ReadAll(src.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(body, dst)
	return nil
}
