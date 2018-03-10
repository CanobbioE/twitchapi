<!--- build:travisyml doc:go -->
# TwitchAPI
This project is a go wrapper for the [New Twitch API](https://dev.twitch.tv/docs/api).  
twitchapi provides a set of functuins to perform calls to the new Twitch API.  
Although most functions checks for input correctness, I highly recomend to check the [API reference](https://dev.twitch.tv/docs/api/reference).

## Why twitchapi?
To this day this is the most updated and complete go wrapper for the new Twitch API.

## Functions supported
- [CreateClip](https://dev.twitch.tv/docs/api/reference#create-clip) - ([code](https://github.com/CanobbioE/twitchapi/blob/08bfab66e2f2ca4136dea52b597b47573c6a0218/clips.go#L10))
- [CreateEntitlementGrantsUploadURL](https://dev.twitch.tv/docs/api/reference#create-entitlement-grants-upload-url) - ([code]())
- [GetClip](https://dev.twitch.tv/docs/api/reference#get-clip) - ([codde](https://github.com/CanobbioE/twitchapi/blob/08bfab66e2f2ca4136dea52b597b47573c6a0218/clips.go#L42))
- [GetGames](https://dev.twitch.tv/docs/api/reference#get-clip) - ([code]())
- [GetStreams](https://dev.twitch.tv/docs/api/reference#get-streams) - ([code]())
- [GetStreamsMetadata](https://dev.twitch.tv/docs/api/reference#get-streams-metadata) - ([code]())
- [GetTopGames](https://dev.twitch.tv/docs/api/reference#get-top-games) - ([code]())
- [GetUsers](https://dev.twitch.tv/docs/api/reference#get-users) - ([code]())
- [GetUsersFollows](https://dev.twitch.tv/docs/api/reference#get-users-follows) - ([code]())
- [UpdateUser](https://dev.twitch.tv/docs/api/reference#update-user) - ([code]())
- [GetVideos](https://dev.twitch.tv/docs/api/reference#get-videos) - ([code]())

## TODO
- Change request to something more significant
- Change Client to TwitchClient
- Try refactoring requests
- Check exported types
- Check functions names
- Tests
- Readme
- Change repo name
