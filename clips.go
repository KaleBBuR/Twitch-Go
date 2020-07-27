package twitchgo

import (
	"encoding/json"
)

func (sess *Session) CreateClip(createClipParams map[string]interface{}) (*CreateClipJSON, error) {
	requiredCreateClipParams := []string{"broadcaster_id"}

	data, dataErr := sess.GetResponse(addParams(createClipParams, CreateGetClipURL, requiredCreateClipParams), "POST", true, true, nil)
	if dataErr != nil {
		return nil, dataErr
	}

	createClipJSON := &CreateClipJSON{}
	json.Unmarshal(data, createClipJSON)

	return createClipJSON, nil
}

func (sess *Session) GetClip(getClipParams map[string]interface{}) (*GetClipJSON, error) {
	requiredGetClipParams := []string{"broadcaster_id", "game_id", "id"}

	data, dataErr := sess.GetResponse(addParams(getClipParams, CreateGetClipURL, requiredGetClipParams), "GET", true, true, nil)
	if dataErr != nil {
		return nil, dataErr
	}

	getClipJSON := &GetClipJSON{}
	json.Unmarshal(data, getClipJSON)

	return getClipJSON, nil
}
