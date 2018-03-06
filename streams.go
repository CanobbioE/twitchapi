package gwat

import "fmt"

type Stream struct {
	ID           string   `json:"id"`
	UserID       string   `json:"user_id"`
	GameID       string   `json:"game_id"`
	ComunityIDs  []string `json:"comunity_ids"`
	Type         string   `json:"type"`
	Title        string   `json:"title"`
	ViewerCount  int      `json:"viewer_count"`
	StartedAt    string   `json:"started_at"`
	Language     string   `json:"language"`
	ThumbnailURL string   `json:"thumbnail_url"`
}

type streams struct {
	Data       []Stream `json:"data"`
	Pagination Cursor   `json:"pagination"`
}

type Cursor struct {
	Cursor string `json:"cursor"`
}

type QueryParameters struct {
	After      string
	Before     string
	ComunityID string
	First      int
	GameID     string
	Language   string
	Type       string
	UserID     string
	UserLogin  string
}

func (c *Client) GetStreams(qp QueryParameters) ([]Stream, Cursor, error) {
	uri := BaseURL + StreamEP
	retCursor := Cursor{}
	params := parseInput(qp)

	uri += "?"
	for k, v := range params {
		uri += fmt.Sprintf("%s=%v&", k, v)
	}
	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	res, err := c.request("GET", uri, h)
	if err != nil {
		return nil, retCursor, err
	}
	defer res.Body.Close()

	retStreams := streams{}
	if err := parseResult(res, &retStreams); err != nil {
		return nil, retCursor, err
	}
	retCursor = retStreams.Pagination
	return retStreams.Data, retCursor, nil
}
