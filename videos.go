package twitchgo

import (
	"encoding/json"
	"fmt"
)

func (oauth2 *OAuth2) getVideos(ID string, userID string, gameID string, optionalParams map[string]interface{}) (*GetVideosJSON, error) {
	getVideosParams := make(map[string]interface{})

	// Required Params
	getVideosParams["id"] = ID
	getVideosParams["user_id"] = userID
	getVideosParams["game_id"] = gameID

	// Optional Params
	optionalParamNames := []string{"after", "before", "first", "language", "period", "sort", "type"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getVideosParams[key] = value
	}

	twitchRequest := &TwitchRequest{
		URL:          addParams(getVideosParams, GetVideosURL, []string{"id", "user_id", "game_id"}),
		HttpMethod:   "GET",
		NeedOAuth2:   false,
		NeedClientID: true,
	}

	responseString, responseErr := oauth2.twitchRequest(twitchRequest, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getVideosJSON GetVideosJSON
	json.Unmarshal([]byte(responseString), &getVideosJSON)

	return &getVideosJSON, nil
}
