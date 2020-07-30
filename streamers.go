package twitchgo

import (
	"fmt"
)

func (sess *Session) CreateStreamMarker(userID string, optionalParams map[string]interface{}) (*CreateStreamMarkerJSON, error) {
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

	var createStreamMarkerJSON CreateStreamMarkerJSON
	dataErr := sess.GetResponse(addParams(createStreamMarkerParams, CreateStreamMarkerURL, []string{"user_id"}), "POST", true, false, nil, &createStreamMarkerJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &createStreamMarkerJSON, nil
}

func (sess *Session) GetStreams(optionalParams map[string]interface{}) (*GetStreamsJSON, error) {
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

	var getStreamsJSON GetStreamsJSON
	dataErr := sess.GetResponse(addParams(getStreamsParams, GetStreamsURL, []string{}), "GET", true, true, nil, &getStreamsJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getStreamsJSON, nil
}

func (sess *Session) GetStreamMarkers(userID string, videoID string, optionalParams map[string]interface{}) (*GetStreamMarkersJSON, error) {
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

	var getStreamMarkersJSON GetStreamMarkersJSON
	dataErr := sess.GetResponse(addParams(getStreamMarkersParams, GetStreamMarkersURL, []string{"user_id", "video_id"}), "GET", true, false, nil, &getStreamMarkersJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getStreamMarkersJSON, nil
}

func (sess *Session) GetStreamsMetadata(optionalParams map[string]interface{}) (*GetStreamsMetadataJSON, error) {
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

	var getStreamsMetadataJSON GetStreamsMetadataJSON
	dataErr := sess.GetResponse(addParams(getStreamsMetadataParams, GetStreamsMetadataURL, []string{}), "GET", false, true, nil, &getStreamsMetadataJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getStreamsMetadataJSON, nil
}

func (sess *Session) GetSubscriptions(broadcasterID string, optionalParams map[string]interface{}) (*GetSubscriptionsJSON, error) {
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

	var getSubsJSON GetSubscriptionsJSON
	dataErr := sess.GetResponse(addParams(getSubsParams, GetSubsURL, []string{"broadcaster_id"}), "GET", true, false, nil, &getSubsJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getSubsJSON, nil
}

func (sess *Session) GetAllStreamTags(optionalParams map[string]interface{}) (*GetAllStreamTagsJSON, error) {
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

	var getAllStreamTagsJSON GetAllStreamTagsJSON
	dataErr := sess.GetResponse(addParams(getAllStreamTagsParams, GetAllStreamTagsURL, []string{}), "GET", false, true, nil, &getAllStreamTagsJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getAllStreamTagsJSON, nil
}

func (sess *Session) GetStreamTags(broadcasterID interface{}) (*GetStreamTagsJSON, error) {
	getStreamTagsParams := make(map[string]interface{})

	// Required Params
	getStreamTagsParams["broadcaster_id"] = broadcasterID
	// Optional Params -> None

	var getStreamTagsJSON GetStreamTagsJSON
	dataErr := sess.GetResponse(addParams(getStreamTagsParams, GetStreamTagsURL, []string{}), "GET", false, true, nil, &getStreamTagsJSON)
	if dataErr != nil {
		return nil, dataErr
	}

	return &getStreamTagsJSON, nil
}

func (sess *Session) ReplaceStreamTag(broadcasterID string, tagIDs []string) (string, error) {
	getStreamTagsParams := make(map[string]interface{})

	// Required Params
	getStreamTagsParams["broadcaster_id"] = broadcasterID

	var data string
	dataErr := sess.GetResponse(addParams(getStreamTagsParams, GetStreamTagsURL, []string{}), "GET", false, true, []byte(fmt.Sprintf(`{"tag_ids":"%s"`, tagIDs)), &data)
	if dataErr != nil {
		return "", dataErr
	}

	return data, nil
}
