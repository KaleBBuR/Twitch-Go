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

func OAuthLoginSession(clientID string, clientSecret string, scope string) (*Session, error) {
	sess := &Session{
		clientID:     clientID,
		clientSecret: clientSecret,
		scope:        scope,
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

		var oauthjson map[string]interface{}
		json.Unmarshal(body, &oauthjson)
		return nil, errors.New(fmt.Sprintf("Did not a return a good status code.\nCode: %d\nError: %s", response.StatusCode, oauthjson["error"]))
	}

	defer response.Body.Close()

	body, readErr := ioutil.ReadAll(response.Body)
	bodyString := string(body)
	if readErr != nil {
		return nil, readErr
	}

	var oauthjson map[string]interface{}
	json.Unmarshal([]byte(bodyString), &oauthjson)
	accessToken, atOK := oauthjson["access_token"].(string)
	expireTime, etOK := oauthjson["expires_in"].(float64)
	scopes, sOK := oauthjson["scope"].([]string)

	if atOK && etOK {
		sess.accessToken = accessToken
		sess.expireTime = expireTime
	} else {
		return nil, errors.New("Can't get Expire Time or Access Token.")
	}

	if sOK {
		sess.scopes = scopes
	}

	go sess.expireTimeCountdown()

	return sess, nil
}

func (sess *Session) TwitchRequest(twitchRequestData *TwitchRequest, requestData []byte) (string, error) {
	request, postErr := http.NewRequest(twitchRequestData.HttpMethod, twitchRequestData.URL, bytes.NewBuffer(requestData))
	if postErr != nil {
		return "", postErr
	}

	if twitchRequestData.NeedClientID {
		request.Header.Set("Client-ID", sess.clientID)
	}

	if twitchRequestData.NeedOAuth2 {
		request.Header.Set("Authorization", "Bearer "+sess.accessToken)
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

func (sess1 *Session) expireTimeCountdown() {
	expireTime := sess1.expireTime
	for {
		time.Sleep(1 * time.Second)
		expireTime -= 1
		if expireTime == 0 {
			sess2, err := OAuthLoginSession(
				sess1.clientID,
				sess1.clientSecret,
				sess1.scope,
			)

			if err != nil {
				log.Fatal(errors.New("Couldn't get new access token."))
			}

			sess1.accessToken = sess2.accessToken
			sess1.expireTimeCountdown()
		}
	}
}
