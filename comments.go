package goddit

import (
	"fmt"
)

func GetReplies(c Comment) []Comment {
	replies := []Comment{}
	for _, replydata := range c.Replies.Data.Children {
		replies = append(replies, replydata.Data)
	}

	return replies
}

func Replied(c Comment, username string) bool {
	for _, reply := range GetReplies(c) {
		if reply.Author == username {
			return true
		}
	}

	return false
}

func (sess *Session) GetComments(submission Submission) (*[]Comment, error) {
	commentURL := baseURL + submission.Permalink
	req, RequestErr := RedditAPIRequest(GET, commentURL, nil)

	if RequestErr != nil {
		return nil, RequestErr
	}

	commentJSON := &GetCommentJSON{}
	ResponseErr := sess.RedditAPIResponse(req, commentJSON)

	if ResponseErr != nil {
		return nil, ResponseErr
	}

	comments := []Comment{}
	for _, comment := range commentJSON.Data.Children {
		comments = append(comments, comment.Data)
	}

	return &comments, nil
}

func (sess *Session) Reply(c Comment, replyParams map[string]interface{}) (*Comment, error) {
	allReplyParams := []string{"api_type", "return_rtjson", "richtext_json", "text"}
	for key, _ := range replyParams {
		if !contains(allReplyParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	if value, ok := replyParams["text"]; ok {
		replyParams["text"] = checkSpace(value.(string))
	}

	replyURL := addParams(commentURL, replyParams)

	req, RequestErr := RedditAPIRequest(POST, replyURL, nil)

	if RequestErr != nil {
		return nil, RequestErr
	}

	postedComment := &GetPostedComment{}
	ResponseErr := sess.RedditAPIResponse(req, postedComment)

	if ResponseErr != nil {
		return nil, ResponseErr
	}

	comment := Comment{}
	for _, commentdata := range postedComment.JSON.Data.Things {
		comment = commentdata.Data
		break
	}

	return &comment, nil
}

func (sess *Session) Comment(sub Submission, commentParams map[string]interface{}) (*Comment, error) {
	allCommentParams := []string{"api_type", "return_rtjson", "richtext_json", "text"}
	for key, _ := range commentParams {
		if !contains(allCommentParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	if value, ok := commentParams["text"]; ok {
		commentParams["text"] = checkSpace(value.(string))
	}

	commentParamURL := addParams(commentURL, commentParams)

	req, RequestErr := RedditAPIRequest(POST, commentParamURL, nil)

	if RequestErr != nil {
		return nil, RequestErr
	}

	postedComment := &GetPostedComment{}
	ResponseErr := sess.RedditAPIResponse(req, postedComment)

	if ResponseErr != nil {
		return nil, ResponseErr
	}

	var comment Comment
	for _, commentdata := range postedComment.JSON.Data.Things {
		comment = commentdata.Data
	}

	return &comment, nil
}
