package twitchapi

import (
	"errors"
)

// GetBitsLeaderboard gets a ranked list of Bits leaderboard information.
// An authorization token is required with scope "bits:read"
func (c *Client) GetBitsLeaderboard(qp BitsQueryParameters, authTkn string) ([]Leaderboard, DateRange, int, error) {
	retBits := bitsLeaderboardData{}

	if qp.Count < 0 {
		qp.Count = 10
	}
	if qp.Count > 100 {
		qp.Count = 100
	}

	if isNil(qp.Period) {
		qp.Period = "all"
	} else {
		valid := []string{"all", "day", "week", "month", "year"}
		err := isValid("period", qp.Period, valid)
		if err != nil {
			return []Leaderboard{}, DateRange{}, -1, err
		}
		if qp.Period == "all" {
			qp.StartedAt = ""
		}
	}

	h := Header{}
	if !isNil(authTkn) {
		h.Field = "Authorization"
		h.Value = "Bearer " + authTkn
	} else {
		err := errors.New("GetBitsLeaderboard: An authorization token is needed")
		return []Leaderboard{}, DateRange{}, -1, err
	}

	uri := makeUri(BaseURL+BitsEP+LeaderboardEP, qp)
	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return []Leaderboard{}, DateRange{}, -1, err
	}
	defer res.Body.Close()

	if err := parseResult(res, &retBits); err != nil {
		return []Leaderboard{}, DateRange{}, -1, err
	}
	if res.Status != "200 OK" {
		err := errors.New("GetBitsLeaderboard returned status: " + res.Status)
		return []Leaderboard{}, DateRange{}, -1, err
	}

	return retBits.Data, retBits.DateRange, retBits.Total, nil
}
