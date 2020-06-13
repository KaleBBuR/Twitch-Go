package twitchgo

import (
	"encoding/json"
)

func (sess *Session) CreateClip(createClipParams map[string]interface{}) (*CreateClipJSON, error) {
	requiredCreateClipParams := []string{"broadcaster_id"}

	twitchRequestData := &TwitchRequest{
		URL:          addParams(createClipParams, CreateGetClipURL, requiredCreateClipParams),
		HttpMethod:   "POST",
		NeedOAuth2:   true,
		NeedClientID: true,
	}

	responseString, responseErr := sess.TwitchRequest(twitchRequestData, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	createClipJSON := &CreateClipJSON{}
	json.Unmarshal([]byte(responseString), createClipJSON)

	return createClipJSON, nil
}

func (sess *Session) GetClip(getClipParams map[string]interface{}) (*GetClipJSON, error) {
	requiredGetClipParams := []string{"broadcaster_id", "game_id", "id"}

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getClipParams, CreateGetClipURL, requiredGetClipParams),
		HttpMethod:   "GET",
		NeedOAuth2:   true,
		NeedClientID: true,
	}

	responseString, responseErr := sess.TwitchRequest(twitchRequestData, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	getClipJSON := &GetClipJSON{}
	json.Unmarshal([]byte(responseString), getClipJSON)

	return getClipJSON, nil
}
