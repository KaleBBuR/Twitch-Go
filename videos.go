package twitchgo

import (
	"encoding/json"
	"fmt"
)

func (sess *Session) GetVideos(ID string, userID string, gameID string, optionalParams map[string]interface{}) (*GetVideosJSON, error) {
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

	data, dataErr := sess.GetResponse(addParams(getVideosParams, GetVideosURL, []string{"id", "user_id", "game_id"}), "GET", false, true, nil)
	if dataErr != nil {
		return nil, dataErr
	}

	var getVideosJSON GetVideosJSON
	json.Unmarshal(data, &getVideosJSON)

	return &getVideosJSON, nil
}
