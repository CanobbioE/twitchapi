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

	// cheking required fields
	if isEmpty(ids) && isEmpty(names) {
		return nil, errors.New("At least one id or name must be specified")
	}
	if (!isEmpty(ids)) && len(ids) > 100 {

		return nil, errors.New("GetGames: A maximum of 100 ids can be specified")
	}
	if (!isEmpty(names)) && len(names) > 100 {
		return nil, errors.New("GetGames: A maximum of 100 names can be specified")
	}

	// create the header
	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	// perform API call
	uri := makeUri(BaseURL+GamesEP, qp)
	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// parse result
	if res.Status != "200 OK" {
		return nil, errors.New("GetGames returned status " + res.Status)
	}

	if err := parseResult(res, &retGames); err != nil {
		return nil, err
	}
	return retGames.Data, nil
}

// GetTopGames gets games sorted by number of current viewers.
func (c *Client) GetTopGames(qp TopGameQueryParameters) ([]Game, error) {
	retGames := gameData{}

	// check for parameters boundries, set default min value
	if qp.First > 100 {
		return nil, errors.New("GetTopGames: \"First\" parameter cannot be greater than 100")
	}
	if qp.First <= 0 {
		qp.First = 20
	}

	// create the header
	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	// perform API call
	uri := makeUri(BaseURL+GamesEP+TopGamesEP, qp)
	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// parse the result
	if res.Status != "200 OK" {
		return nil, errors.New("GetTopGames returned status " + res.Status)
	}

	if err := parseResult(res, &retGames); err != nil {
		return nil, err
	}

	return retGames.Data, nil
}

// GetGameAnalytics gets an URL that game developers can use to download
// analytics reports for their games. The URL is valid for one minute.
// An authorization token is required with scope "analytics:read:games"
func (c *Client) GetGameAnalytics(id string, authTkn string) ([]Analytic, error) {
	uri := BaseURL + AnalyticsEP + GamesEP

	// checking required parameters
	if !isEmpty(id) {
		uri += "?id=" + id
	}

	// create the header
	h := Header{}
	if !isEmpty(authTkn) {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	} else {
		return nil, errors.New("GetGameAnalytics: An authorization token is needed")
	}

	// perform API call
	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// parse the result
	if res.Status != "200 OK" {
		return nil, errors.New("GetGameAnalytics returned status: " + res.Status)
	}

	retAnalytics := analyticsData{}
	if err := parseResult(res, &retAnalytics); err != nil {
		return nil, err
	}

	return retAnalytics.Data, nil
}
