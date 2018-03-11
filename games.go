package twitchapi

import (
	"errors"
	"fmt"
)

// GetGames gets information about one or more games specified by id and/or name.
// At least one between ids and names must be specified.
func (c *Client) GetGames(qp GameQueryParameters) ([]Game, error) {
	uri := BaseURL + GamesEP
	retGames := gameData{}

	params := parseInput(qp)

	ids, idIsOk := params["id"]
	names, nameIsOk := params["name"]

	if !idIsOk && !nameIsOk {
		return nil, errors.New("At least one id or name must be specified")
	}

	if len(ids.([]string)) > 10 || len(names.([]string)) > 10 {
		return nil, errors.New("A maximum of 10 ids or names can be specified")
	}

	uri += "?"
	addParameters(&uri, "id", ids.([]string))
	addParameters(&uri, "name", names.([]string))
	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		return nil, errors.New("GetGames returned status" + res.Status)
	}

	if err := parseResult(res, &retGames); err != nil {
		return nil, err
	}
	return retGames.Data, nil
}

// GetTopGames gets games sorted by number of current viewers.
func (c *Client) GetTopGames(qp TopGameQueryParameters) ([]Game, error) {
	uri := BaseURL + GamesEP + TopGamesEP
	retGames := gameData{}

	params := parseInput(qp)

	uri += "?"
	for k, v := range params {
		uri += fmt.Sprintf("%s=%v&", k, v)
	}

	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		return nil, errors.New("GetTopGames returned status" + res.Status)
	}

	if err := parseResult(res, &retGames); err != nil {
		return nil, err
	}

	return retGames.Data, nil
}
