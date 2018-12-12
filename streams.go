package twitchapi

import "errors"

// GetStream gets information about active streams using options specified in a QueryParameters struct.
// Streams are returned sorted by number of current viewers, in descending order.
func (c *Client) GetStreams(qp StreamQueryParameters) ([]Stream, Cursor, error) {
	retCursor := Cursor{}
	retStreams := streamData{}

	qp.First = setDefaultValueIf(qp.First > 100, qp.First, 100).(int)
	qp.First = setDefaultValueIf(qp.First <= 0, qp.First, 20).(int)

	if len(qp.Language) > 100 || len(qp.ComunityID) > 100 || len(qp.UserID) > 100 || len(qp.UserLogin) > 100 {
		return nil, retCursor, errors.New("GetStreams: a parameter exceed the 100 characters limit")
	}

	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	uri := makeUri(BaseURL+StreamEP, qp)
	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, retCursor, err
	}
	defer res.Body.Close()

	// parse the result
	if res.Status != "200 OK" {
		return nil, retCursor, errors.New("GetStreams returned status:" + res.Status)
	}
	if err := parseResult(res, &retStreams); err != nil {
		return nil, retCursor, err
	}
	retCursor = retStreams.Pagination
	return retStreams.Data, retCursor, nil
}

// GetStreamsMetadata gets metadata information about active streams playing Overwatch or Heartstone.
// Streams are sorted by number of current viewers, in descending order
func (c *Client) GetStreamsMetadata(qp StreamQueryParameters) ([]StreamMetadata, Cursor, error) {
	retCursor := Cursor{}
	retMetaStreams := metaData{}

	qp.First = setDefaultValueIf(qp.First > 100, qp.First, 100).(int)
	qp.First = setDefaultValueIf(qp.First <= 0, qp.First, 20).(int)

	uri := makeUri(BaseURL+StreamEP, qp)
	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, retCursor, err
	}
	defer res.Body.Close()

	if err := parseResult(res, &retMetaStreams); err != nil {
		return nil, retCursor, err
	}
	retCursor = retMetaStreams.Pagination
	return retMetaStreams.Data, retCursor, nil
}
