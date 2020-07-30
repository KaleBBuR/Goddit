package goddit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Moderator struct {
	Name  string `json:"name"`
	Flair string `json:"author_flair_text"`
	ID    string `json:"id"`
}

type GetModerators struct {
	Data struct {
		Mods []Moderator `json:"children"`
	} `json:"data"`
}

type ModLog struct {
	Description       string `json:"description"`
	TargetCommentBody string `json:"target_body"`
	ModID             string `json:"mod_id"`
	Subreddit         string `json:"subredddit"`
	TargetTitle       string `json:"target_title"`
	TargetPermalink   string `json:"target_permalink"`
	Details           string `json:"details"`
	Action            string `json:"action"`
	TargetAuthor      string `json:"target_author"`
	TargetFullname    string `json:"target_fullname"`
	ID                string `json:"id"`
	Mod               string `json:"mod"`
}

type GetModLog struct {
	Data struct {
		Children []struct {
			Data ModLog `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type User struct {
	IsEmployee       bool   `json:"is_employee"`
	Icon             string `json:"icon_img"`
	Name             string `json:"name"`
	IsFriend         bool   `json:"is_friend"`
	HasSubsribed     bool   `json:"has_subscribed"`
	CreatedUTC       int64  `json:"created_utc"`
	PostKarma        int64  `json:"link_karma"`
	CommentKarma     int64  `json:"comment_karma"`
	IsGold           bool   `json:"is_gold"`
	IsMod            bool   `json:"is_mod"`
	Verified         string `json:"verified"`
	HasVerifiedEmail bool   `json:"has_verified_email"`
	ID               string `json:"id"`
}

type GetUser struct {
	Data User `json:"data"`
}

type GetSubmission struct {
	Data struct {
		Children []struct {
			Data Submission `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type Submission struct {
	Subreddit   string `json:"subreddit"`
	Saved       bool   `json:"saved"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	NumComments int    `json:"num_comments"`
	Permalink   string `json:"permalink"`
	URL         string `json:"url"`
	IsVideo     bool   `json:"is_video"`
	UTC         int    `json:"created_utc"`
	Awards      []struct {
		Name   string `json:"name"`
		Amount string `json:"count"`
	} `json:"all_awardings"`
	Removed_by    string `json:"removed_by"`
	ID            int    `json:"id"`
	Upvotes       int    `json:"ups"`
	FullID        string `json:"name"`
	Total_rewards int    `json:"total_awards_received"`
	Edited        bool   `json:"edited"`
	Quarantined   bool   `json:"quarantine"`
	NSFW          bool   `json:"over_18"`
	Pinned        bool   `json:"pinned"`
	Locked        bool   `json:"locked"`
	Spolier       bool   `json:"spolier"`
	Stickied      bool   `json:"stickied"`
	Hidden        bool   `json:"hidden"`
	Upvote_ratio  int    `json:"upvote_ration"`
	Thumbnail     string `json:"thumbnail"`
}

type Subreddit struct {
	Name        string `json:"display_name"`
	Title       string `json:"title"`
	Thumbnail   string `json:"header_img"`
	OnlineUsers int64  `json:"active_user_count"`
	Members     int64  `json:"subscribers"`
	FullID      string `json:"name"`
	Quarantine  bool   `json:"quarantine"`
	Description string `json:"public_description"`
	Gifs        bool   `json:"allow_videogifs"`
	Videos      bool   `json:"allow_videos"`
	ID          string `json:"id"`
	NSFW        bool   `json:"over18"`
	Images      bool   `json:"allow_images"`
	Commenting  bool   `json:"restrict_commenting"`
	Lang        string `json:"lang"`
	UTC         string `json:"utc"`
}

type Rule struct {
	Kind            string `json:"kind"`
	Description     string `json:"description"`
	ShortName       string `json:"short_name"`
	ViolationReason string `json:"violation_reason"`
	CreatedUTC      int64  `json:"created_utc"`
	Priority        int64  `json:"priority"`
}

type GetSubredditData struct {
	Data Subreddit `json:"data"`
}

type GetRules struct {
	Rules []Rule `json:"rules"`
}

type Me struct {
	Employee               bool   `json:"is_employee"`
	NoProfanity            bool   `json:"pref_no_profanity"`
	HasExteralAccount      bool   `json:"has_external_account"`
	IsSponser              bool   `json:"is_sponser"`
	HasGoldSubscription    bool   `json:"has_gold_subscription"`
	NumFriends             int64  `json:"num_friends"`
	Verified               bool   `json:"verified"`
	Coins                  int64  `json:"coins"`
	HasPaypalSubscription  bool   `json:"has_paypal_subscription"`
	HasPremiumSubscription bool   `json:"has_subscribed_to_premium"`
	ID                     int    `json:"id"`
	CanCreateSubreddit     bool   `json:"can_create_subreddit"`
	NSFW                   bool   `json:"over_18"`
	IsGold                 bool   `json:"is_gold"`
	IsMod                  bool   `json:"is_mod"`
	VerifiedEmail          bool   `json:"jas_verified_email"`
	IsSuspended            bool   `json:"is_suspended"`
	Icon                   string `json:"icon_img"`
	HasModMail             bool   `json:"has_mod_mail"`
	ClientID               string `json:"oauth_client_id"`
	PostKarma              int64  `json:"link_karma"`
	InboxCount             int64  `json:"inbox_count"`
	Name                   string `json:"name"`
	CreatedUTC             int64  `json:"created_utc"`
	Created                int64  `json:"created"`
	CommentKarma           int64  `json:"comment_karma"`
}

type MeBlocked struct {
}

type MeFriends struct {
}

type MePrefs struct {
	ThreadedMessages      bool   `json:"threaded_messages"`
	HideDownvotes         bool   `json:"hide_downs"`
	LabelNSFW             bool   `json:"label_nsfw"`
	VideoAutoplay         bool   `json:"video_autoplay"`
	ThirdPartySiteContent bool   `json:"third_party_site_data_personalized_content"`
	ShowLinkFlair         bool   `json:"show_link_flair"`
	ShowTrending          bool   `json:"show_trending"`
	SendWelcomeMessages   bool   `json:"send_welcome_messages"`
	PrivateFeeds          bool   `json:"private_feeds"`
	MonitorMentions       bool   `json:"monitor_mentions"`
	ShowAvatar            bool   `json:"show_snoovatar"`
	NSFW                  bool   `json:"over_18"`
	EmailMessages         bool   `json:"email_messages"`
	LiveOrangeReds        bool   `json:"live_orangereds"`
	EnableDefaultThemes   bool   `json:"enabled_default_themes"`
	Language              string `json:"lang"`
	ThirdPartyAds         bool   `json:"third_part_data_personalized_ads"`
	AllowClickTracking    bool   `json:"allow_clicktracking"`
	HideFromRobots        bool   `json:"hide_from_robots"`
	ShowTwitter           bool   `json:"show_twitter"`
	LowestPostScore       int64  `json:"min_link_score"`
	Nightmode             bool   `json:"nightmode"`
	ThirdPartySiteAds     bool   `json:"third_party_site_data_personalized_ads"`
	LowestCommentScore    int64  `json:"min_comment_score"`
	PublicVotes           bool   `json:"public_votes"`
	ShowFlair             bool   `json:"show_flair"`
	MarkMessagesRead      bool   `json:"mark_messages_read"`
	SearchNSFW            bool   `json:"search_include_over_18"`
	NoProfanity           bool   `json:"no_profanity"`
	HideAds               bool   `json:"hide_ads"`
	NumSites              int64  `json:"numsites"`
	NumComments           int64  `json:"num_comments"`
	ShowGoldExpiration    bool   `json:"show_gold_expiration"`
	HighlightNewComments  bool   `json:"highlight_new_comments"`
	EmailUnsubscribeAll   bool   `json:"email_unsubscribe_all"`
	DefaultCommentSort    string `json:"default_comment_sort"`
	AcceptPMS             string `json:"accept_pms"`
}

type MeKarma struct {
	Data []KarmaData `json:"data"`
}

type KarmaData struct {
	Subreddit    string `json:"sr"`
	CommentKarma int64  `json:"comment_karma"`
	PostKarma    int64  `json:"link_karma"`
}

type HttpErrorJSON struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type Trophie struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	AwardID     string `json:"award_id"`
	ID          string `json:"id"`
	Description string `json:"description"`
}

type GetTrophies struct {
	Data struct {
		TrophieList []struct {
			Data Trophie `json:"data"`
		} `json:"trophies"`
	} `json:"data"`
}

type FriendData struct {
	Date  int64  `json:"date"`
	RelID string `json:"rel_id"`
	Name  string `json:"name"`
	ID    string `json:"id"`
}

type GetPostedComment struct {
	JSON struct {
		Errors []string `json:"errors"`
		Data   struct {
			Things []struct {
				Kind        string  `json:"kind"`
				CommentData Comment `json:"data"`
			} `json:"things"`
		} `json:"data"`
	} `json:"json"`
}

type GetCommentJSON struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Dist     int    `json:"dist"`
		Children []struct {
			CommentData Comment `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type Comment struct {
	AwardsAmount     int64          `json:"total_awards_received"`
	Upvotes          int64          `json:"ups"`
	LinkID           string         `json:"link_id"`
	Replies          GetCommentJSON `json:"replies"`
	Saved            bool           `json:"saved"`
	ID               string         `json:"id"`
	Archived         bool           `json:"archived"`
	Author           string         `json:"author"`
	CanModPost       bool           `json:"can_mod_post"`
	SendReplies      bool           `json:"send_replies"`
	AuthorFullname   string         `json:"author_fullname"`
	SubredditID      string         `json:"subreddit_id"`
	Content          string         `json:"body"`
	Edited           bool           `json:"edited"`
	Stickied         bool           `json:"stickied"`
	AuthorPremium    bool           `json:"author_premium"`
	Subreddit        string         `json:"subreddit"`
	ScoreHidden      bool           `json:"score_hidden"`
	Permalink        string         `json:"permalink"`
	Locked           bool           `json:"locked"`
	FullID           string         `json:"name"`
	Created          int64          `json:"created"`
	CreatedUTC       int64          `json:"created_utc"`
	Controversiality int64          `json:"controversiality"`
	Depth            int64          `json:"depth"`
}

type Session struct {
	username     string
	password     string
	clientID     string
	clientSecret string
	useragent    string
	accesstoken  string
	expiretime   float64
}

func RedditAPIRequest(method string, url string, post_data []byte) (*http.Request, error) {
	var request *http.Request
	var err error
	if post_data != nil {
		request, err = http.NewRequest(method, url, bytes.NewBuffer(post_data))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("HTTP Request error. URL: %s\nMethod:%s\nHost:%s\n", request.URL, request.Method, request.Host))
		}
	} else {
		request, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("HTTP Request error. URL: %s\nMethod:%s\nHost:%s\n", request.URL, request.Method, request.Host))
		}
	}

	return request, nil
}

