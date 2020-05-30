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
	bot.NewComment()
	bot.DeleteComment()
	bot.EditComment()
	bot.LockComment()
	bot.ReportComment()
	bot.SaveComment()
	bot.UnsaveComment()
	bot.VoteComment()
	bot.UnlockComment()

	// All functions from me.go
	bot.MyData()
	bot.MyPrefs()
	bot.MyKarma()
	bot.MyTrophies()

	// All functions from mod.go
	bot.ModLog()
	bot.ApproveComment()
	bot.ApprovePost()
	bot.AccpetModInvite()
	bot.RemovePost()
	bot.RemoveComment()
	bot.ShowComment()

	// All functions from submission.go
	bot.GetSubmission()
	bot.HotSubmissions()
	bot.NewSubmissions()
	bot.TopSubmissions()
	bot.RisingSubmissions()
	bot.BestSubmissions()
	bot.ControversialSubmissions()
	bot.SearchSubmissions()
	bot.LockPost()
	bot.MarkPostNSFW()
	bot.ReportPost()
	bot.DeletePost()
	bot.EditPost()
	bot.SavePost()
	bot.SpolierLink()
	bot.Post()
	bot.UnlockPost()
	bot.UnmarkNSFWPost()
	bot.unsavePost()
	bot.UnspoilerPost()
	bot.VotePost()

	// All functions from subreddit.go
	bot.SubredditData()
	bot.SubredditRules()
	bot.Moderators()

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
