package twitchgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func OAuthLoginSession(clientID string, clientSecret string, scope string) (*OAuth2, error) {
	Oauth := &OAuth2{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	formData := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"scope":         {scope},
	}

	request, postErr := http.NewRequest("POST", OAUTH2URL, bytes.NewBufferString(formData.Encode()))
	if postErr != nil {
		return nil, postErr
	}

	client := &http.Client{}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return nil, responseErr
	}

	if response.StatusCode != http.StatusOK {
		body, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			return nil, readErr
		}

		bodyString := string(body)
		var oauthjson map[string]interface{}
		json.Unmarshal([]byte(bodyString), &oauthjson)
		return nil, errors.New(fmt.Sprintf("Did not a return a good status code.\nCode: %d\nError: %s", response.StatusCode, oauthjson["error"]))
	}

	defer response.Body.Close()

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}

	bodyString := string(body)
	json.Unmarshal([]byte(bodyString), Oauth)

	go Oauth.expireTimeCountdown()

	return Oauth, nil
}

func (oauth *OAuth2) twitchRequest(twitchRequestData *TwitchRequest, requestData []byte) (string, error) {
	request, postErr := http.NewRequest(twitchRequestData.HttpMethod, twitchRequestData.URL, bytes.NewBuffer(requestData))
	if postErr != nil {
		return "", postErr
	}

	if twitchRequestData.NeedClientID {
		request.Header.Set("Client-ID", oauth.ClientID)
	}

	if twitchRequestData.NeedOAuth2 {
		request.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	}

	client := &http.Client{}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return "", responseErr
	}

	if response.StatusCode != http.StatusOK {
		body, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			return "", readErr
		}

		bodyString := string(body)
		var oauthjson map[string]interface{}
		json.Unmarshal([]byte(bodyString), &oauthjson)
		return "", errors.New(fmt.Sprintf("Did not a return a good status code.\nCode: %d\nError: %s", oauthjson["code"], oauthjson["error"]))
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return "", readErr
	}

	bodyString := string(body)
	return bodyString, nil
}

func (oauth *OAuth2) expireTimeCountdown() {
	expireTime := oauth.ExpireTime
	for {
		time.Sleep(1 * time.Second)
		expireTime -= 1
		if expireTime == 0 {
			oAuth2, err := OAuthLoginSession(
				oauth.ClientID,
				oauth.ClientSecret,
				oauth.Scope,
			)

			if err != nil {
				log.Fatal(errors.New("Couldn't get new access token."))
			}

			oauth.AccessToken = oAuth2.AccessToken
			oauth.expireTimeCountdown()
		}
	}
}
