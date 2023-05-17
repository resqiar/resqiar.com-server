package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"resdev-server/entities"
)

func ConvertToken(accessToken string) (*entities.GooglePayload, error) {
	resp, httpErr := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken))
	if httpErr != nil {
		return nil, httpErr
	}

	// clean up when this function returns (destroyed)
	defer resp.Body.Close()

	respBody, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	// Bind JSON into struct
	var data entities.GooglePayload
	err := json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
