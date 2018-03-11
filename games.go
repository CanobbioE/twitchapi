package twitchapi

import (
	"errors"
	"strconv"
)

// GetGames gets information about one or more games specified by id and/or name.
// At least one between ids and names must be specified.
func (c *Client) GetGames(ids []string, names []string) ([]Game, error) {
	uri := BaseURL + GamesEP
	retGames := gameData{}

	if ids == nil && names == nil {
		return nil, errors.New("At least one id or name must be specified")
	}
	if len(ids) > 10 || len(names) > 10 {
		return nil, errors.New("A maximum of 10 ids or names can be specified")
	}

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

	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := parseResult(res, &retGames); err != nil {
		return nil, err
	}
	return retGames.Data, nil
}

// GetTopGames gets games sorted by number of current viewers.
func (c *Client) GetTopGames(after, before string, first int) ([]Game, error) {
	uri := BaseURL + GamesEP + TopGamesEP
	retGames := gameData{}

	uri += "?"
	if !isNil(after) {
		uri += "after=" + after + "&"
	}
	if !isNil(before) {
		uri += "before=" + before + "&"
	}
	if !isNil(first) {
		uri += "first=" + strconv.Itoa(first) + "&"
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

	if err := parseResult(res, &retGames); err != nil {
		return nil, err
	}

	return retGames.Data, nil
}
