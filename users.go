package twitchapi

import (
	"errors"
	"fmt"
)

// GetUsers gets information about one or more specified Twitch users, identified by id or login.
// The authentication token must have scope 'user:read:mail'
func (c *Client) GetUsers(qp UserQueryParameters, authTkn string) ([]User, error) {
	uri := BaseURL + UsersEP
	retUsers := userData{}

	params := parseInput(qp)

	ids := params["id"]
	logins := params["login"]

	if len(ids.([]string)) > 100 {
		return nil, errors.New("GetUsers: A maximum of 100 ids can be specified")
	}
	if len(logins.([]string)) > 100 {
		return nil, errors.New("GetUsers: A maximum of 100 logins can be specified")
	}

	uri += "?"
	addParameters(&uri, "id", ids.([]string))
	addParameters(&uri, "login", logins.([]string))
	h := Header{}
	if !isNil(authTkn) {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	} else {
		return nil, errors.New("GetUsers: An authorization token is needed")
	}

	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		return nil, errors.New("GetUsers returned status" + res.Status)
	}

	if err := parseResult(res, &retUsers); err != nil {
		return nil, err
	}

	return retUsers.Data, nil
}

// GetUsersFollows gets information on follow relationships between two Twitch users.
func (c *Client) GetUserFollows(qp FollowQueryParameters) ([]UserFollows, error) {
	uri := BaseURL + UsersEP + FollowsEP
	retUsersFollows := userFollowData{}

	params := parseInput(qp)
	uri += "?"
	if params["first"].(int) > 100 {
		return nil, errors.New("GetUsersFollows: \"First\" parameter cannot be greater than 100")
	}

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
		return nil, errors.New("GetUsersFollows returned status" + res.Status)
	}

	if err := parseResult(res, &retUsersFollows); err != nil {
		return nil, err
	}
	return retUsersFollows.Data, nil
}

// UpdateUser updates the description of a user specified by the authentication token (authTkn).
// The authentication token must have scope user:edit
func (c *Client) UpdateUser(description, authTkn string) ([]User, error) {
	uri := BaseURL + UsersEP
	retUsers := userData{}

	uri += "?description=" + description
	h := Header{}
	if !isNil(authTkn) {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	}

	res, err := c.apiCall("PUT", uri, h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		return nil, errors.New("UpdateUser returned status" + res.Status)
	}

	if err := parseResult(res, &retUsers); err != nil {
		return nil, err
	}

	return retUsers.Data, nil
}
