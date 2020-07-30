package goddit

func (sess *Session) MyData() (*Me, error) {
	var me Me
	dataErr := sess.GetResponse(MeURL, GET, nil, &me)
	if dataErr != nil {
		return nil, dataErr
	}

	return &me, nil
}

func (sess *Session) MyPrefs() (*MePrefs, error) {
	var mePrefs MePrefs
	dataErr := sess.GetResponse(MePrefURL, GET, nil, &mePrefs)
	if dataErr != nil {
		return nil, dataErr
	}

	return &mePrefs, nil
}

func (sess *Session) MyKarma() ([]KarmaData, error) {
	var meKarma MeKarma
	dataErr := sess.GetResponse(MeKarmaURL, GET, nil, &meKarma)
	if dataErr != nil {
		return nil, dataErr
	}

	karmaData := []KarmaData{}
	for _, karmadatastruct := range meKarma.Data {
		karmaData = append(karmaData, karmadatastruct)
	}

	return karmaData, nil
}

func (sess *Session) MyTrophies() ([]Trophie, error) {
	var getTrophies GetTrophies
	dataErr := sess.GetResponse(MeTrophieURL, GET, nil, &getTrophies)
	if dataErr != nil {
		return nil, dataErr
	}

	trophies := []Trophie{}
	for _, trophieget := range getTrophies.Data.TrophieList {
		trophies = append(trophies, trophieget.Data)
	}

	return trophies, nil
}
