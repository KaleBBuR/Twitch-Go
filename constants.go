package twitchgo

import (
	"fmt"
	"strings"
)

var (
	BASEURL                    = "https://api.twitch.tv/helix"
	OAUTH2URL                  = "https://id.twitch.tv/oauth2/token"
	CreateGetClipURL           = BASEURL + "/clips"
	GetTopGamesURL             = BASEURL + "/games/top"
	GetGamesURL                = BASEURL + "/games"
	CheckAutoModStatusURL      = BASEURL + "/moderation/enforments/status"
	GetBannedEventsURL         = BASEURL + "/moderation/banned/events"
	GetBannedUsersURL          = BASEURL + "/moderation/banned"
	GetModsURL                 = BASEURL + "/moderation/moderators"
	GetModEventsURL            = BASEURL + "/moderation/moderators/events"
	CreateStreamMarkerURL      = BASEURL + "/streams/markers"
	GetStreamsURL              = BASEURL + "/streams"
	GetStreamMarkersURL        = BASEURL + "/streams/markers"
	GetStreamsMetadataURL      = BASEURL + "/streams/metadata"
	GetSubsURL                 = BASEURL + "/subscriptions"
	GetAllStreamTagsURL        = BASEURL + "/tags/streams"
	GetStreamTagsURL           = BASEURL + "/streams/tags"
	GetUsersURL                = BASEURL + "/users"
	GetUsersFollowsURL         = BASEURL + "/users/follows"
	GetUserExtensionsURL       = BASEURL + "/users/extensions/list"
	GetUserActiveExtensionsURL = BASEURL + "/users/extensions"
	GetVideosURL               = BASEURL + "/videos"
)

type CreateClipJSON struct {
	Data []struct {
		ID      string `json:"id"`
		EditURL string `json:"edit_url"`
	} `json:"data"`
}

