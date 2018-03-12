package twitchapi

import (
	"errors"
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

	if !isNil(qp.ID) {
		var err error
		err = isValid("period", qp.Period, []string{"all", "day", "month", "week"})
		if err != nil {
			return nil, retCursor, err
		}
		err = isValid("sort", qp.Sort, []string{"time", "trending", "views"})
		if err != nil {
			return nil, retCursor, err
		}
		err = isValid("type", qp.Type, []string{"all", "upload", "archive", "highlight"})
		if err != nil {
			return nil, retCursor, err
		}
		if qp.First > 100 {
			return nil, retCursor, errors.New("GetVideos: \"First\" parameter cannot be greater than 100")
		}
		uri = makeUri(BaseURL+VideosEP, qp)
	} else {
		fmt.Println("GetVideos: \"id\" was specified. Ignoring all the other parameters")
		var u url.URL
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

	// parse
	if err := parseResult(res, &retVideos); err != nil {
		return nil, retCursor, err
	}
	retCursor = retVideos.Pagination
	return retVideos.Data, retCursor, nil
}
