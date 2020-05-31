package goddit

import (
	"fmt"
)

func (sess *Session) UserAbout(user string) (*User, error) {
	userAboutURL := fmt.Sprintf("%s/%s%s", UserURL, user, aboutEnding)
	req, RequestErr := RedditAPIRequest(GET, userAboutURL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	userAboutJson := &GetUser{}
	ResponseErr := sess.RedditAPIResponse(req, userAboutJson)

	if ResponseErr != nil {
		return nil, ResponseErr
	}

	return &userAboutJson.Data, nil
}

func (sess *Session) UserTrophies(user string) ([]Trophie, error) {
	userTrophiesURL := fmt.Sprintf("%s/%s%s", User2URL, user, trophiesEnding)
	req, RequestErr := RedditAPIRequest(GET, userTrophiesURL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	trophiesStruct := &GetTrophies{}
	ResponseErr := sess.RedditAPIResponse(req, trophiesStruct)
	if ResponseErr != nil {
		return nil, ResponseErr
	}

	var trophies []Trophie
	for _, trophieObject := range trophiesStruct.Data.TrophieList {
		trophies = append(trophies, trophieObject.Data)
	}

	return trophies, nil
}

func (sess *Session) UserFriends(user string) ([]FriendData, error) {
	userFriendsURL := fmt.Sprintf("%s/%s", UserFriendsBeginning, user)
	req, RequestErr := RedditAPIRequest(GET, userFriendsURL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	friends := []FriendData{}
	ResponseErr := sess.RedditAPIResponse(req, friends)

	if ResponseErr != nil {
		return nil, ResponseErr
	}

	return friends, nil
}

func (sess *Session) UserComments(user string) ([]Comment, error) {
	userCommentsURL := fmt.Sprintf("%s/%s%s", UserURL, user, commentsEnding)
	req, RequestErr := RedditAPIRequest(GET, userCommentsURL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	commentData := &GetCommentJSON{}
	ResponseErr := sess.RedditAPIResponse(req, commentData)

	if ResponseErr != nil {
		return nil, ResponseErr
	}

	comments := []Comment{}
	for _, commentStruct := range commentData.Data.Children {
		comments = append(comments, commentStruct.CommentData)
	}

	return comments, nil
}

func (sess *Session) UserPosts(user string) ([]Submission, error) {
	userPostsURL := fmt.Sprintf("%s/%s%s", UserURL, user, postsEnding)
	req, RequestErr := RedditAPIRequest(GET, userPostsURL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	submissionsStruct := &GetSubmission{}
	ResponseErr := sess.RedditAPIResponse(req, submissionsStruct)

	if ResponseErr != nil {
		return nil, ResponseErr
	}

	submissions := []Submission{}
	for _, submissionData := range submissionsStruct.Data.Children {
		submissions = append(submissions, submissionData.Data)
	}

	return submissions, nil
}

// func (sess *Session) upvoted() {

// }

// func (sess *Session) downvoted() {

// }

// func (sess *Session) hidden() {

// }

// func (sess *Session) saved() {

// }
