package gwat

// ------------ Users -------------

// userData represents and array of user
type userData struct {
	Data []User
}

// User represents a user information as described by the Twitch API documentation.
type User struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadCasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Email           string `json:"email"`
}

// userFollowData represents the information returned by GetUsersFollows
type userFollowData struct {
	Total int           `json:"total"`
	Data  []UserFollows `json:"data"`
}

// UserFollows represents two users follow relationship
type UserFollows struct {
	FromID     string `json:"from_id"`
	ToID       string `json:"to_id"`
	FollowedAt string `json:"followed_at"`
}

// FollowQueryParameters represents the optional query string parameters used for API calls to the "follow" endpoint.
type FollowQueryParameters struct {
	After  string `after`
	Before string `before`
	First  int    `first`
	FromID string `from_id`
	ToID   string `to_id`
}

// ------------- Clips -------------

// clipData represents an array of Clip
type clipData struct {
	Data []Clip `json:"data"`
}

// Clip represent a clip as described by the twitch API documentation.
type Clip struct {
	ID            string `json:"id"`
	URL           string `json:"url"`
	EmbedURL      string `json:"embed_url"`
	BroadcasterID string `json:"broadcaster_id"`
	CreatorID     string `json:"creator_id"`
	VideoID       string `json:"video_id"`
	GameID        string `json:"game_id"`
	Language      string `json:"language"`
	Title         string `json:"title"`
	ViewCount     int    `json:"view_count"`
	CreatedAt     string `json:"created_at"`
	ThumbnailURL  string `json:"thumbnail_url"`
}

// clipInfoData represents an array of ClipInfo
type clipInfoData struct {
	Data []ClipInfo `json:"data"`
}

// DataClip represent the response generated by CreateClip as described by the twitch API documentation.
type ClipInfo struct {
	ID      string `json:"id"`
	editURL string `json:"edit_url"`
}

// ------------- Uploads -------------

// uploadData represents an array of uploadURL
type uploadData struct {
	Data []uploadURL `json:"data"`
}

// uploadURL represent the response returned by CreateEntitlementGrantsUploadURL.
type uploadURL struct {
	url string `json:"url"`
}

// ------------- Games -------------

// Game represents a game as described by the twitch API documentation.
type Game struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BoxArtURL string `json:"box_art_url"`
}

// games represents an array of Game
type games struct {
	Data []Game `json:data`
}

// ------------- Metadatas -------------

// StreamMetadata represents metadata information about a stream as described by the Twitch API documentation.
type StreamMetadata struct {
	UserID      string      `json:"user_id"`
	GameID      string      `json:"game_id"`
	Overwatch   Overwatch   `json:"overwatch"`
	Hearthstone Hearthstone `json:"hearthstone"`
}

// Overwatch represents the overwatch metadata information.
type Overwatch struct {
	Broadcaster OwBroadcaster `json:"broadcaster"`
}

// OwBroadcaster represents an Overwatch broadcaster.
type OwBroadcaster struct {
	Role    string `json:"role"`
	Name    string `json:"name"`
	Ability string `json:"ability"`
}

// Hearthstone represents the hearthstone metadata information.
type Hearthstone struct {
	Broadcaster Hero `json:"broadcaster"`
	Opponent    Hero `json:"opponent"`
}

// Hero represents the hero metadata information.
type Hero struct {
	Type  string `json:"type"`
	Class string `json:"class"`
	Name  string `json:"name"`
}

// metas represents an array of StreamMetadata
type metas struct {
	Data       []StreamMetadata `json:"data"`
	Pagination Cursor           `json:"pagination"`
}

// ------------- Streams -------------

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