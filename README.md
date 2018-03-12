[![GoDoc](https://godoc.org/github.com/CanobbioE/twitchapi?status.png)](https://godoc.org/github.com/CanobbioE/twitchapi)
[![Build Status](https://travis-ci.org/CanobbioE/twitchapi.svg?branch=master)](https://travis-ci.org/CanobbioE/twitchapi)
<!--[![twitchapi](https://gocover.io/_badge/github.com/CanobbioE/twitchapi)](http://gocover.io/github.com/CanobbioE/twitchapi)-->
# TwitchAPI

## Description
This project is a go wrapper for the [New Twitch API](https://dev.twitch.tv/docs/api).  
twitchapi provides a set of functuins to perform calls to the new Twitch API.  
Although most functions checks for input correctness, I highly recomend to check the [API reference](https://dev.twitch.tv/docs/api/reference).

## Why twitchapi?
To this day this is the most updated and complete go wrapper for the new Twitch API.

## Installation
Using go get:
```
go get github.com/canobbioe/twitchapi
```

## Functions supported
- [CreateClip](https://dev.twitch.tv/docs/api/reference#create-clip) 
- [CreateEntitlementGrantsUploadURL](https://dev.twitch.tv/docs/api/reference#create-entitlement-grants-upload-url)
- [GetClip](https://dev.twitch.tv/docs/api/reference#get-clip)
- [GetGames](https://dev.twitch.tv/docs/api/reference#get-clip)
- [GetStreams](https://dev.twitch.tv/docs/api/reference#get-streams)
- [GetStreamsMetadata](https://dev.twitch.tv/docs/api/reference#get-streams-metadata)
- [GetTopGames](https://dev.twitch.tv/docs/api/reference#get-top-games)
- [GetUsers](https://dev.twitch.tv/docs/api/reference#get-users)
- [GetUsersFollows](https://dev.twitch.tv/docs/api/reference#get-users-follows)
- [UpdateUser](https://dev.twitch.tv/docs/api/reference#update-user)
- [GetVideos](https://dev.twitch.tv/docs/api/reference#get-videos)

## Authentication
Authentication is not yet implemented within this pkg.  
Authentication involves:  
- [Registering](https://dev.twitch.tv/dashboard/apps/create) your client.
- [Getting a token](https://dev.twitch.tv/docs/authentication#getting-tokens).
- [Sending a token](https://dev.twitch.tv/docs/authentication#sending-user-access-and-app-access-tokens).

## Usage
First of all make sure to have a [registered](https://dev.twitch.tv/docs/authentication#registration) client, then proceed as shown.  
```go
// Import the package
import "github.com/CanobbioE/twitchapi"

// Create a new client
c := twitchapi.NewClient("client-id")

// Create an input struct as needed
qp := StreamQueryParameters{
	First: 20,
	Language: []string{"en"},
}

// Perform the API call
streams, cursor, err := c.GetStreams(qp)
if err != nil {
	// handle error
}

// Access the values
for _, stream := range streams {
	fmt.Printf("%s : %d\n", stream.Title, stream.ViewerCount)
}
```
Output:
```
We gaming | @Ninja on twitter and Instagram : 110299

buff kaisa : 18888

hey there B U D   : 16138

༼ ºل͟º ༼ ºل͟º ༼ ºل͟º ༽ ºل͟º ༽ ºل͟º ༽ : 15337

★AMAZ★ No more Rag in Arena nooooo =D | Battlerite later!! : 11316

...
```


## TODO
- Try refactoring requests
- Use net/url
- Authentication
- Complete testing
- Committing test files
