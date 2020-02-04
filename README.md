[![GoDoc](https://godoc.org/github.com/CanobbioE/twitchapi?status.png)](https://godoc.org/github.com/CanobbioE/twitchapi)
[![Build Status](https://travis-ci.org/CanobbioE/twitchapi.svg?branch=master)](https://travis-ci.org/CanobbioE/twitchapi)
<!-- [![twitchapi](https://gocover.io/_badge/github.com/CanobbioE/twitchapi)](http://gocover.io/github.com/CanobbioE/twitchapi) -->

# TwitchAPI

# THIS IS NOT UP TO DATE
This project is no longer up to date with the newest endpoints that twitch API offers. Furthermore, now that I feel more confident with Go I am working on a better, cooler and possibly easier to use version of this.

## Description
This project is a go wrapper for the [New Twitch API](https://dev.twitch.tv/docs/api).  
The package provides a set of functcions to perform calls to the new Twitch API.  
Although most functions checks for input correctness, I highly recomend to check the [API reference](https://dev.twitch.tv/docs/api/reference) before starting to use _twitchapi_ .

## Why twitchapi?
To this day this is the most updated and complete go wrapper for the new Twitch API.

## Installation
Using go get:
```
go get github.com/canobbioe/twitchapi
```

## Functions supported
- [GetBitsLeaderboard](https://dev.twitch.tv/docs/api/reference/#get-bits-leaderboard)
- [CreateClip](https://dev.twitch.tv/docs/api/reference#create-clip) 
- [CreateEntitlementGrantsUploadURL](https://dev.twitch.tv/docs/api/reference#create-entitlement-grants-upload-url)
- [GetClip](https://dev.twitch.tv/docs/api/reference#get-clip)
- [GetGames](https://dev.twitch.tv/docs/api/reference#get-clip)
- [GetTopGames](https://dev.twitch.tv/docs/api/reference#get-top-games)
- [GetStreams](https://dev.twitch.tv/docs/api/reference#get-streams)
- [GetStreamsMetadata](https://dev.twitch.tv/docs/api/reference#get-streams-metadata)
- [GetUsers](https://dev.twitch.tv/docs/api/reference#get-users)
- [GetUsersFollows](https://dev.twitch.tv/docs/api/reference#get-users-follows)
- [UpdateUser](https://dev.twitch.tv/docs/api/reference#update-user)
- [GetVideos](https://dev.twitch.tv/docs/api/reference#get-videos)

## Authentication
Authentication is not yet implemented within this package.  
Authentication involves:  
- [Registering](https://dev.twitch.tv/dashboard/apps/create) your client.
- [Getting a token](https://dev.twitch.tv/docs/authentication#getting-tokens).
- [Sending a token](https://dev.twitch.tv/docs/authentication#sending-user-access-and-app-access-tokens).

## Usage
The documentation can be found on [go doc](https://godoc.org/github.com/CanobbioE/twitchapi).

### Example
Make sure to have a [registered](https://dev.twitch.tv/docs/authentication#registration) client, hence a client id.

This code gets the top five English streams and prints the resutlt:
```go
// Import the package
import (
	"fmt"
	"github.com/canobbioe/twitchapi"
)

func main() {
	// Create a new client
	c := twitchapi.NewClient("your-client-id")

	// Create an input struct as needed
	qp := twitchapi.StreamQueryParameters{
		First:    5,
		Language: []string{"en"},
	}

	// Perform the API call
	streams, _, err := c.GetStreams(qp)
	if err != nil {
		// handle error
	}

	// Access the values
	for _, stream := range streams {
		fmt.Printf("%s : %d\n", stream.Title, stream.ViewerCount)
	}
}
```
Output:
```
 TSM Dakotaz - 1400+ Wins 🏆 | youtube.com/dakotaz | twitter.com/dakotaz : 18768
LCK Spring: ROX vs. KZ - KSV vs. JAG : 15100
Giving GabeN all my money || [A] @AdmiralBulldog : 10714
IWD - Jungle Abuse : 6600
[PC] Day 2 of troll alert sounds, Sanity barely being held on to | Twitter @JacobHysteria 1100+ wins : 5570
```


## TODO
- There must be a way to simplify the code...
- Authentication
- GetExtensionAnalytic
- GetGameAnalytics
- CreateStreamMarker
- GetStreamMarkers
- GetUserExtensions
- GetUserActiveExtensions
- UpdateUserExtensions
- GetWebHookSubscriptions

