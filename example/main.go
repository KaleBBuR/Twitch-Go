package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	twitchgo ".."
)

type Private struct {
	ClientID     string `json:"client ID"`
	ClientSecret string `json:"client Secret"`
	Scope        string `json:"scope"`
}

func main() {
	private := Private{}
	raw, err := ioutil.ReadFile("private.json")
	if err != nil {
		return
	}

	json.Unmarshal(raw, &private)

	clientID := private.ClientID
	clientSecret := private.ClientSecret
	scope := private.Scope
	oauth, err := twitchgo.OAuthLoginSession(clientID, clientSecret, scope)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	getStreamsMap := make(map[string]interface{})
	getStreamsMap["first"] = 100

	twitchgorequest, err := oauth.getStreams(getStreamsMap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, value := range twitchgorequest.Data {
		fmt.Printf("\n\nUsername: %s\nTitle: %s\nViewer Count: %d\n\n", value.Username, value.Title, value.ViewerCount)
	}
}
