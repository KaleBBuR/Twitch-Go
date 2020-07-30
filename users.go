package twitchgo

import (
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

	var getUsersJSON GetUsersJSON
	dataErr := sess.GetResponse(addParams(getUsersParams, GetUsersURL, []string{}), "GET", true, true, nil, &getUsersJSON)
	if dataErr != nil {
		return nil, dataErr
	}

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

	var getUsersFollowsJSON GetUsersFollowsJSON
	dataErr := sess.GetResponse(addParams(getUsersFollowsParams, GetUsersFollowsURL, []string{}), "GET", false, true, nil, &getUsersFollowsJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getUsersFollowsJSON, nil
}

func (sess *Session) UpdateUser(description string) (*GetUsersJSON, error) {
	getUpdateUserParams := make(map[string]interface{})

	// Required Params -> None
	// Optional Params
	getUpdateUserParams["description"] = description

	var getUpdateUserJSON GetUsersJSON
	dataErr := sess.GetResponse(addParams(getUpdateUserParams, GetUsersURL, []string{}), "PUT", true, false, nil, &getUpdateUserJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getUpdateUserJSON, nil
}

func (sess *Session) GetUserExtensions() (*GetUserExtensionsJSON, error) {
	var getUserExtensionsJSON GetUserExtensionsJSON
	dataErr := sess.GetResponse(GetUserExtensionsURL, "GET", true, false, nil, &getUserExtensionsJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getUserExtensionsJSON, nil
}

// func (sess *Session) GetUserActiveExtensions() {}

// func (sess *Session) UpdateUserExtensions() {}
