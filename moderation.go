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

	twitchRequestData := &TwitchRequest{
		URL:          addParams(postCheckAutoModStatusParams, CheckAutoModStatusURL, []string{"broadcaster_id"}),
		HttpMethod:   "POST",
		NeedOAuth2:   false,
		NeedClientID: false,
	}

	autoModStatusBody, jsonErr := json.Marshal(AutoModStatusFullBody{Data: body})
	if jsonErr != nil {
		return nil, jsonErr
	}

	responseString, responseErr := sess.TwitchRequest(twitchRequestData, autoModStatusBody)
	if responseErr != nil {
		return nil, responseErr
	}

	var postAutoModStatusJSON PostAutoModStatusJSON
	json.Unmarshal([]byte(responseString), &postAutoModStatusJSON)

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

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getBannedEventsParams, GetBannedEventsURL, []string{"broadcaster_id"}),
		HttpMethod:   "GET",
		NeedOAuth2:   true,
		NeedClientID: false,
	}

	responseString, responseErr := sess.TwitchRequest(twitchRequestData, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getBannedEventsJSON GetBannedEventsJSON
	json.Unmarshal([]byte(responseString), &getBannedEventsJSON)

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

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getBannedUsersParams, GetBannedUsersURL, []string{"broadcaster_id"}),
		HttpMethod:   "GET",
		NeedOAuth2:   true,
		NeedClientID: false,
	}

	responseString, responseErr := sess.TwitchRequest(twitchRequestData, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getBannedUsersJSON GetBannedUsersJSON
	json.Unmarshal([]byte(responseString), &getBannedUsersJSON)

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

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getModsParams, GetModsURL, []string{"broadcaster_id"}),
		HttpMethod:   "GET",
		NeedOAuth2:   true,
		NeedClientID: false,
	}

	responseString, responseErr := sess.TwitchRequest(twitchRequestData, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getModsJSON GetModsJSON
	json.Unmarshal([]byte(responseString), &getModsJSON)

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

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getModEventsParams, GetModEventsURL, []string{"broadcaster_id"}),
		HttpMethod:   "GET",
		NeedOAuth2:   true,
		NeedClientID: false,
	}

	responseString, responseErr := sess.TwitchRequest(twitchRequestData, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getModEventsJSON GetModEventsJSON
	json.Unmarshal([]byte(responseString), &getModEventsJSON)

	return &getModEventsJSON, nil
}
