package twitchgo

import (
	"encoding/json"
	"fmt"
)

func (sess *Session) GetTopGames(optionalParams map[string]interface{}) (*GetTopGamesJSON, error) {

	// Required Params -> None
	// Optional Params
	optionalParamNames := []string{"after", "before", "first"}

	for key := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}
	}

	data, dataErr := sess.GetResponse(addParams(optionalParams, GetTopGamesURL, []string{}), "GET", false, true, nil)
	if dataErr != nil {
		return nil, dataErr
	}

	var getTopGamesJSON GetTopGamesJSON
	json.Unmarshal(data, &getTopGamesJSON)

	return &getTopGamesJSON, nil
}

func (sess *Session) GetGames(ID string, name string) (*GetGamesJSON, error) {
	getGamesParams := make(map[string]interface{})

	// Required Params
	getGamesParams["id"] = ID
	getGamesParams["name"] = name
	// Optional Params -> None

	data, dataErr := sess.GetResponse(addParams(getGamesParams, GetGamesURL, []string{"id", "name"}), "GET", false, true, nil)
	if dataErr != nil {
		return nil, dataErr
	}

	var getGamesJSON GetGamesJSON
	json.Unmarshal(data, &getGamesJSON)

	return &getGamesJSON, nil
}
