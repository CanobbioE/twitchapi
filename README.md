<!--- build:travisyml doc:go -->
# TwitchAPI
This project is a go wrapper for the [New Twitch API](https://dev.twitch.tv/docs/api).  
twitchapi provides a set of functuins to perform calls to the new Twitch API.  
Although most functions checks for input correctness, I highly recomend to check the [API reference](https://dev.twitch.tv/docs/api/reference).

## Why twitchapi?
To this day this is the most updated and complete go wrapper for the new Twitch API.

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
- Change request to something more significant
- Change Client to TwitchClient
- Try refactoring requests
- Check exported types
- Check functions names
- Tests
- Readme
- Change repo name
