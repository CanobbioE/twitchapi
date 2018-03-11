<!--- build:travisyml doc:go -->
# TwitchAPI
This project is a go wrapper for the [New Twitch API](https://dev.twitch.tv/docs/api).  
twitchapi provides a set of functuins to perform calls to the new Twitch API.  
Although most functions checks for input correctness, I highly recomend to check the [API reference](https://dev.twitch.tv/docs/api/reference).

## Why twitchapi?
To this day this is the most updated and complete go wrapper for the new Twitch API.

## Functions supported
- [CreateClip](https://dev.twitch.tv/docs/api/reference#create-clip) - ([code](https://github.com/CanobbioE/twitchapi/blob/08bfab66e2f2ca4136dea52b597b47573c6a0218/clips.go#L10))
- [CreateEntitlementGrantsUploadURL](https://dev.twitch.tv/docs/api/reference#create-entitlement-grants-upload-url) - ([code](https://github.com/CanobbioE/twitchapi/blob/d419966d62121a7cd8836d0dafbaaba1fe6fe513/entitlements.go#L8))
- [GetClip](https://dev.twitch.tv/docs/api/reference#get-clip) - ([code](https://github.com/CanobbioE/twitchapi/blob/08bfab66e2f2ca4136dea52b597b47573c6a0218/clips.go#L42))
- [GetGames](https://dev.twitch.tv/docs/api/reference#get-clip) - ([code](https://github.com/CanobbioE/twitchapi/blob/d419966d62121a7cd8836d0dafbaaba1fe6fe513/games.go#L10))
- [GetStreams](https://dev.twitch.tv/docs/api/reference#get-streams) - ([code](https://github.com/CanobbioE/twitchapi/blob/d419966d62121a7cd8836d0dafbaaba1fe6fe513/streams.go#L5))
- [GetStreamsMetadata](https://dev.twitch.tv/docs/api/reference#get-streams-metadata) - ([code](https://github.com/CanobbioE/twitchapi/blob/d419966d62121a7cd8836d0dafbaaba1fe6fe513/streams.go#L25))
- [GetTopGames](https://dev.twitch.tv/docs/api/reference#get-top-games) - ([code](https://github.com/CanobbioE/twitchapi/blob/d419966d62121a7cd8836d0dafbaaba1fe6fe513/streams.go#L25))
- [GetUsers](https://dev.twitch.tv/docs/api/reference#get-users) - ([code](https://github.com/CanobbioE/twitchapi/blob/d419966d62121a7cd8836d0dafbaaba1fe6fe513/users.go#L11))
- [GetUsersFollows](https://dev.twitch.tv/docs/api/reference#get-users-follows) - ([code](https://github.com/CanobbioE/twitchapi/blob/d419966d62121a7cd8836d0dafbaaba1fe6fe513/users.go#L52))
- [UpdateUser](https://dev.twitch.tv/docs/api/reference#update-user) - ([code](https://github.com/CanobbioE/twitchapi/blob/d419966d62121a7cd8836d0dafbaaba1fe6fe513/users.go#L86))
- [GetVideos](https://dev.twitch.tv/docs/api/reference#get-videos) - ([code](https://github.com/CanobbioE/twitchapi/blob/d419966d62121a7cd8836d0dafbaaba1fe6fe513/videos.go#L10))

## TODO
- Try refactoring requests
- Check exported types
- Tests
- Readme
- Use net/url
