package goddit

import (
	"fmt"
)

func (sess *Session) GetSubmission(url string) (Submission, error) {
	req, RequestErr := RedditAPIRequest(GET, url, nil)
	if RequestErr != nil {
		return Submission{}, RequestErr
	}

	submissionsStruct := &GetSubmission{}
	ResponseErr := sess.RedditAPIResponse(req, &submissionsStruct)
	if ResponseErr != nil {
		return Submission{}, ResponseErr
	}

	return submissionsStruct.Data.Children[0].Data, nil
}

func (sess *Session) GetSubmissions(subreddit string, sort string, optionalParams map[string]interface{}) (*[]Submission, error) {
	sorts := []string{"hot", "new", "top", "random", "rising", "controversial"}
	baseParams := []string{"after", "before", "count", "limit", "show", "sr_detail"}
	sortParams := make(map[string][]string)
	sortParams["hot"] = append(baseParams, "g")
	sortParams["new"] = baseParams
	sortParams["random"] = []string{}
	sortParams["rising"] = baseParams
	sortParams["top"] = append(baseParams, "t")
	sortParams["controversial"] = append(baseParams, "t")

	if sort != "" {
		if !contains(sorts, sort) {
			panic(fmt.Sprintf("This is not a sorting option -> %s\n\nYour sorting options are:\nhot\nnew\ntop\nrandom\nrising\ncontroversial", sort))
		}

		if params, ok := sortParams[sort]; ok {
			for key, _ := range optionalParams {
				if !contains(params, key) {
					panic(fmt.Sprintf("This is not a valid parameter -> %s", key))
				}
			}
		}

		url := fmt.Sprintf("%s/%s/%s.json", SubredditURL, subreddit, sort)
		url = addParams(url, optionalParams)

		req, RequestErr := RedditAPIRequest(GET, url, nil)
		if RequestErr != nil {
			return nil, RequestErr
		}

		submissionsStruct := &GetSubmission{}
		ResponseErr := sess.RedditAPIResponse(req, submissionsStruct)

		if ResponseErr != nil {
			return nil, ResponseErr
		}

		submissions := []Submission{}

		for _, children := range submissionsStruct.Data.Children {
			submissions = append(submissions, children.Data)
		}

		return &submissions, nil

	} else {
		panic("You must have some kind of sort.")
	}
}

