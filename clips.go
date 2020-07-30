package twitchgo

func (sess *Session) CreateClip(createClipParams map[string]interface{}) (*CreateClipJSON, error) {
	requiredCreateClipParams := []string{"broadcaster_id"}

	var createClipJSON CreateClipJSON
	dataErr := sess.GetResponse(addParams(createClipParams, CreateGetClipURL, requiredCreateClipParams), "POST", true, true, nil, &createClipJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &createClipJSON, nil
}

func (sess *Session) GetClip(getClipParams map[string]interface{}) (*GetClipJSON, error) {
	requiredGetClipParams := []string{"broadcaster_id", "game_id", "id"}

	var getClipJSON GetClipJSON
	dataErr := sess.GetResponse(addParams(getClipParams, CreateGetClipURL, requiredGetClipParams), "GET", true, true, nil, &getClipJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getClipJSON, nil
}
