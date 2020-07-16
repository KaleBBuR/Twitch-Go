package twitchgo

import (
	"encoding/json"
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

	data, dataErr := sess.GetResponse(addParams(createStreamMarkerParams, CreateStreamMarkerURL, []string{"user_id"}), "POST", true, false, nil)
	if dataErr != nil {
		PanicErr(dataErr)
	}

	var createStreamMarkerJSON CreateStreamMarkerJSON
	json.Unmarshal(data, &createStreamMarkerJSON)

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

	data, dataErr := sess.GetResponse(addParams(getStreamsParams, GetStreamsURL, []string{}), "GET", true, true, nil)
	if dataErr != nil {
		PanicErr(dataErr)
	}

	var getStreamsJSON GetStreamsJSON
	json.Unmarshal(data, &getStreamsJSON)

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

	data, dataErr := sess.GetResponse(addParams(getStreamMarkersParams, GetStreamMarkersURL, []string{"user_id", "video_id"}), "GET", true, false, nil)
	if dataErr != nil {
		PanicErr(dataErr)
	}

	var getStreamMarkersJSON GetStreamMarkersJSON
	json.Unmarshal(data, &getStreamMarkersJSON)

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

	data, dataErr := sess.GetResponse(addParams(getStreamsMetadataParams, GetStreamsMetadataURL, []string{}), "GET", false, true, nil)
	if dataErr != nil {
		PanicErr(dataErr)
	}

	var getStreamsMetadataJSON GetStreamsMetadataJSON
	json.Unmarshal(data, &getStreamsMetadataJSON)

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

	data, dataErr := sess.GetResponse(addParams(getSubsParams, GetSubsURL, []string{"broadcaster_id"}), "GET", true, false, nil)
	if dataErr != nil {
		PanicErr(dataErr)
	}

	var getSubsJSON GetSubscriptionsJSON
	json.Unmarshal(data, &getSubsJSON)

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

	data, dataErr := sess.GetResponse(addParams(getAllStreamTagsParams, GetAllStreamTagsURL, []string{}), "GET", false, true, nil)
	if dataErr != nil {
		PanicErr(dataErr)
	}

	var getAllStreamTagsJSON GetAllStreamTagsJSON
	json.Unmarshal(data, &getAllStreamTagsJSON)

	return &getAllStreamTagsJSON, nil
}

func (sess *Session) GetStreamTags(broadcasterID interface{}) (*GetStreamTagsJSON, error) {
	getStreamTagsParams := make(map[string]interface{})

	// Required Params
	getStreamTagsParams["broadcaster_id"] = broadcasterID
	// Optional Params -> None

	data, dataErr := sess.GetResponse(addParams(getStreamTagsParams, GetStreamTagsURL, []string{}), "GET", false, true, nil)
	if dataErr != nil {
		PanicErr(dataErr)
	}

	var getStreamTagsJSON GetStreamTagsJSON
	json.Unmarshal(data, &getStreamTagsJSON)

	return &getStreamTagsJSON, nil
}

func (sess *Session) ReplaceStreamTag(broadcasterID string, tagIDs []string) (string, error) {
	getStreamTagsParams := make(map[string]interface{})

	// Required Params
	getStreamTagsParams["broadcaster_id"] = broadcasterID

	data, dataErr := sess.GetResponse(addParams(getStreamTagsParams, GetStreamTagsURL, []string{}), "GET", false, true, []byte(fmt.Sprintf(`{"tag_ids":"%s"`, tagIDs)))
	if dataErr != nil {
		PanicErr(dataErr)
	}

	return string(data), nil
}
