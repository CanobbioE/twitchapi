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

	if len(qp.IDs) > 100 {
		qp.IDs = qp.IDs[:99]
	}
	if len(qp.Logins) > 100 {
		qp.Logins = qp.Logins[:99]
	}

	h := Header{}
	if !isEmpty(authTkn) {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	} else {
		return nil, errors.New("GetUsers: an authorization token is required")
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

	qp.First = setDefaultValueIf(qp.First > 100, qp.First, 100).(int)
	qp.First = setDefaultValueIf(qp.First <= 0, qp.First, 1).(int)

	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	uri := makeUri(BaseURL+UsersEP+FollowsEP, qp)
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
	if !isEmpty(authTkn) {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	}

	// create the uri
	uri, err := url.Parse(BaseURL + UsersEP)
	if err != nil {
		return nil, err
	}
	v := url.Values{}
	v.Add("description", description)
	uri.RawQuery = v.Encode()

	// perform api call
	res, err := c.apiCall("PUT", uri.String(), h)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// parse result
	if res.Status != "200 OK" {
		return nil, errors.New("UpdateUser returned status " + res.Status)
	}
	if err := parseResult(res, &retUsers); err != nil {
		return nil, err
	}

	return retUsers.Data, nil
}
