package gwat

import (
	"fmt"
)

// GetVideos gets video information by video ID (one or more), user ID (one only), or game ID (one only).
// If id is specified any other parameter is ignored.
// For lookup by user or game ID a Cursor is returned.
func (c *Client) GetVideos(vq VideoQueryParameters) ([]Video, Cursor, error) {
	uri := BaseURL + VideosEP
	retVideos := videoData{}
	retCursor := Cursor{}

	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	// check params
	params := parseInput(vq)
	uri += "?"
	if id, ok := params["id"]; !ok {
		for k, v := range params {
			if k == "period" {
				err := isValid(k, v.(string), []string{"all", "day", "month", "week"})
				if err != nil {
					return nil, retCursor, err
				}
			}
			if k == "sort" {
				err := isValid(k, v.(string), []string{"time", "trending", "views"})
				if err != nil {
					return nil, retCursor, err
				}
			}
			if k == "type" {
				err := isValid(k, v.(string), []string{"all", "upload", "archive", "highlight"})
				if err != nil {
					return nil, retCursor, err
				}
			}
			// if k == "first" && strconv.Atoi(v) > 100 { return nil, errors.New()}
			uri += fmt.Sprintf("%s=%v&", k, v)
		}
	} else {
		fmt.Println("GetVideos: \"id\" was specified. Ignoring all the other parameters")
		if t, ok := id.([]string); ok {
			addParameters(&uri, "id", t)
		} else {
			uri += "id=" + id.(string)
		}
	}

	// perform request
	res, err := c.request("GET", uri, h)
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
