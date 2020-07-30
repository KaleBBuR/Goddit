package goddit

import (
	"fmt"
)

func (sess *Session) GetSubmission(url string) (*Submission, error) {
	var getSubmission GetSubmission
	dataErr := sess.GetResponse(url, GET, nil, &getSubmission)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getSubmission.Data.Children[0].Data, nil
}

func (sess *Session) GetSubmissions(subreddit string, sort string, optionalParams map[string]interface{}) ([]Submission, error) {
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
			for key := range optionalParams {
				if !contains(params, key) {
					panic(fmt.Sprintf("This is not a valid parameter -> %s", key))
				}
			}
		}

		url := fmt.Sprintf("%s/%s/%s.json", SubredditURL, subreddit, sort)
		url = addParams(url, optionalParams)

		var getSubmission GetSubmission
		dataErr := sess.GetResponse(url, GET, nil, &getSubmission)
		if dataErr != nil {
			return nil, dataErr
		}

		submissions := []Submission{}

		for _, children := range getSubmission.Data.Children {
			submissions = append(submissions, children.Data)
		}

		return submissions, nil

	} else {
		panic("You must have some kind of sort.")
	}
}

func (sess *Session) SearchSubmissions(subreddit string, optionalParams map[string]interface{}) ([]Submission, error) {
	allOptionalParams := []string{"after", "before", "category", "count", "include_facets", "limit", "q", "restrict_sr", "show", "sr_detail", "t", "type"}
	for key := range optionalParams {
		if !contains(allOptionalParams, key) {
			panic(fmt.Sprintf("Invalid key for searching submissions -> %s", key))
		}
	}

	URL := fmt.Sprintf("%s/%s/search.json", SubredditURL, subreddit)
	URL = addParams(URL, optionalParams)
	var getSubmission GetSubmission
	dataErr := sess.GetResponse(URL, GET, nil, &getSubmission)
	if dataErr != nil {
		return nil, dataErr
	}
	submissions := []Submission{}

	for _, submission := range getSubmission.Data.Children {
		submissions = append(submissions, submission.Data)
	}

	return submissions, nil
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
	dataErr := sess.GetResponse(hideURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Lock(a interface{}) error {
	lockParams := make(map[string]interface{})
	lockParams["id"] = sess.getFullID(a)
	lockURL := addParams(lockCommentPostURL, lockParams)
	dataErr := sess.GetResponse(lockURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) MarkNSFW(a interface{}) error {
	nsfwParams := make(map[string]interface{})
	nsfwParams["id"] = sess.getFullID(a)
	nsfwURL := addParams(markNsfwURL, nsfwParams)
	dataErr := sess.GetResponse(nsfwURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Report(a interface{}, reportParams map[string]interface{}) error {
	possibleReportParams := []string{"additional_info", "api_type", "custom_text", "from_help_desk", "from_modmail", "modmail_conv_id", "other_reason", "reason", "rule_reason", "site_reason", "sr_name", "thing_id", "usernames"}

	for key := range reportParams {
		if !contains(possibleReportParams, key) {
			panic(fmt.Sprintf("Unusuable Parameter -> %s", key))
		}
	}

	reportParams["thing_id"] = sess.getFullID(a)
	reportingURL := addParams(reportURL, reportParams)

	dataErr := sess.GetResponse(reportingURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Delete(a interface{}) error {
	deleteParams := make(map[string]interface{})
	deleteParams["id"] = sess.getFullID(a)
	deleteURL := addParams(deleteCommentPostURL, deleteParams)

	dataErr := sess.GetResponse(deleteURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Edit(a interface{}, editParams map[string]interface{}) error {
	possibleEditParams := []string{"api_type", "return_rtjson", "richtext_json", "text"}
	for key := range editParams {
		if !contains(possibleEditParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	editParams["thing_id"] = sess.getFullID(a)
	editURL := addParams(editCommentPostURL, editParams)

	dataErr := sess.GetResponse(editURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Save(a interface{}, saveParams map[string]interface{}) error {
	possibleSaveParam := "category"

	for key := range saveParams {
		if key != possibleSaveParam {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	saveParams["id"] = sess.getFullID(a)
	saveParamsURL := addParams(saveURL, saveParams)
	dataErr := sess.GetResponse(saveParamsURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) SetSubredditSticky(a interface{}, stickyParams map[string]interface{}) error {
	possibleSubStickyParams := []string{"api_type", "num", "state", "to_profile"}
	for key := range stickyParams {
		if !contains(possibleSubStickyParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	stickyParams["id"] = sess.getFullID(a)
	subStickyParamURL := addParams(setSubredditStickyURL, stickyParams)
	dataErr := sess.GetResponse(subStickyParamURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) SetSuggestedSort(a interface{}, setSuggSortParams map[string]interface{}) error {
	possibleSetSuggSortParams := []string{"api_type", "sort"}
	setSuggSortParams["id"] = sess.getFullID(a)

	for key := range setSuggSortParams {
		if !contains(possibleSetSuggSortParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	setSuggSortParamURL := addParams(setSuggestedSortURL, setSuggSortParams)
	dataErr := sess.GetResponse(setSuggSortParamURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Spolier(a interface{}) error {
	spoilerParams := make(map[string]interface{})
	spoilerParams["id"] = sess.getFullID(a)

	spoilerParamURL := addParams(spoilerURL, spoilerParams)

	dataErr := sess.GetResponse(spoilerParamURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) StoreVists(links string) error {
	storeVistsParams := make(map[string]interface{})
	storeVistsParams["links"] = links

	storeVistsParamsURL := addParams(storeVisitsURL, storeVistsParams)

	dataErr := sess.GetResponse(storeVistsParamsURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Post(postParams map[string]interface{}) error {
	possibleSubmitParams := []string{"ad", "api_type", "app", "collection_id", "event_end", "event_start", "event_tz", "extension", "flair_id", "flair_text", "g-recaptcha-response", "kind", "nsfw", "resubmit", "richtext_json", "sendreplies", "spoiler", "sr", "text", "title", "url", "video_poster_url"}
	for key := range postParams {
		if !contains(possibleSubmitParams, key) {
			panic(fmt.Sprintf("Unusable Parameter -> %s", key))
		}
	}

	postURL := addParams(submitURL, postParams)

	dataErr := sess.GetResponse(postURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Unhide(a interface{}) error {
	unhideParams := make(map[string]interface{})
	unhideParams["id"] = sess.getFullID(a)

	unhideParamURL := addParams(unhidePostURL, unhideParams)

	dataErr := sess.GetResponse(unhideParamURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Unlock(a interface{}) error {
	unlockParams := make(map[string]interface{})
	unlockParams["id"] = sess.getFullID(a)

	unlockParamURL := addParams(unlockCommentPostURL, unlockParams)

	dataErr := sess.GetResponse(unlockParamURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) UnmarkNSFW(a interface{}) error {
	unmarkNSFWParams := make(map[string]interface{})
	unmarkNSFWParams["id"] = sess.getFullID(a)

	unMarkNSFWURL := addParams(unmarkNsfwURL, unmarkNSFWParams)

	dataErr := sess.GetResponse(unMarkNSFWURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Unsave(a interface{}) error {
	unsaveParams := make(map[string]interface{})
	unsaveParams["id"] = sess.getFullID(a)

	unsaveParamsURL := addParams(unSaveURL, unsaveParams)

	dataErr := sess.GetResponse(unsaveParamsURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Unspoiler(a interface{}) error {
	unspoilerParams := make(map[string]interface{})
	unspoilerParams["id"] = sess.getFullID(a)

	unspoilerParamsURL := addParams(unspoilerURL, unspoilerParams)

	dataErr := sess.GetResponse(unspoilerParamsURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}

func (sess *Session) Vote(a interface{}, vote int) error {
	voteParams := make(map[string]interface{})
	voteParams["id"] = sess.getFullID(a)
	voteParams["dir"] = vote
	voteParamURL := addParams(voteURL, voteParams)

	dataErr := sess.GetResponse(voteParamURL, POST, nil, nil)
	if dataErr != nil {
		return dataErr
	}

	return nil
}
