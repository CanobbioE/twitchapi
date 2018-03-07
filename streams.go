package gwat

// Stream represents a stream as described by the Twitch API documentation.
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

// streams represents an array of Stream.
type streams struct {
	Data       []Stream `json:"data"`
	Pagination Cursor   `json:"pagination"`
}

// Cursor represents a cursor as described by the Twitch API documentation.
type Cursor struct {
	Cursor string `json:"cursor"`
}

// QueryParameters represents the optional query string parameters used for API calls.
type QueryParameters struct {
	After      string   `after`
	Before     string   `before`
	ComunityID []string `comunity_id`
	First      int      `first`
	GameID     string   `game_id`
	Language   string   `language`
	Type       string   `type`
	UserID     string   `user_id`
	UserLogin  string   `user_login`
}

// GetStream gets information about active streams using options specified in a QueryParameters struct.
// Streams are returned sorted by number of current viewers, in descending order.
func (c *Client) GetStreams(qp QueryParameters) ([]Stream, Cursor, error) {
	uri := BaseURL + StreamEP
	retCursor := Cursor{}

	res, err := c.streamRequest(&uri, qp)
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

//GetStreamsMetadata gets metadata information about active streams playing Overwatch or Heartstone.
// Streams are sorted by number of current viewers, in descending order
func (c *Client) GetStreamsMetadata(qp QueryParameters) ([]StreamMetadata, Cursor, error) {
	uri := BaseURL + StreamEP + MetaDataEP
	retCursor := Cursor{}

	res, err := c.streamRequest(&uri, qp)
	if err != nil {
		return nil, retCursor, err
	}
	defer res.Body.Close()

	retMetaStreams := metas{}
	if err := parseResult(res, &retMetaStreams); err != nil {
		return nil, retCursor, err
	}
	retCursor = retMetaStreams.Pagination
	return retMetaStreams.Data, retCursor, nil
}
