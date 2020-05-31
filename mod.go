package goddit

import (
	"fmt"
)

func (sess *Session) ModLog(subreddit string, modParams map[string]interface{}) (*[]ModLog, error) {
	allModLogParams := []string{"after", "before", "count", "limit", "mod", "show", "sr_detail", "type"}
	for key, _ := range modParams {
		if !contains(allModLogParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	modLogURL := fmt.Sprintf("%s/%s%s", SubredditURL, subreddit, modLogEnding)
	modLogURL = addParams(modLogURL, modParams)

	req, RequestErr := RedditAPIRequest(GET, modLogURL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	modLogJSON := &GetModLog{}
	ResponseErr := sess.RedditAPIResponse(req, modLogJSON)
	if ResponseErr != nil {
		return nil, ResponseErr
	}

	var modLog []ModLog
	for _, modLogObject := range modLogJSON.Data.Children {
		modLog = append(modLog, modLogObject.Data)
	}

	return &modLog, nil
}

func (sess *Session) Approve(a interface{}) error {
	approveParams := make(map[string]interface{})
	approveParams["id"] = sess.getFullID(a)

	approveParamURL := addParams(approveURL, approveParams)
	req, RequestErr := RedditAPIRequest(POST, approveParamURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) AcceptModInvite(subreddit string) error {
	acceptModParams := make(map[string]interface{})
	acceptModParams["api_type"] = "json"

	acceptModURL := fmt.Sprintf("%s/%s%s", SubredditURL, subreddit, acceptModEnding)
	acceptModURL = addParams(acceptModURL, acceptModParams)
	req, RequestErr := RedditAPIRequest(POST, acceptModURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Remove(a interface{}, spam bool) error {
	removeParams := make(map[string]interface{})
	removeParams["id"] = sess.getFullID(a)
	removeParams["spam"] = spam

	removeParamURL := addParams(removeURL, removeParams)
	req, RequestErr := RedditAPIRequest(POST, removeParamURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) ShowComment(c Comment) error {
	showCommentParams := make(map[string]interface{})
	showCommentParams["id"] = c.FullID

	showCommentParamURL := addParams(showCommentURL, showCommentParams)

	req, RequestErr := RedditAPIRequest(POST, showCommentParamURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}
