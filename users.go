package twitchapi

import (
	"errors"
	"net/url"
)

// GetUsers gets information about one or more specified Twitch users,
// identified by id or login.
// The authentication token must have scope 'user:read:mail'
func (c *Client) GetUsers(qp UserQueryParameters, authTkn string) ([]User, error) {
	retUsers := userData{}

	ids := qp.IDs
	logins := qp.Logins

	if len(ids) > 100 {
		return nil, errors.New("GetUsers: A maximum of 100 ids can be specified")
	}
	if len(logins) > 100 {
		return nil, errors.New("GetUsers: A maximum of 100 logins can be specified")
	}

	h := Header{}
	if !isNil(authTkn) {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	} else {
		return nil, errors.New("GetUsers: An authorization token is needed")
	}
	uri := makeUri(BaseURL+UsersEP, qp)

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

// GetUsersFollows gets information on follow relationships between two Twitch
// users. At minimum, from id or to_id must be porvided.
// It also returns the number of items returned.
//
// - If only from_id was in the request, returns the total number of followed users.
//
// - If only to_id was in the request, returns the total number of followers.
//
// - If both from_id and to_id were in the request, returns 1 (if the "from" user follows the "to" user) or 0.
func (c *Client) GetUserFollows(qp FollowQueryParameters) ([]UserFollows, int, error) {
	retUsersFollows := userFollowData{}
	var retTotal int

	if qp.First > 100 {
		return nil, retTotal, errors.New("GetUsersFollows: \"First\" parameter cannot be greater than 100")
	}

	uri := makeUri(BaseURL+UsersEP+FollowsEP, qp)
	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, retTotal, err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		return nil, retTotal, errors.New("GetUsersFollows returned status " + res.Status)
	}

	if err := parseResult(res, &retUsersFollows); err != nil {
		return nil, retTotal, err
	}
	retTotal = retUsersFollows.Total
	return retUsersFollows.Data, retTotal, nil
}

// UpdateUser updates the description of a user specified by the authentication
// token (authTkn).
// The authentication token must have scope user:edit
func (c *Client) UpdateUser(description, authTkn string) ([]User, error) {
	retUsers := userData{}

	h := Header{}
	if !isNil(authTkn) {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	}

	uri, err := url.Parse(BaseURL + UsersEP)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("description", description)
	uri.RawQuery = v.Encode()

	res, err := c.apiCall("PUT", uri.String(), h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		return nil, errors.New("UpdateUser returned status " + res.Status)
	}

	if err := parseResult(res, &retUsers); err != nil {
		return nil, err
	}

	return retUsers.Data, nil
}
