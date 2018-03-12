package twitchapi

import "errors"

// GetStream gets information about active streams using options specified in a QueryParameters struct.
// Streams are returned sorted by number of current viewers, in descending order.
func (c *Client) GetStreams(qp StreamQueryParameters) ([]Stream, Cursor, error) {
	retCursor := Cursor{}
	retStreams := streamData{}

	if qp.First > 100 {
		return nil, retCursor, errors.New("\"First\" parameter cannot be greater than 100")
	}

	uri := makeUri(StreamEP, qp)
	h := Header{
		Field: "Client-ID",
		Value: c.ClientID,
	}

	res, err := c.apiCall("GET", uri, h)
	if err != nil {
		return nil, retCursor, err
	}
	defer res.Body.Close()

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
	if qp.First > 100 {
		return nil, retCursor, errors.New("\"First\" parameter cannot be greater than 100")
	}

	uri := makeUri(StreamEP, qp)
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
