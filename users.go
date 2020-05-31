package twitchgo

import (
	"encoding/json"
	"fmt"
)

func (oauth2 *OAuth2) getUsers(optionalParams map[string]interface{}) (*GetUsersJSON, error) {
	getUsersParams := make(map[string]interface{})

	// Required Params -> None
	// Optional Params
	optionalParamNames := []string{"id", "login"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getUsersParams[key] = value
	}

	twitchRequest := &TwitchRequest{
		URL:          addParams(getUsersParams, GetUsersURL, []string{}),
		HttpMethod:   "GET",
		NeedOAuth2:   true,
		NeedClientID: false,
	}

	responseString, responseErr := oauth2.twitchRequest(twitchRequest, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getUsersJSON GetUsersJSON
	json.Unmarshal([]byte(responseString), &getUsersJSON)

	return &getUsersJSON, nil
}

func (oauth2 *OAuth2) getUsersFollowers(optionalParams map[string]interface{}) (*GetUsersFollowsJSON, error) {
	getUsersFollowsParams := make(map[string]interface{})

	// Required Params -> None
	// Optional Params
	optionalParamNames := []string{"after", "first", "from_id", "to_id"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getUsersFollowsParams[key] = value
	}

	twitchRequest := &TwitchRequest{
		URL:          addParams(getUsersFollowsParams, GetUsersFollowsURL, []string{}),
		HttpMethod:   "GET",
		NeedOAuth2:   false,
		NeedClientID: true,
	}

	responseString, responseErr := oauth2.twitchRequest(twitchRequest, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getUsersFollowsJSON GetUsersFollowsJSON
	json.Unmarshal([]byte(responseString), &getUsersFollowsJSON)
	return &getUsersFollowsJSON, nil
}

func (oauth2 *OAuth2) updateUser(description string) (*GetUsersJSON, error) {
	getUpdateUserParams := make(map[string]interface{})

	// Required Params -> None
	// Optional Params
	getUpdateUserParams["description"] = description

	twitchRequest := &TwitchRequest{
		URL:          addParams(getUpdateUserParams, GetUsersURL, []string{}),
		HttpMethod:   "PUT",
		NeedOAuth2:   true,
		NeedClientID: false,
	}

	responseString, responseErr := oauth2.twitchRequest(twitchRequest, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getUpdateUserJSON GetUsersJSON
	json.Unmarshal([]byte(responseString), &getUpdateUserJSON)

	return &getUpdateUserJSON, nil
}

func (oauth2 *OAuth2) getUserExtensions() (*GetUserExtensionsJSON, error) {
	twitchRequest := &TwitchRequest{
		URL:          GetUserExtensionsURL,
		HttpMethod:   "GET",
		NeedOAuth2:   true,
		NeedClientID: false,
	}

	responseString, responseErr := oauth2.twitchRequest(twitchRequest, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getUserExtensionsJSON GetUserExtensionsJSON
	json.Unmarshal([]byte(responseString), &getUserExtensionsJSON)

	return &getUserExtensionsJSON, nil
}

// func (oauth2 *OAuth2) getUserActiveExtensions() {}

// func (oauth2 *OAuth2) updateUserExtensions() {}
