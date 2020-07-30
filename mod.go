package goddit

import (
	"fmt"
)

func (sess *Session) ModLog(subreddit string, modParams map[string]interface{}) ([]ModLog, error) {
	allModLogParams := []string{"after", "before", "count", "limit", "mod", "show", "sr_detail", "type"}
	for key := range modParams {
		if !contains(allModLogParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	modLogURL := fmt.Sprintf("%s/%s%s", SubredditURL, subreddit, modLogEnding)
	modLogURL = addParams(modLogURL, modParams)

	var getModLog GetModLog
	dataErr := sess.GetResponse(modLogURL, GET, nil, &getModLog)
	if dataErr != nil {
		return nil, dataErr
	}

	var modLog []ModLog
	for _, modLogObject := range getModLog.Data.Children {
		modLog = append(modLog, modLogObject.Data)
	}

	return modLog, nil
}

func (sess *Session) Approve(a interface{}) error {
	approveParams := make(map[string]interface{})
	approveParams["id"] = sess.getFullID(a)

	approveParamURL := addParams(approveURL, approveParams)

	dataErr := sess.GetResponse(approveParamURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) AcceptModInvite(subreddit string) error {
	acceptModParams := make(map[string]interface{})
	acceptModParams["api_type"] = "json"

	acceptModURL := fmt.Sprintf("%s/%s%s", SubredditURL, subreddit, acceptModEnding)
	acceptModURL = addParams(acceptModURL, acceptModParams)
	dataErr := sess.GetResponse(acceptModURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Remove(a interface{}, spam bool) error {
	removeParams := make(map[string]interface{})
	removeParams["id"] = sess.getFullID(a)
	removeParams["spam"] = spam

	removeParamURL := addParams(removeURL, removeParams)
	dataErr := sess.GetResponse(removeParamURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) ShowComment(c Comment) error {
	showCommentParams := make(map[string]interface{})
	showCommentParams["id"] = c.FullID

	showCommentParamURL := addParams(showCommentURL, showCommentParams)

	dataErr := sess.GetResponse(showCommentParamURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}
