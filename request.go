package gwat

import "net/http"

// Header represent a simplified http.Header.
type Header struct {
	Field string // field name
	Value string // field value
}

// Request performs a http request and returns the response.
func (c *Client) request(method, uri string, h Header) (*http.Response, error) {
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
