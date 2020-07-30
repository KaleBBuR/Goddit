package goddit

import (
	"fmt"
)

func GetReplies(c Comment) []Comment {
	replies := []Comment{}
	for _, replydata := range c.Replies.Data.Children {
		replies = append(replies, replydata.CommentData)
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

func (sess *Session) GetComments(submission Submission) ([]Comment, error) {
	commentURL := baseURL + submission.Permalink
	var getCommentJSON []GetCommentJSON
	dataErr := sess.GetResponse(commentURL, GET, nil, getCommentJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	comments := []Comment{}
	for _, commentData := range getCommentJSON {
		for _, comment := range commentData.Data.Children {
			comments = append(comments, comment.CommentData)
		}
	}

	return comments, nil
}

func (sess *Session) Reply(c Comment, replyParams map[string]interface{}) (*Comment, error) {
	allReplyParams := []string{"api_type", "return_rtjson", "richtext_json", "text"}
	for key := range replyParams {
		if !contains(allReplyParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	if value, ok := replyParams["text"]; ok {
		replyParams["text"] = checkSpace(value.(string))
	}

	replyURL := addParams(commentURL, replyParams)
	var getPostedComment GetPostedComment
	dataErr := sess.GetResponse(replyURL, POST, nil, &getPostedComment)
	if dataErr != nil {
		return nil, dataErr
	}

	comment := Comment{}
	for _, commentdata := range getPostedComment.JSON.Data.Things {
		comment = commentdata.CommentData
		break
	}

	return &comment, nil
}

func (sess *Session) Comment(sub Submission, commentParams map[string]interface{}) (*Comment, error) {
	allCommentParams := []string{"api_type", "return_rtjson", "richtext_json", "text"}
	for key := range commentParams {
		if !contains(allCommentParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	if value, ok := commentParams["text"]; ok {
		commentParams["text"] = checkSpace(value.(string))
	}

	commentParamURL := addParams(commentURL, commentParams)
	var getPostedComment GetPostedComment
	dataErr := sess.GetResponse(commentParamURL, POST, nil, &getPostedComment)
	if dataErr != nil {
		return nil, dataErr
	}

	var comment Comment
	for _, commentdata := range getPostedComment.JSON.Data.Things {
		comment = commentdata.CommentData
	}

	return &comment, nil
}
