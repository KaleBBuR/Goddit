package goddit

func (sess *Session) MyData() (*Me, error) {
	req, RequestErr := RedditAPIRequest(GET, MeURL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	me := &Me{}
	ResponseErr := sess.RedditAPIResponse(req, me)
	if ResponseErr != nil {
		return nil, ResponseErr
	}

	return me, nil
}

func (sess *Session) MyPrefs() (*MePrefs, error) {
	req, RequestErr := RedditAPIRequest(GET, MePrefURL, nil)
	if RequestErr != nil {
		return nil, RequestErr
	}

	mePrefs := &MePrefs{}
	ResponseErr := sess.RedditAPIResponse(req, mePrefs)
	if ResponseErr != nil {
		return nil, ResponseErr
	}

	return mePrefs, nil
}

func (sess *Session) MyKarma() (*[]KarmaData, error) {
	req, RequestErr := RedditAPIRequest(GET, MeKarmaURL, nil)

	if RequestErr != nil {
		return nil, RequestErr
	}

	meKarma := &MeKarma{}
	ResponseErr := sess.RedditAPIResponse(req, meKarma)
	if ResponseErr != nil {
		return nil, ResponseErr
	}

	karmaData := []KarmaData{}
	for _, karmadatastruct := range meKarma.Data {
		karmaData = append(karmaData, karmadatastruct)
	}

	return &karmaData, nil
}

func (sess *Session) MyTrophies() (*[]Trophie, error) {
	req, RequestErr := RedditAPIRequest(GET, MeTrophieURL, nil)

	if RequestErr != nil {
		return nil, RequestErr
	}

	trophieStruct := &GetTrophies{}
	ResponseErr := sess.RedditAPIResponse(req, trophieStruct)
	if ResponseErr != nil {
		return nil, ResponseErr
	}

	trophies := []Trophie{}
	for _, trophieget := range trophieStruct.Data.TrophieList {
		trophies = append(trophies, trophieget.Data)
	}

	return &trophies, nil
}
