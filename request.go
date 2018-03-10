package twitchapi

import (
	"encoding/json"
	"errors"
	"fmt"
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

// streamRequest prepares a stream request whether it's for metadata or not
func (c *Client) streamRequest(uri *string, qp StreamQueryParameters) (*http.Response, error) {
	params := parseInput(qp)

	if params["first"].(int) > 100 {
		err := errors.New("\"First\" parameter cannot be greater than 100")
		return &http.Response{}, err
	}

	*uri += "?"
	for k, v := range params {
		if k == "comunity_id" {
			addParameters(uri, "comunity_id", v.([]string))
		} else {
			*uri += fmt.Sprintf("%s=%v&", k, v)
		}
	}

	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	res, err := c.apiCall("GET", *uri, h)
	return res, err
}
