package twitchapi

/*
	This file contains all the input used for the different API calls.
*/

// FollowQueryParameters represents the optional query string's parameters used for API calls to the "follow" endpoint.
// At minimum from_id or to_id must be provided for a query to be valid.
type FollowQueryParameters struct {
	After  string `after`   // Optional. Cursor for forward pagination: tells the server where to start fetching the next set of results.
	Before string `before`  // Optional. Cursor for backward pagination: tells the server where to start fetching the next set of results.
	First  int    `first`   // Optional. Maximum number of objects to return. Maximum: 100. Default: 20.
	FromID string `from_id` // User ID. The request returns information about users who are being followed by the from_id user.
	ToID   string `to_id`   // User ID. The request returns information about users who are following the to_id user.
}

// StreamQueryParameters represents the optional query string parameters used for API calls to the "stream" endpoint.
type StreamQueryParameters struct {
	After      string   `after`       // Optional. Cursor for forward pagination: tells the server where to start fetching the next set of results.
	Before     string   `before`      //Optional. Cursor for backward pagination: tells the server where to start fetching the next set of results.
	ComunityID []string `comunity_id` //Optional. Returns streams in a specified community ID. You can specify up to 100 IDs.
	First      int      `first`       //Optional. Maximum number of objects to return. Maximum: 100. Default: 20.
	GameID     string   `game_id`     //Optional. Returns streams broadcasting the specified game ID. You can specify up to 100 IDs.
	Language   string   `language`    //Optional. Stream language. You can specify up to 100 languages.
	Type       string   `type`        //Optional. Stream type: "all", "live", "vodcast". Default: "all".
	UserID     string   `user_id`     //Optional. Returns streams broadcast by one or more of the specified user IDs. You can specify up to 100 IDs.
	UserLogin  string   `user_login`  //Optional. Returns streams broadcast by one or more of the specified user login names. You can specify up to 100 names.
}

// VideoQueryParameters represents the query string's parameters used for API calls to the "video" endpoint.
// Each request must specify one or more video ids, one user_id, or one game_id.
type VideoQueryParameters struct {
	ID       string `id`       // Required. ID of the video being queried. Limit: 100. If this is specified, you cannot use any of the optional parameters below.
	UserID   string `user_id`  // Required. ID of the user who owns the video. Limit 1.
	GameID   string `game_id`  // Required. ID of the game the video is of. Limit 1.
	After    string `after`    // Optional. Cursor for forward pagination: tells the server where to start fetching the next set of results.
	Before   string `before`   // Optional. Cursor for backward pagination: tells the server where to start fetching the next set of results.
	First    int    `first`    // Optional. Number of values to be returned when getting videos by user or game ID. Limit: 100. Default: 20.
	Language string `language` // Optional. Language of the video being queried. Limit: 1.
	Period   string `period`   // Optional. Period during which the video was created. Valid values: "all", "day", "month", and "week". Default: "all".
	Sort     string `sort`     // Optional. Sort order of the videos. Valid values: "time", "trending", and "views". Default: "time".
	Type     string `type`     // Optional. Type of video. Valid values: "all", "upload", "archive", and "highlight". Default: "all".
}

// EntitlementURLQueryParameters represents the query string's parameters used for API calls to the "upload" endpoint.
type EntitlementURLQueryParameters struct {
	ManifestID string `manifest_id` // Required. Unique identifier of the manifest file to be uploaded. Must be 1-64 characters.
	Type       string `type`        // Required. Type of entitlement being granted. Only "bulk_drops_grant" is supported.
}

// GameQueryParameters represents the query string's parameters used for API calls to the "games" endpoint.
// For a query to be valid, name and/or id must be specified.
type GameQueryParameters struct {
	IDs   []string `id`   // Required. Game ID. At most 100 id values can be specified.
	Names []string `name` // Required. Game name. The name must be an exact match. At most 100 name values can be specified.
}

// TopGameQueryParameters represents the query string's parameters used for API calls to the "games/top" endpoint.
type TopGameQueryParameters struct {
	After  string `after`  // Optional. Cursor for forward pagination: tells the server where to start fetching the next set of results.
	Before string `before` // Optional. Cursor for backward pagination: tells the server where to start fetching the next set of results.
	First  int    `first`  // Optional. Maximum number of objects to return. Maximum: 100. Default: 20.
}

// UserQueryParameters represents the query string's parameters used for API calls to the "user" endpoint.
type UserQueryParameters struct {
	IDs    []string `id`    // Optional. User ID. Multiple user IDs can be specified. Limit: 100.
	Logins []string `login` // Optional. User login name. Multiple login names can be specified. Limit: 100.
}
