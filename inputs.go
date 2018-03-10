package twitchapi

/*
	This file contains all the input used for the different API calls
*/

// FollowQueryParameters represents the optional query string's parameters used for API calls to the "follow" endpoint.
type FollowQueryParameters struct {
	After  string `after`
	Before string `before`
	First  int    `first`
	FromID string `from_id`
	ToID   string `to_id`
}

// StreamQueryParameters represents the optional query string parameters used for API calls to the "stream" endpoint.
type StreamQueryParameters struct {
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

// VideoQueryParameters represents the query string's parameters used for API calls to the "video" endpoint.
type VideoQueryParameters struct {
	ID       string `id`      // required
	UserID   string `user_id` // required
	GameID   string `game_id` // required
	After    string `after`
	Before   string `before`
	First    string `first`
	Language string `language`
	Period   string `period`
	Sort     string `sort`
	Type     string `type`
}
