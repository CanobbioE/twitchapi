package twitchapi

import (
	"fmt"
	"net/url"
)

// GetVideos gets video information by video ID (one or more), user ID (one only), or game ID (one only).
// If id is specified any other parameter is ignored.
// For lookup by user or game ID a Cursor is returned.
func (c *Client) GetVideos(qp VideoQueryParameters) ([]Video, Cursor, error) {
	retVideos := videoData{}
	retCursor := Cursor{}
	var uri string

	if isEmpty(qp.ID) {

		// default values
		qp.Period = setDefaultValueIf(isEmpty(qp.Period), qp.Period, "all").(string)
		qp.Sort = setDefaultValueIf(isEmpty(qp.Sort), qp.Sort, "time").(string)
		qp.Type = setDefaultValueIf(isEmpty(qp.Type), qp.Type, "all").(string)
		qp.First = setDefaultValueIf(qp.First > 100, qp.First, 100).(int)
		qp.First = setDefaultValueIf(qp.First <= 0, qp.First, 20).(int)

		// check parameters
		if err := isValid("period", qp.Period, []string{"all", "day", "month", "week"}); err != nil {
			return nil, retCursor, err
		}
		if err := isValid("sort", qp.Sort, []string{"time", "trending", "views"}); err != nil {
			return nil, retCursor, err
		}
		if err := isValid("type", qp.Type, []string{"all", "upload", "archive", "highlight"}); err != nil {
			return nil, retCursor, err
		}
		uri = makeUri(BaseURL+VideosEP, qp)

	} else {
		fmt.Println("GetVideos: \"id\" was specified. Ignoring all the other parameters")
		u, err := url.Parse(BaseURL + VideosEP)
		if err != nil {
			return nil, retCursor, err
		}
		v := url.Values{}
		v.Add("id", qp.ID)
		u.RawQuery = v.Encode()
		uri = u.String()
	}

	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, retCursor, err
	}
	defer res.Body.Close()

	if err := parseResult(res, &retVideos); err != nil {
		return nil, retCursor, err
	}
	retCursor = retVideos.Pagination
	return retVideos.Data, retCursor, nil
}
