package twitchgo

import (
	"encoding/json"
	"fmt"
)

func (oauth2 *OAuth2) createStreamMarker(userID string, optionalParams map[string]interface{}) (*CreateStreamMarkerJSON, error) {
	createStreamMarkerParams := make(map[string]interface{})

	// Required Params
	createStreamMarkerParams["user_id"] = userID

	// Optional Params
	optionalParamNames := []string{"description"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		createStreamMarkerParams[key] = value
	}

	twitchRequestData := &TwitchRequest{
		URL:          addParams(createStreamMarkerParams, CreateStreamMarkerURL, []string{"user_id"}),
		HttpMethod:   "POST",
		NeedOAuth2:   true,
		NeedClientID: false,
	}

	responseString, twitchRequestErr := oauth2.twitchRequest(twitchRequestData, nil)
	if twitchRequestErr != nil {
		return nil, twitchRequestErr
	}

	var createStreamMarkerJSON CreateStreamMarkerJSON
	json.Unmarshal([]byte(responseString), &createStreamMarkerJSON)

	return &createStreamMarkerJSON, nil
}

func (oauth2 *OAuth2) getStreams(optionalParams map[string]interface{}) (*GetStreamsJSON, error) {
	getStreamsParams := make(map[string]interface{})

	// Required Params -> None
	// Optional Params
	optionalParamNames := []string{"after", "before", "first", "game_id", "language", "user_id", "user_login"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getStreamsParams[key] = value
	}

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getStreamsParams, GetStreamsURL, []string{}),
		HttpMethod:   "GET",
		NeedOAuth2:   true,
		NeedClientID: true,
	}

	responseString, twitchRequestErr := oauth2.twitchRequest(twitchRequestData, nil)
	if twitchRequestErr != nil {
		return nil, twitchRequestErr
	}

	var getStreamsJSON GetStreamsJSON
	json.Unmarshal([]byte(responseString), &getStreamsJSON)

	return &getStreamsJSON, nil
}

func (oauth2 *OAuth2) getStreamMarkers(userID string, videoID string, optionalParams map[string]interface{}) (*GetStreamMarkersJSON, error) {
	getStreamMarkersParams := make(map[string]interface{})

	// Required Params
	getStreamMarkersParams["user_id"] = userID
	getStreamMarkersParams["video_id"] = videoID

	// Optional Params
	optionalParamNames := []string{"after", "before", "first"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getStreamMarkersParams[key] = value
	}

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getStreamMarkersParams, GetStreamMarkersURL, []string{"user_id", "video_id"}),
		HttpMethod:   "GET",
		NeedOAuth2:   true,
		NeedClientID: false,
	}

	responseString, responseErr := oauth2.twitchRequest(twitchRequestData, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getStreamMarkersJSON GetStreamMarkersJSON
	json.Unmarshal([]byte(responseString), &getStreamMarkersJSON)

	return &getStreamMarkersJSON, nil
}

func (oauth2 *OAuth2) getStreamsMetadata(optionalParams map[string]interface{}) (*GetStreamsMetadataJSON, error) {
	getStreamsMetadataParams := make(map[string]interface{})

	// Required Params -> None
	// Optional Params
	optionalParamNames := []string{"after", "before", "first", "game_id", "language", "user_id", "user_login"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getStreamsMetadataParams[key] = value
	}

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getStreamsMetadataParams, GetStreamsMetadataURL, []string{}),
		HttpMethod:   "GET",
		NeedOAuth2:   false,
		NeedClientID: true,
	}

	responseString, responseErr := oauth2.twitchRequest(twitchRequestData, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getStreamsMetadataJSON GetStreamsMetadataJSON
	json.Unmarshal([]byte(responseString), &getStreamsMetadataJSON)

	return &getStreamsMetadataJSON, nil
}

func (oauth2 *OAuth2) getSubscriptions(broadcasterID string, optionalParams map[string]interface{}) (*GetSubscriptionsJSON, error) {
	getSubsParams := make(map[string]interface{})

	// Required Params
	getSubsParams["broadcaster_id"] = broadcasterID

	// Optional Params
	optionalParamNames := []string{"user_id"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getSubsParams[key] = value
	}

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getSubsParams, GetSubsURL, []string{"broadcaster_id"}),
		HttpMethod:   "GET",
		NeedOAuth2:   true,
		NeedClientID: false,
	}

	responseString, responseErr := oauth2.twitchRequest(twitchRequestData, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getSubsJSON GetSubscriptionsJSON
	json.Unmarshal([]byte(responseString), &getSubsJSON)

	return &getSubsJSON, nil
}

func (oauth2 *OAuth2) getAllStreamTags(optionalParams map[string]interface{}) (*GetAllStreamTagsJSON, error) {
	getAllStreamTagsParams := make(map[string]interface{})

	// Required Params -> None
	// Optional Params
	optionalParamNames := []string{"after", "first", "tag_id"}

	for key, value := range optionalParams {
		if !Contains(optionalParamNames, key) {
			panic(fmt.Sprintf("Unusable parameter -> %s", key))
		}

		getAllStreamTagsParams[key] = value
	}

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getAllStreamTagsParams, GetAllStreamTagsURL, []string{}),
		HttpMethod:   "GET",
		NeedOAuth2:   false,
		NeedClientID: true,
	}

	responseString, responseErr := oauth2.twitchRequest(twitchRequestData, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getAllStreamTagsJSON GetAllStreamTagsJSON
	json.Unmarshal([]byte(responseString), &getAllStreamTagsJSON)

	return &getAllStreamTagsJSON, nil
}

func (oauth2 *OAuth2) getStreamTags(broadcasterID interface{}) (*GetStreamTagsJSON, error) {
	getStreamTagsParams := make(map[string]interface{})

	// Required Params
	getStreamTagsParams["broadcaster_id"] = broadcasterID
	// Optional Params -> None

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getStreamTagsParams, GetStreamTagsURL, []string{}),
		HttpMethod:   "GET",
		NeedOAuth2:   false,
		NeedClientID: true,
	}

	responseString, responseErr := oauth2.twitchRequest(twitchRequestData, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	var getStreamTagsJSON GetStreamTagsJSON
	json.Unmarshal([]byte(responseString), &getStreamTagsJSON)

	return &getStreamTagsJSON, nil
}

func (oauth2 *OAuth2) replaceStreamTag(broadcasterID string, tagIDs []string) (string, error) {
	getStreamTagsParams := make(map[string]interface{})

	// Required Params
	getStreamTagsParams["broadcaster_id"] = broadcasterID

	twitchRequestData := &TwitchRequest{
		URL:          addParams(getStreamTagsParams, GetStreamTagsURL, []string{}),
		HttpMethod:   "GET",
		NeedOAuth2:   false,
		NeedClientID: true,
	}

	replaceStreamTagData := []byte(fmt.Sprintf(`{"tag_ids":"%s"`, tagIDs))
	fmt.Printf("\n%s\n", fmt.Sprintf(`{"tag_ids":"%s"`, tagIDs))
	responseString, responseErr := oauth2.twitchRequest(twitchRequestData, replaceStreamTagData)
	if responseErr != nil {
		return "", responseErr
	}

	return responseString, nil
}