func (sess *Session) SearchSubmissions(subreddit string, optionalParams map[string]interface{}) (*[]Submission, error) {
	allOptionalParams := []string{"after", "before", "category", "count", "include_facets", "limit", "q", "restrict_sr", "show", "sr_detail", "t", "type"}
	for key, _ := range optionalParams {
		if !contains(allOptionalParams, key) {
			panic(fmt.Sprintf("Invalid key for searching submissions -> %s", key))
		}
	}

	URL := fmt.Sprintf("%s/%s/search.json", SubredditURL, subreddit)
	URL = addParams(URL, optionalParams)

	req, RequestErr := RedditAPIRequest(GET, URL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	submissionsStruct := &GetSubmission{}
	ResponseErr := sess.RedditAPIResponse(req, submissionsStruct)

	if ResponseErr != nil {
		return nil, ResponseErr
	}

	submissions := []Submission{}

	for _, submission := range submissionsStruct.Data.Children {
		submissions = append(submissions, submission.Data)
	}

	return &submissions, nil
}

/*
 * -----------------------------------------------------------------------------------------------------------------------------------------
 * Comments and Links
 * -----------------------------------------------------------------------------------------------------------------------------------------
 */

func (sess *Session) getFullID(a interface{}) string {
	switch a.(type) {
	case string:
		// If it's a string, check if it's a link in which we can get a post or comment out of.
		fmt.Printf("It's a string!")
		link := a.(string)
		sub, getSubErr := sess.GetSubmission(link)
		if getSubErr != nil {
			panic(fmt.Sprintf("Error! -> %s", getSubErr))
		}

		return sub.FullID
	case Comment:
		// Check if it's a Comment type, which we can get the full ID of
		fmt.Printf("It's a comment!")
		return a.(Comment).FullID
	case Submission:
		// Check if it's a Submission type, which we can get the full ID of
		fmt.Printf("It's a submission!")
		return a.(Submission).FullID
	default:
		panic("Must be a submission, comment, or string!")
	}
}

func (sess *Session) Hide(a interface{}) error {
	hideParams := make(map[string]interface{})
	hideParams["id"] = sess.getFullID(a)
	hideURL := addParams(hidePostURL, hideParams)

	req, RequestErr := RedditAPIRequest(POST, hideURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Lock(a interface{}) error {
	lockParams := make(map[string]interface{})
	lockParams["id"] = sess.getFullID(a)
	lockURL := addParams(lockCommentPostURL, lockParams)
	req, RequestErr := RedditAPIRequest(POST, lockURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) MarkNSFW(a interface{}) error {
	nsfwParams := make(map[string]interface{})
	nsfwParams["id"] = sess.getFullID(a)
	nsfwURL := addParams(markNsfwURL, nsfwParams)
	req, RequestErr := RedditAPIRequest(POST, nsfwURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Report(a interface{}, reportParams map[string]interface{}) error {
	possibleReportParams := []string{"additional_info", "api_type", "custom_text", "from_help_desk", "from_modmail", "modmail_conv_id", "other_reason", "reason", "rule_reason", "site_reason", "sr_name", "thing_id", "usernames"}

	for key, _ := range reportParams {
		if !contains(possibleReportParams, key) {
			panic(fmt.Sprintf("Unusuable Parameter -> %s", key))
		}
	}

	reportParams["thing_id"] = sess.getFullID(a)
	reportingURL := addParams(reportURL, reportParams)

	req, RequestErr := RedditAPIRequest(POST, reportingURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Delete(a interface{}) error {
	deleteParams := make(map[string]interface{})
	deleteParams["id"] = sess.getFullID(a)
	deleteURL := addParams(deleteCommentPostURL, deleteParams)
	req, RequestErr := RedditAPIRequest(POST, deleteURL, nil)

	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Edit(a interface{}, editParams map[string]interface{}) error {
	possibleEditParams := []string{"api_type", "return_rtjson", "richtext_json", "text"}
	for key, _ := range editParams {
		if !contains(possibleEditParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	editParams["thing_id"] = sess.getFullID(a)
	editURL := addParams(editCommentPostURL, editParams)

	req, RequestErr := RedditAPIRequest(POST, editURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Save(a interface{}, saveParams map[string]interface{}) error {
	possibleSaveParam := "category"

	for key, _ := range saveParams {
		if key != possibleSaveParam {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	saveParams["id"] = sess.getFullID(a)
	saveParamsURL := addParams(saveURL, saveParams)
	req, RequestErr := RedditAPIRequest(POST, saveParamsURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) setSubredditSticky(a interface{}, stickyParams map[string]interface{}) error {
	possibleSubStickyParams := []string{"api_type", "num", "state", "to_profile"}
	for key, _ := range stickyParams {
		if !contains(possibleSubStickyParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	stickyParams["id"] = sess.getFullID(a)

	subStickyParamURL := addParams(setSubredditStickyURL, stickyParams)
	req, RequestErr := RedditAPIRequest(POST, subStickyParamURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) SetSuggestedSort(a interface{}, setSuggSortParams map[string]interface{}) error {
	possibleSetSuggSortParams := []string{"api_type", "sort"}
	setSuggSortParams["id"] = sess.getFullID(a)

	for key, _ := range setSuggSortParams {
		if !contains(possibleSetSuggSortParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	setSuggSortParamURL := addParams(setSuggestedSortURL, setSuggSortParams)

	req, RequestErr := RedditAPIRequest(POST, setSuggSortParamURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Spolier(a interface{}) error {
	spoilerParams := make(map[string]interface{})
	spoilerParams["id"] = sess.getFullID(a)

	spoilerParamURL := addParams(spoilerURL, spoilerParams)

	req, RequestErr := RedditAPIRequest(POST, spoilerParamURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) StoreVists(links string) error {
	storeVistsParams := make(map[string]interface{})
	storeVistsParams["links"] = links

	storeVistsParamsURL := addParams(storeVisitsURL, storeVistsParams)
	req, RequestErr := RedditAPIRequest(POST, storeVistsParamsURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Post(postParams map[string]interface{}) error {
	possibleSubmitParams := []string{"ad", "api_type", "app", "collection_id", "event_end", "event_start", "event_tz", "extension", "flair_id", "flair_text", "g-recaptcha-response", "kind", "nsfw", "resubmit", "richtext_json", "sendreplies", "spoiler", "sr", "text", "title", "url", "video_poster_url"}
	for key, _ := range postParams {
		if !contains(possibleSubmitParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	postURL := addParams(submitURL, postParams)

	req, RequestErr := RedditAPIRequest(POST, postURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return RequestErr
	}

	return nil
}

func (sess *Session) Unhide(a interface{}) error {
	unhideParams := make(map[string]interface{})
	unhideParams["id"] = sess.getFullID(a)

	unhideParamURL := addParams(unhidePostURL, unhideParams)
	req, RequestErr := RedditAPIRequest(POST, unhideParamURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Unlock(a interface{}) error {
	unlockParams := make(map[string]interface{})
	unlockParams["id"] = sess.getFullID(a)

	unlockParamURL := addParams(unlockCommentPostURL, unlockParams)

	req, RequestErr := RedditAPIRequest(POST, unlockParamURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) UnmarkNSFW(a interface{}) error {
	unmarkNSFWParams := make(map[string]interface{})
	unmarkNSFWParams["id"] = sess.getFullID(a)

	unMarkNSFWURL := addParams(unmarkNsfwURL, unmarkNSFWParams)

	req, RequestErr := RedditAPIRequest(POST, unMarkNSFWURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Unsave(a interface{}) error {
	unsaveParams := make(map[string]interface{})
	unsaveParams["id"] = sess.getFullID(a)

	unsaveParamsURL := addParams(unSaveURL, unsaveParams)
	req, RequestErr := RedditAPIRequest(POST, unsaveParamsURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Unspoiler(a interface{}) error {
	unspoilerParams := make(map[string]interface{})
	unspoilerParams["id"] = sess.getFullID(a)

	unspoilerParamsURL := addParams(unspoilerURL, unspoilerParams)

	req, RequestErr := RedditAPIRequest(POST, unspoilerParamsURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}

func (sess *Session) Vote(a interface{}, vote int) error {
	voteParams := make(map[string]interface{})
	voteParams["id"] = sess.getFullID(a)
	voteParams["dir"] = vote
	voteParamURL := addParams(voteURL, voteParams)

	req, RequestErr := RedditAPIRequest(POST, voteParamURL, nil)
	if RequestErr != nil {
		return RequestErr
	}

	ResponseErr := sess.RedditAPIResponse(req, nil)
	if ResponseErr != nil {
		return ResponseErr
	}

	return nil
}
