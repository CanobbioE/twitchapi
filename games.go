package twitchapi

import (
	"errors"
)

// GetGames gets information about one or more games specified by id and/or name.
// At least one between ids and names must be specified.
func (c *Client) GetGames(qp GameQueryParameters) ([]Game, error) {
	retGames := gameData{}

	ids := qp.IDs
	names := qp.Names

	if isNil(ids) && isNil(names) {
		return nil, errors.New("At least one id or name must be specified")
	}
	if (!isNil(ids)) && len(ids) > 100 {

		return nil, errors.New("GetGames: A maximum of 100 ids can be specified")
	}
	if (!isNil(names)) && len(names) > 100 {
		return nil, errors.New("GetGames: A maximum of 100 names can be specified")
	}

	uri := makeUri(GamesEP, qp)
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
	retGames := gameData{}

	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}
	uri := makeUri(GamesEP+TopGamesEP, qp)

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
