package twitchgo

import (
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

	var getTopGamesJSON GetTopGamesJSON
	dataErr := sess.GetResponse(addParams(optionalParams, GetTopGamesURL, []string{}), "GET", false, true, nil, &getTopGamesJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getTopGamesJSON, nil
}

func (sess *Session) GetGames(ID string) (*GetGamesJSON, error) {
	getGamesParams := make(map[string]interface{})

	// Required Params
	getGamesParams["id"] = ID
	// Optional Params -> None

	var getGamesJSON GetGamesJSON
	dataErr := sess.GetResponse(addParams(getGamesParams, GetGamesURL, []string{"id", "name"}), "GET", true, true, nil, &getGamesJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getGamesJSON, nil
}
