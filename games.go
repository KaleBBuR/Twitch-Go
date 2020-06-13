package twitchgo

import (
	"encoding/json"
	"fmt"
)

func (sess *Session) GetTopGames(optionalParams map[string]interface{}) (*GetTopGamesJSON, error) {

	// Required Params -> None
	// Optional Params
	optionalParamNames := []string{"after", "before", "first"}

	for key, _ := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}
	}

	twitchRequestData := &TwitchRequest{
		URL:          addParams(optionalParams, GetTopGamesURL, []string{}),
		HttpMethod:   "GET",
		NeedOAuth2:   false,
		NeedClientID: true,
	}

	responseString, twitchRequestErr := sess.TwitchRequest(twitchRequestData, nil)
	if twitchRequestErr != nil {
		return nil, twitchRequestErr
	}

	var getTopGamesJSON GetTopGamesJSON
	json.Unmarshal([]byte(responseString), &getTopGamesJSON)

	return &getTopGamesJSON, nil
}

func (sess *Session) GetGames(ID string, name string) (*GetGamesJSON, error) {
	getGamesParams := make(map[string]interface{})

	// Required Params
	getGamesParams["id"] = ID
	getGamesParams["name"] = name
	// Optional Params -> None

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getGamesParams, GetGamesURL, []string{"id", "name"}),
		HttpMethod:   "GET",
		NeedOAuth2:   false,
		NeedClientID: true,
	}

	responseString, twitchRequestErr := sess.TwitchRequest(twitchRequestData, nil)
	if twitchRequestErr != nil {
		return nil, twitchRequestErr
	}

	var getGamesJSON GetGamesJSON
	json.Unmarshal([]byte(responseString), &getGamesJSON)

	return &getGamesJSON, nil
}
