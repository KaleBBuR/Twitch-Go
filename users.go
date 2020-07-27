package twitchgo

import (
	"encoding/json"
	"fmt"
)

func (sess *Session) GetUsers(optionalParams map[string]interface{}) (*GetUsersJSON, error) {
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

	data, dataErr := sess.GetResponse(addParams(getUsersParams, GetUsersURL, []string{}), "GET", true, true, nil)
	if dataErr != nil {
		return nil, dataErr
	}

	var getUsersJSON GetUsersJSON
	json.Unmarshal(data, &getUsersJSON)

	return &getUsersJSON, nil
}

func (sess *Session) GetUsersFollowers(optionalParams map[string]interface{}) (*GetUsersFollowsJSON, error) {
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

	data, dataErr := sess.GetResponse(addParams(getUsersFollowsParams, GetUsersFollowsURL, []string{}), "GET", false, true, nil)
	if dataErr != nil {
		return nil, dataErr
	}

	var getUsersFollowsJSON GetUsersFollowsJSON
	json.Unmarshal(data, &getUsersFollowsJSON)
	return &getUsersFollowsJSON, nil
}

func (sess *Session) UpdateUser(description string) (*GetUsersJSON, error) {
	getUpdateUserParams := make(map[string]interface{})

	// Required Params -> None
	// Optional Params
	getUpdateUserParams["description"] = description

	data, dataErr := sess.GetResponse(addParams(getUpdateUserParams, GetUsersURL, []string{}), "PUT", true, false, nil)
	if dataErr != nil {
		return nil, dataErr
	}

	var getUpdateUserJSON GetUsersJSON
	json.Unmarshal(data, &getUpdateUserJSON)

	return &getUpdateUserJSON, nil
}

func (sess *Session) GetUserExtensions() (*GetUserExtensionsJSON, error) {
	data, dataErr := sess.GetResponse(GetUserExtensionsURL, "GET", true, false, nil)
	if dataErr != nil {
		return nil, dataErr
	}

	var getUserExtensionsJSON GetUserExtensionsJSON
	json.Unmarshal(data, &getUserExtensionsJSON)

	return &getUserExtensionsJSON, nil
}

// func (sess *Session) GetUserActiveExtensions() {}

// func (sess *Session) UpdateUserExtensions() {}
