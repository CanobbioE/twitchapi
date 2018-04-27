package twitchapi

/*
	This file contains all the input used for the different API calls.
*/

// FollowQueryParameters represents the optional query string's parameters used for API calls to the "follow" endpoint.
// At minimum from_id or to_id must be provided for a query to be valid.
type FollowQueryParameters struct {
	// Optional. Cursor for forward pagination: tells the server where to start fetching the next set of results.
	After string `param:"after"`
	// Optional. Cursor for backward pagination: tells the server where to start fetching the next set of results.
	Before string `param:"before"`
	// Optional. Maximum number of objects to return. Maximum: 100. Default: 20.
	First int `param:"first"`
	// User ID. The request returns information about users who are being followed by the from_id user.
	FromID string `param:"from_id"`
	// User ID. The request returns information about users who are following the to_id user.
	ToID string `param:"to_id"`
}

// StreamQueryParameters represents the optional query string parameters used for API calls to the "stream" endpoint.
type StreamQueryParameters struct {
	// Optional. Cursor for forward pagination: tells the server where to start fetching the next set of results.
	After string `param:"after"`
	//Optional. Cursor for backward pagination: tells the server where to start fetching the next set of results.
	Before string `param:"before"`
	//Optional. Returns streams in a specified community ID. You can specify up to 100 IDs.
	ComunityID []string `param:"comunity_id"`
	//Optional. Maximum number of objects to return. Maximum: 100. Default: 20.
	First int `param:"first"`
	//Optional. Returns streams broadcasting the specified game ID. You can specify up to 100 IDs.
	GameID string `param:"game_id"`
	//Optional. Stream language. You can specify up to 100 languages.
	Language []string `param:"language"`
	//Optional. Stream type: "all", "live", "vodcast". Default: "all".
	Type string `param:"type"`
	//Optional. Returns streams broadcast by one or more of the specified user IDs. You can specify up to 100 IDs.
	UserID []string `param:"user_id"`
	//Optional. Returns streams broadcast by one or more of the specified user login names. You can specify up to 100 names.
	UserLogin []string `param:"user_login"`
}

// VideoQueryParameters represents the query string's parameters used for API calls to the "video" endpoint.
// Each request must specify one or more video ids, one user_id, or one game_id.
type VideoQueryParameters struct {
	// Required. ID of the video being queried. Limit: 100. If this is specified, you cannot use any of the optional parameters below.
	ID string `param:"id"`
	// Required. ID of the user who owns the video. Limit 1.
	UserID string `param:"user_id"`
	// Required. ID of the game the video is of. Limit 1.
	GameID string `param:"game_id"`
	// Optional. Cursor for forward pagination: tells the server where to start fetching the next set of results.
	After string `param:"after"`
	// Optional. Cursor for backward pagination: tells the server where to start fetching the next set of results.
	Before string `param:"before"`
	// Optional. Number of values to be returned when getting videos by user or game ID. Limit: 100. Default: 20.
	First int `param:"first"`
	// Optional. Language of the video being queried. Limit: 1.
	Language string `param:"language"`
	// Optional. Period during which the video was created. Valid values: "all", "day", "month", and "week". Default: "all".
	Period string `param:"period"`
	// Optional. Sort order of the videos. Valid values: "time", "trending", and "views". Default: "time".
	Sort string `param:"sort"`
	// Optional. Type of video. Valid values: "all", "upload", "archive", and "highlight". Default: "all".
	Type string `param:"type"`
}

// EntitlementURLQueryParameters represents the query string's parameters used for API calls to the "upload" endpoint.
type EntitlementURLQueryParameters struct {
	// Required. Unique identifier of the manifest file to be uploaded. Must be 1-64 characters.
	ManifestID string `param:"manifest_id"`
	// Required. Type of entitlement being granted. Only "bulk_drops_grant" is supported.
	Type string `param:"type"`
}

// GameQueryParameters represents the query string's parameters used for API calls to the "games" endpoint.
// For a query to be valid, name and/or id must be specified.
type GameQueryParameters struct {
	// Required. Game ID. At most 100 id values can be specified.
	IDs []string `param:"id"`
	// Required. Game name. The name must be an exact match. At most 100 name values can be specified.
	Names []string `param:"name"`
}

// TopGameQueryParameters represents the query string's parameters used for API calls to the "games/top" endpoint.
type TopGameQueryParameters struct {
	// Optional. Cursor for forward pagination: tells the server where to start fetching the next set of results.
	After string `param:"after"`
	// Optional. Cursor for backward pagination: tells the server where to start fetching the next set of results.
	Before string `param:"before"`
	// Optional. Maximum number of objects to return. Maximum: 100. Default: 20.
	First int `param:"first"`
}

// UserQueryParameters represents the query string's parameters used for API calls to the "user" endpoint.
type UserQueryParameters struct {
	// Optional. User ID. Multiple user IDs can be specified. Limit: 100.
	IDs []string `param:"id"`
	// Optional. User login name. Multiple login names can be specified. Limit: 100.
	Logins []string `param:"login"`
}