type GetClipJSON struct {
	Data []struct {
		ID              string `json:"id"`
		URL             string `json:"url"`
		EmbedURL        string `json:"embed_url"`
		BroadcasterID   string `json:"broadcaster_id"`
		BroadcasterName string `json:"broadcaster_name"`
		CreatorID       string `json:"creator_id"`
		CreatorName     string `json:"creator_name"`
		VideoID         string `json:"video_id"`
		GameID          string `json:"game_id"`
		Language        string `json:"language"`
		Title           string `json:"title"`
		ViewCount       int    `json:"view_count"`
		CreatedAt       string `json:"created_at"`
		ThumbnailURL    string `json:"thumbnail_url"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type GetTopGamesJSON struct {
	Data []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		BoxArtURL string `json:"box_art_url"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type GetGamesJSON struct {
	Data []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		BoxArtURL string `json:"box_art_url"`
	} `json:"data"`
}

type AutoModStatusBody struct {
	MsgID   string `json:"msg_id"`
	MsgText string `json:"msg_text"`
	UserID  string `json:"user_id"`
}

type AutoModStatusFullBody struct {
	Data []AutoModStatusBody `json:"data"`
}

type PostAutoModStatusJSON struct {
	Data []struct {
		MessageID   string `json:"msg_id"`
		IsPermitted bool   `json:"is_permitted"`
	} `json:"data"`
}

type GetBannedEventsJSON struct {
	Data []struct {
		ID             string `json:"id"`
		EventType      string `json:"event_type"`
		EventTimestamp string `json:"event_timestamp"`
		Version        string `json:"version"`
		EventData      struct {
			BroadcasterID   string `json:"broadcaster_id"`
			BroadcasterName string `json:"broadcaster_name"`
			UserID          string `json:"user_id"`
			Username        string `json:"user_name"`
			ExpireTime      string `json:"expires_at"`
		} `json:"event_data"`
	} `json:"data"`
}

type GetBannedUsersJSON struct {
	Data []struct {
		UserID      string `json:"user_id"`
		Username    string `json:"user_name"`
		ExpiresTime string `json:"expires_at"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type GetModsJSON struct {
	Data []struct {
		UserID   string `json:"user_id"`
		Username string `json:"user_name"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type GetModEventsJSON struct {
	Data []struct {
		ID             string `json:"id"`
		EventType      string `json:"event_type"`
		EventTimestamp string `json:"event_timestamp"`
		Version        string `json:"version"`
		EventData      struct {
			BroadcasterID   string `json:"broadcaster_id"`
			BroadcasterName string `json:"broadcaster_name"`
			UserID          string `json:"user_id"`
			Username        string `json:"user_name"`
		} `json:"event_data"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type CreateStreamMarkerJSON struct {
	Data []struct {
		ID              int    `json:"id"`
		CreatedAt       string `json:"created_at"`
		Description     string `json:"description"`
		PositionSeconds int    `json:"position_seconds"`
	} `json:"data"`
}

type GetStreamsJSON struct {
	Data []struct {
		ID          string   `json:"id"`
		UserID      string   `json:"user_id"`
		Username    string   `json:"user_name"`
		GameID      string   `json:"game_id"`
		Type        string   `json:"type"`
		Title       string   `json:"title"`
		ViewerCount int      `json:"viewer_count"`
		StartedAt   string   `json:"started_at"`
		Language    string   `json:"language"`
		Thumbnail   string   `json:"thumbnail_url"`
		TagIDs      []string `json:"tag_ids"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type GetStreamMarkersJSON struct {
	Data []struct {
		UserID   string `json:"user_id"`
		Username string `json:"user_name"`
		Videos   []struct {
			VideoID string `json:"video_id"`
			Markers []struct {
				ID              string `json:"id"`
				CreatedAt       string `json:"created_at"`
				Description     string `json:"description"`
				PositionSeconds int    `json:"position_seconds"`
				URL             string `json:"URL"`
			} `json:"markers"`
		} `json:"videos"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type GetStreamsMetadataJSON struct {
	Data []struct {
		UserID    string `json:"user_id"`
		Username  string `json:"user_name"`
		GameID    string `json:"game_id"`
		Overwatch struct {
			Broadcaster struct {
				Hero struct {
					Role    string `json:"role"`
					Name    string `json:"name"`
					Ability string `json:"ability"`
				} `json:"hero"`
			} `json:"broadcaster"`
		} `json:"overwatch"`
		Hearthstrong struct {
			Broadcaster struct {
				Hero struct {
					Type  string `json:"type"`
					Class string `json:"class"`
					Name  string `json:"name"`
				} `json:"hero"`
			} `json:"broadcaster"`
			Opponent struct {
				Hero struct {
					Type  string `json:"type"`
					Class string `json:"class"`
					Name  string `json:"name"`
				} `json:"hero"`
			} `json:"opponent"`
		} `json:"hearthstone"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type GetSubscriptionsJSON struct {
	Data []struct {
		BroadcasterID   string `json:"broadcaster_id"`
		BroadcasterName string `json:"broadcaster_name"`
		IsGift          bool   `json:"is_gift"`
		Tier            string `json:"tier"`
		PlanName        string `json:"plan_name"`
		UserID          string `json:"user_id"`
		Username        string `json:"user_name"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type GetAllStreamTagsJSON struct {
	Data []struct {
		TagID                    string                 `json:"tag_id"`
		IsAuto                   bool                   `json:"is_auto"`
		LocalizationNames        map[string]interface{} `json:"localization_names"`
		LocalizationDescriptions map[string]interface{} `json:"localization_descriptions"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type GetStreamTagsJSON struct {
	Data []struct {
		TagID                    string                 `json:"tag_id"`
		IsAuto                   bool                   `json:"is_auto"`
		LocalizationNames        map[string]interface{} `json:"localization_names"`
		LocalizationDescriptions map[string]interface{} `json:"localization_descriptions"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type GetUsersJSON struct {
	Data []struct {
		ID              string `json:"id"`
		Login           string `json:"login"`
		DisplayName     string `json:"display_name"`
		Type            string `json:"type"`
		BroadcasterType string `json:"broadcaster_type"`
		Description     string `json:"description"`
		ProfileImageURL string `json:"profile_image_url"`
		OfflineImageURL string `json:"offline_image_url"`
		Viewcount       int    `json:"view_count"`
		Email           string `json:"email"`
	} `json:"data"`
}

type GetUsersFollowsJSON struct {
	Total int `json:"total"`
	Data  []struct {
		FromID     string `json:"from_id"`
		FromName   string `json:"from_name"`
		ToID       string `json:"to_id"`
		ToName     string `json:"to_name"`
		FollowedAt string `json:"followed_at"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type GetUserExtensionsJSON struct {
	Data []struct {
		ID          string   `json:"id"`
		Version     string   `json:"verison"`
		Name        string   `json:"name"`
		CanActivate bool     `json:"can_activate"`
		Type        []string `json:"type"`
	} `json:"data"`
}

type GetVideosJSON struct {
	Data []struct {
		ID           string `json:"id"`
		UserID       string `json:"user_id"`
		Username     string `json:"user_name"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		CreatedAt    string `json:"created_at"`
		PublishedAt  string `json:"published_at"`
		URL          string `json:"url"`
		ThumbnailURL string `json:"thumbnail_url"`
		Viewable     string `json:"viewable"`
		ViewCount    int    `json:"view_count"`
		Language     string `json:"language"`
		Type         string `json:"type"`
		Duration     string `json:"duration"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// type GetUserActiveExtensions struct {
// }

type Session struct {
	clientID     string
	clientSecret string
	scope        string
	accessToken  string
	expireTime   float64
	scopes       []string
}

type TwitchRequest struct {
	URL          string
	HttpMethod   string
	NeedOAuth2   bool
	NeedClientID bool
}

func addParams(params map[string]interface{}, url string, requiredParams []string) string {
	url += "?"

	for _, requiredParam := range requiredParams {
		if _, ok := params[requiredParam]; !ok {
			panic(fmt.Sprintf("Need Parameter -> %s", requiredParam))
		}
	}

	for key, value := range params {
		switch value.(type) {
		case string:
			value = value.(string)
			if value != "" {
				url = url + fmt.Sprintf("%s=%s&", key, value)
			} else {
				if Contains(requiredParams, key) {
					panic(fmt.Sprintf("You must have a value for this key -> %s", key))
				}
			}
		case int:
			url = url + fmt.Sprintf("%s=%d&", key, value.(int))
		case bool:
			url = url + fmt.Sprintf("%s=%t&", key, value.(bool))
		default:
			panic("Parameter value must be String, Integer, or Boolean")
		}
	}

	url = strings.TrimSuffix(url, "&")
	return url
}

func Contains(a []string, key string) bool {
	for _, n := range a {
		if key == n {
			return true
		}
	}

	return false
}

func (sess *Session) GetResponse(URL string, method string, oauth bool, id bool, body []byte) ([]byte, error) {
	twitchRequestData := &TwitchRequest{
		URL:          URL,
		HttpMethod:   method,
		NeedOAuth2:   oauth,
		NeedClientID: id,
	}

	responseString, responseErr := sess.TwitchRequest(twitchRequestData, body)
	if responseErr != nil {
		return nil, responseErr
	}

	return []byte(responseString), nil
}

func PanicErr(err error) {
	panic(fmt.Sprintf("ERROR: %s", err.Error()))
}
