package gwat

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
