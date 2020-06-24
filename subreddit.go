package goddit

import (
	"fmt"
)

func (sess *Session) SubredditData(subreddit string) (*Subreddit, error) {
	subredditAboutURL := fmt.Sprintf("%s/%s%s", SubredditURL, subreddit, aboutEnding)
	req, RequestErr := RedditAPIRequest(GET, subredditAboutURL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	basesubredditjson := &GetSubredditData{}
	ResponseErr := sess.RedditAPIResponse(req, basesubredditjson)
	if ResponseErr != nil {
		return nil, ResponseErr
	}

	return &basesubredditjson.Data, nil
}

func (sess *Session) SubredditRules(subreddit string, ruleParams map[string]interface{}) ([]Rule, error) {
	rulesURL := fmt.Sprintf("%s/%s%s", SubredditURL, subreddit, aboutRulesEnding)
	req, RequestErr := RedditAPIRequest(GET, rulesURL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	getRules := &GetRules{}
	ResponseErr := sess.RedditAPIResponse(req, getRules)

	if ResponseErr != nil {
		return nil, ResponseErr
	}

	return getRules.Rules, nil
}

func (sess *Session) Moderators(subreddit string, modParams map[string]interface{}) ([]Moderator, error) {
	possibleSubModParams := []string{"after", "before", "count", "limit", "show", "sr_detail", "user"}

	for key := range modParams {
		if !contains(possibleSubModParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	moderatorsURL := fmt.Sprintf("%s/%s%s", SubredditURL, subreddit, aboutModsEnding)
	moderatorsURL = addParams(moderatorsURL, modParams)

	req, RequestErr := RedditAPIRequest(GET, moderatorsURL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	getMods := &GetModerators{}

	ResponseErr := sess.RedditAPIResponse(req, getMods)
	if ResponseErr != nil {
		return nil, ResponseErr
	}

	return getMods.Data.Mods, nil
}

func (sess *Session) Subscribe(a interface{}, subscribeParams map[string]interface{}) error {
	possibleSubscribeParams := []string{"action", "skip_initial_defaults"}
	for key := range subscribeParams {
		if !contains(possibleSubscribeParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	switch a.(type) {
	case Subreddit:
		subscribeParams["sr"] = a.(Subreddit).FullID
	case string:
		subscribeParams["sr_name"] = a.(string)
	default:
		panic("Must be a string or Subreddit type.")
	}

	subscribeParamURL := addParams(subscribeURL, subscribeParams)

	req, RequestErr := RedditAPIRequest(POST, subscribeParamURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}