func (sess *Session) RedditAPIResponse(request *http.Request, a interface{}) error {
	request.Header.Add("Authorization", "bearer "+sess.accesstoken)
	request.Header.Add("User-Agent", sess.useragent)

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		bodyString := string(body)

		var HTTPErrorJSON HttpErrorJSON
		json.Unmarshal([]byte(bodyString), &HTTPErrorJSON)

		if resp.StatusCode == 403 || resp.StatusCode == 404 {
			return errors.New(fmt.Sprintf("HTTP ERROR\nStatus Code: %d\nError: %s\nReason: %s\nMessage: %s", resp.StatusCode, resp.Status, HTTPErrorJSON.Reason, HTTPErrorJSON.Message))
		} else {
			return errors.New(fmt.Sprintf("HTTP ERROR\nStatus Code: %d\nError: %s\n", resp.StatusCode, resp.Status))
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if JSONerr := json.Unmarshal(body, a); JSONerr != nil && a != nil {
		return err
	}

	return nil
}

func checkSpace(text string) string {
	return strings.Replace(text, " ", "%20", -1)
}

func (s *Submission) GetLink() string {
	return normalURL + s.Permalink
}

func contains(strArr []string, key string) bool {
	for _, item := range strArr {
		if item == key {
			return true
		}
	}

	return false
}

func printJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "	")
	return out.Bytes(), err
}

