package twitchapi

import (
	"errors"
)

// GetBitsLeaderboard gets a ranked list of Bits leaderboard information.
// An authorization token is required with scope "bits:read"
func (c *Client) GetBitsLeaderboard(qp BitsQueryParameters, authTkn string) ([]Leaderboard, DateRange, int, error) {
	retBits := bitsLeaderboardData{}

	qp.Count = setDefaultValueIf(qp.Count <= 0, &qp.Count, 10).(int)
	qp.Count = setDefaultValueIf(qp.Count > 100, &qp.Count, 100).(int)

	// checking for required fields
	if err := checkRequiredFields("GetBitsLeaderboard", qp.Period); err != nil {
		qp.Period = "all"
	}
	valid := []string{"all", "day", "week", "month", "year"}
	if err := isValid("period", qp.Period, valid); err != nil {
		return []Leaderboard{}, DateRange{}, -1, err
	}
	if qp.Period == "all" {
		qp.StartedAt = ""
	}

	// creating the header
	h := Header{}
	if !isEmpty(authTkn) {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	} else {
		err := errors.New("GetBitsLeaderboard: an authorization token is needed")
		return []Leaderboard{}, DateRange{}, -1, err
	}

	// perform API call
	uri := makeUri(BaseURL+BitsEP+LeaderboardEP, qp)
	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return []Leaderboard{}, DateRange{}, -1, err
	}
	defer res.Body.Close()

	// parse the response
	if err := parseResult(res, &retBits); err != nil {
		return []Leaderboard{}, DateRange{}, -1, err
	}
	if res.Status != "200 OK" {
		err := errors.New("GetBitsLeaderboard returned status: " + res.Status)
		return []Leaderboard{}, DateRange{}, -1, err
	}

	return retBits.Data, retBits.DateRange, retBits.Total, nil
}
