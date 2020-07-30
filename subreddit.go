package goddit

import (
	"fmt"
)

func (sess *Session) SubredditData(subreddit string) (*Subreddit, error) {
	subredditAboutURL := fmt.Sprintf("%s/%s%s", SubredditURL, subreddit, aboutEnding)
	var getSubredditData GetSubredditData
	dataErr := sess.GetResponse(subredditAboutURL, GET, nil, &getSubredditData)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getSubredditData.Data, nil
}

func (sess *Session) SubredditRules(subreddit string, ruleParams map[string]interface{}) ([]Rule, error) {
	rulesURL := fmt.Sprintf("%s/%s%s", SubredditURL, subreddit, aboutRulesEnding)
	var getRules GetRules
	dataErr := sess.GetResponse(rulesURL, GET, nil, &getRules)
	if dataErr != nil {
		return nil, dataErr
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
	var getModerators GetModerators
	dataErr := sess.GetResponse(moderatorsURL, GET, nil, &getModerators)
	if dataErr != nil {
		return nil, dataErr
	}

	return getModerators.Data.Mods, nil
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
	dataErr := sess.GetResponse(subscribeParamURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}
