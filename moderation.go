package twitchgo

import (
	"encoding/json"
	"fmt"
)

func (sess *Session) CheckAutoModStatus(broadcasterID string, body []AutoModStatusBody) (*PostAutoModStatusJSON, error) {
	postCheckAutoModStatusParams := make(map[string]interface{})

	// Required Params
	postCheckAutoModStatusParams["broadcaster_id"] = broadcasterID
	// Optional Params -> None

	autoModStatusBody, jsonErr := json.Marshal(AutoModStatusFullBody{Data: body})
	if jsonErr != nil {
		return nil, jsonErr
	}

	var postAutoModStatusJSON PostAutoModStatusJSON
	dataErr := sess.GetResponse(addParams(postCheckAutoModStatusParams, CheckAutoModStatusURL, []string{"broadcaster_id"}), "POST", false, false, autoModStatusBody, &postAutoModStatusJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &postAutoModStatusJSON, nil
}

func (sess *Session) GetBannedEvents(broadcasterID string, optionalParams map[string]interface{}) (*GetBannedEventsJSON, error) {
	getBannedEventsParams := make(map[string]interface{})

	// Required Params
	getBannedEventsParams["broadcaster_id"] = broadcasterID

	// Optional Params
	optionalParamNames := []string{"user_id", "after", "before"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getBannedEventsParams[key] = value
	}

	var getBannedEventsJSON GetBannedEventsJSON
	dataErr := sess.GetResponse(addParams(getBannedEventsParams, GetBannedEventsURL, []string{"broadcaster_id"}), "GET", true, false, nil, &getBannedEventsJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getBannedEventsJSON, nil
}

func (sess *Session) GetBannedUsers(broadcasterID string, optionalParams map[string]interface{}) (*GetBannedUsersJSON, error) {
	getBannedUsersParams := make(map[string]interface{})

	// Required Params
	getBannedUsersParams["broadcaster_id"] = broadcasterID

	// Optional Params
	optionalParamNames := []string{"user_id", "after", "before"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getBannedUsersParams[key] = value
	}

	var getBannedUsersJSON GetBannedUsersJSON
	dataErr := sess.GetResponse(addParams(getBannedUsersParams, GetBannedUsersURL, []string{"broadcaster_id"}), "GET", true, false, nil, &getBannedUsersJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getBannedUsersJSON, nil
}

func (sess *Session) GetMods(broadcasterID string, optionalParams map[string]interface{}) (*GetModsJSON, error) {
	getModsParams := make(map[string]interface{})

	// Required Params
	getModsParams["broadcaster_id"] = broadcasterID

	// Optional Params
	optionalParamNames := []string{"user_id", "after"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getModsParams[key] = value
	}

	var getModsJSON GetModsJSON
	dataErr := sess.GetResponse(addParams(getModsParams, GetModsURL, []string{"broadcaster_id"}), "GET", true, false, nil, &getModsJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getModsJSON, nil
}

func (sess *Session) GetModEvents(broadcasterID string, optionalParams map[string]interface{}) (*GetModEventsJSON, error) {
	getModEventsParams := make(map[string]interface{})

	// Required Params
	getModEventsParams["broadcaster_id"] = broadcasterID

	// Optional Params
	optionalParamNames := []string{"user_id"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getModEventsParams[key] = value
	}

	var getModEventsJSON GetModEventsJSON
	dataErr := sess.GetResponse(addParams(getModEventsParams, GetModEventsURL, []string{"broadcaster_id"}), "GET", true, false, nil, &getModEventsJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getModEventsJSON, nil
}
