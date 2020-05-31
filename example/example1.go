package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	goddit ".."
)

var AllComments []goddit.Comment

type Private struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Useragent    string `json:"useragent"`
	ClientID     string `json:"client ID"`
	ClientSecret string `json:"client Secret"`
}

func main() {
	start := time.Now()

	private := Private{}
	raw, err := ioutil.ReadFile("private.json")
	if err != nil {
		return
	}

	json.Unmarshal(raw, &private)

	username := private.Username
	password := private.Password
	useragent := private.Useragent
	clientID := private.ClientID
	clientSecret := private.ClientSecret

	bot, err := goddit.OAuthLoginSession(
		username,
		password,
		useragent,
		clientID,
		clientSecret,
	)

	if err != nil {
		panic(err)
	}

	subreddits := []string{
		"AskReddit",
		"GamersRiseUp",
		"memes",
	}

	submissionParams := make(map[string]interface{})
	submissionParams["limit"] = 10

	for _, subreddit := range subreddits {
		if submissions, submissionErr := bot.GetSubmissions(subreddit, "hot", submissionParams); submissionErr == nil {
			for _, submission := range submissions {
				if comments, commentsErr := bot.GetComments(submission); commentsErr == nil {
					for _, comment := range comments {
						if comment.Author != "" {
							AllComments = append(AllComments, comment)
							Replies(comment)
						}
					}
				}
			}
		}
	}

	fmt.Println(len(AllComments))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s\n", elapsed)
}

func Replies(comment goddit.Comment) {
	for _, reply := range goddit.GetReplies(comment) {
		if reply.Author != "" && !Contain(reply) {
			AllComments = append(AllComments, reply)
			Replies(reply)
		}
	}
}

func Contain(comment goddit.Comment) bool {
	for _, arrComment := range AllComments {
		if arrComment.FullID == comment.FullID {
			return true
		}
	}

	return false
}