func addParams(url string, params map[string]interface{}) string {
	url += "?"
	for key, value := range params {
		switch value.(type) {
		case string:
			value = value.(string)
			if value != "" {
				url = url + fmt.Sprintf("%s=%s&", key, value)
			}
		case int:
			url = url + fmt.Sprintf("%s=%d&", key, value.(int))
		case bool:
			url = url + fmt.Sprintf("%s=%t&", key, value.(bool))
		}
	}

	url = strings.TrimSuffix(url, "&")
	return url
}

func (sess *Session) GetResponse(url string, method string, body []byte, a interface{}) error {
	req, RequestErr := RedditAPIRequest(method, url, body)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, a)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

const (
	baseURL   = "https://oauth.reddit.com"
	normalURL = "https://www.reddit.com"

	// Reddit API links towards the user
	MeURL        = baseURL + "/api/v1/me"
	MePrefURL    = baseURL + "/api/v1/me/prefs"
	MeKarmaURL   = baseURL + "/api/v1/me/karma"
	MeTrophieURL = baseURL + "/api/v1/me/trophies"

	// Reddit API link towards subreddit
	SubredditURL = baseURL + "/r"

	// Links & Comments
	commentURL            = baseURL + "/api/comment"
	deleteCommentPostURL  = baseURL + "/api/del"
	editCommentPostURL    = baseURL + "/api/editusertext"
	postEventTimeURL      = baseURL + "/api/event_post_time"
	followPostURL         = baseURL + "/api/follow_post"
	hidePostURL           = baseURL + "/api/hide"
	unhidePostURL         = baseURL + "/api/unhide"
	infoURL               = "/api/info"
	lockCommentPostURL    = baseURL + "/api/lock"
	unlockCommentPostURL  = baseURL + "/api/unlock"
	markNsfwURL           = baseURL + "/api/marknsfw"
	unmarkNsfwURL         = baseURL + "/api/unmarknsfw"
	moreChildrenURL       = baseURL + "/api/morechildren"
	reportURL             = baseURL + "/api/report"
	reportAwardURL        = baseURL + "/api/report_award"
	saveURL               = baseURL + "/api/save"
	unSaveURL             = baseURL + "/api/unsave"
	savedCategoriesURL    = baseURL + "/api/saved_categories"
	sendRepliesURL        = baseURL + "/api/saved_categories"
	setContestModeURL     = baseURL + "/api/set_contest_mode"
	setSubredditStickyURL = baseURL + "/api/set_subreddit_sticky"
	setSuggestedSortURL   = baseURL + "/api/set_suggested_sort"
	spoilerURL            = baseURL + "/api/spoiler"
	unspoilerURL          = baseURL + "/api/unspoiler"
	storeVisitsURL        = baseURL + "/api/store_visits"
	submitURL             = baseURL + "/api/submit"
	voteURL               = baseURL + "/api/vote"

	// Subreddit
	aboutEnding      = "/about"
	aboutRulesEnding = aboutEnding + "/rules"
	aboutModsEnding  = aboutEnding + "/moderators"
	subscribeURL     = baseURL + "/api/subscribe"

	// Moderation
	modLogEnding    = aboutEnding + "/log"
	acceptModEnding = "/api/accept_moderator_invite"
	approveURL      = baseURL + "/api/approve"
	removeURL       = baseURL + "/api/remove"
	showCommentURL  = baseURL + "/api/show_comment"

	// User
	UserURL              = baseURL + "/user"
	User2URL             = baseURL + "/api/v1/user"
	UserFriendsBeginning = baseURL + "/api/v1/me/friends"
	trophiesEnding       = "/trophies"
	commentsEnding       = "/comments"
	postsEnding          = "/submitted"

	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"

	UpVote     = 1
	DownVote   = -1
	RemoveVote = 0
)
