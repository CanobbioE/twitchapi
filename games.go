package gwat

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Game represents a game as described by the twitch API documentation.
type Game struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BoxArtURL string `json:"box_art_url"`
}

// games represents an array of Game
type games struct {
	Data []Game `json:data`
}

// GetGames gets information about one or more games specified by id and/or name.
// At least one between ids and names must be specified.
func (c *Client) GetGames(ids []string, names []string) ([]Game, error) {
	uri := BaseURL + GamesEP

	// checking input
	if ids == nil && names == nil {
		return nil, errors.New("At least one id or name must be specified")
	}
	if len(ids) > 10 || len(names) > 10 {
		return nil, errors.New("A maximum of 10 ids or names can be specified")
	}

	// creating uri and header
	if ids != nil {
		uri += "?id="
		addParameters(&uri, "id", ids)
	}
	if names != nil {
		uri += "?name="
		addParameters(&uri, "name", names)
	}
	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	// performing request
	res, err := c.request("GET", uri, h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// parsing result
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	retGames := games{}
	json.Unmarshal(body, &retGames)
	return retGames.Data, nil
}

// TODO: use net/url library
func addParameters(uri *string, paramName string, values []string) {
	for _, val := range values {
		*uri += val
		*uri += "&" + paramName + "="
	}
}
