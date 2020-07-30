package goddit

import (
	"fmt"
)

func (sess *Session) UserAbout(user string) (*User, error) {
	userAboutURL := fmt.Sprintf("%s/%s%s", UserURL, user, aboutEnding)

	var getUser GetUser
	dataErr := sess.GetResponse(userAboutURL, GET, nil, &getUser)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getUser.Data, nil
}

func (sess *Session) UserTrophies(user string) ([]Trophie, error) {
	userTrophiesURL := fmt.Sprintf("%s/%s%s", User2URL, user, trophiesEnding)
	var getTrophies GetTrophies
	dataErr := sess.GetResponse(userTrophiesURL, GET, nil, &getTrophies)
	if dataErr != nil {
		return nil, dataErr
	}

	var trophies []Trophie
	for _, trophieObject := range getTrophies.Data.TrophieList {
		trophies = append(trophies, trophieObject.Data)
	}

	return trophies, nil
}

func (sess *Session) UserFriends(user string) ([]FriendData, error) {
	userFriendsURL := fmt.Sprintf("%s/%s", UserFriendsBeginning, user)
	var friendData []FriendData
	dataErr := sess.GetResponse(userFriendsURL, GET, nil, friendData)
	if dataErr != nil {
		return nil, dataErr
	}

	return friendData, nil
}

func (sess *Session) UserComments(user string) ([]Comment, error) {
	userCommentsURL := fmt.Sprintf("%s/%s%s", UserURL, user, commentsEnding)
	var getCommentJSON GetCommentJSON
	dataErr := sess.GetResponse(userCommentsURL, GET, nil, &getCommentJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	comments := []Comment{}
	for _, commentStruct := range getCommentJSON.Data.Children {
		comments = append(comments, commentStruct.CommentData)
	}

	return comments, nil
}

func (sess *Session) UserPosts(user string) ([]Submission, error) {
	userPostsURL := fmt.Sprintf("%s/%s%s", UserURL, user, postsEnding)
	var getSubmission GetSubmission
	dataErr := sess.GetResponse(userPostsURL, GET, nil, &getSubmission)
	if dataErr != nil {
		return nil, dataErr
	}

	submissions := []Submission{}
	for _, submissionData := range getSubmission.Data.Children {
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
