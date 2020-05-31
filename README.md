# Goddit (A Reddit API Wrapper)

This is wrapper that is used to scrape off reddit, and or make of bot for reddit. With many built in functions. Also, my first project in Go.


### Logging in
```go
package main

import (
	goddit "github.com/KaleBBuR/Goddit"
)

func main() {
	username := "USERNAME"
	password := "PASSWORD"
	useragent := "USERAGENT123456789"
	clientID := "CLIENTID"
	clientSecret := "CLIENTSECRET"

	bot, err := goddit.OAuthLoginSession(
		username,
		password,
		useragent,
		clientID,
		clientSecret,
	)

	// All functions from comments.go
	bot.GetComments()
	bot.Reply()
	bot.Comment()

	// All functions from me.go
	bot.MyData()
	bot.MyPrefs()
	bot.MyKarma()
	bot.MyTrophies()

	// All functions from mod.go
	bot.ModLog()
	bot.Approve()
	bot.AccpetModInvite()
	bot.Remove()
	bot.ShowComment()

	// All functions from submission.go
	bot.GetSubmission()
	bot.GetSubmissions()
	bot.SearchSubmissions()
	bot.Hide()
	bot.Lock()
	bot.MarkNSFW()
	bot.Report()
	bot.Delete()
	bot.Edit()
	bot.Save()
	bot.SetSubredditSticky()
	bot.SetSuggestedSort()
	bot.Spolier()
	bot.StoreVists()
	bot.Post()
	bot.Unhide()
	bot.Unlock()
	bot.UnmarkNSFW()
	bot.Unsave()
	bot.Unspoiler()
	bot.Vote()

	// All functions from subreddit.go
	bot.SubredditData()
	bot.SubredditRules()
	bot.Moderators()
	bot.Subscribe()

	// All Functions from user.go
	bot.UserAbout()
	bot.UserTrophies()
	bot.UserFriends()
	bot.UserComments()
	bot.UserPosts()
}
```

I would love feedback on this.

Yes, I do plan on updating the "Documentation"

Hope you guys use this for some amazing projects.

Thanks for looking!
