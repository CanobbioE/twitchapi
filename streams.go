package gwat

// GetStream gets information about active streams using options specified in a QueryParameters struct.
// Streams are returned sorted by number of current viewers, in descending order.
func (c *Client) GetStreams(qp StreamQueryParameters) ([]Stream, Cursor, error) {
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
func (c *Client) GetStreamsMetadata(qp StreamQueryParameters) ([]StreamMetadata, Cursor, error) {
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
