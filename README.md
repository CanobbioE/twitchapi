[![GoDoc](https://godoc.org/github.com/CanobbioE/twitchapi?status.png)](https://godoc.org/github.com/CanobbioE/twitchapi)
[![Build Status](https://travis-ci.org/CanobbioE/twitchapi.svg?branch=master)](https://travis-ci.org/CanobbioE/twitchapi)
[![twitchapi](https://gocover.io/_badge/github.com/CanobbioE/twitchapi)](http://gocover.io/github.com/CanobbioE/twitchapi)
# TwitchAPI
This project is a go wrapper for the [New Twitch API](https://dev.twitch.tv/docs/api).  
twitchapi provides a set of functuins to perform calls to the new Twitch API.  
Although most functions checks for input correctness, I highly recomend to check the [API reference](https://dev.twitch.tv/docs/api/reference).

## Why twitchapi?
To this day this is the most updated and complete go wrapper for the new Twitch API.

## Usage
Import the package:
```go
import "github.com/CanobbioE/twitchapi"
```

[Register](https://dev.twitch.tv/docs/authentication#registration) a client.

Create a new client:
```go
c := twitchapi.NewClient("client-id")
```

Create an input struct as needed:
```go
qp := GameQueryParameters{
	ID: "493057"
}
```

Perform the API call:
```go
games, err := c.GetGames()
if err != nil {
	// do something with error
}
```

Access the values:
```go
for _, game := range games {
	fmt.Println(game.Name)
}
```
Output:
```
PLAYER'S UNKOWN BATTLEGROUND
PLAYER'S UNKOWN BATTLEGROUND
...
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

## TODO
- Try refactoring requests
- Check exported types
- Tests
- Readme
- Use net/url
- Authentication
- Complete testing
