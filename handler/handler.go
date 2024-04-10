package authorizationApiController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/farismnrr/go-auth-api-consume/helper"
	"github.com/farismnrr/go-auth-api-consume/model"
)

func GetDataFromAPI() (*model.ResponseData, error) {
	_, authToken := helper.ReadJsonFile()

	req, err := http.NewRequest("GET", "https://authorization-api-dot-farismnrr-gclouds.as.r.appspot.com/copyright", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	var response model.ResponseData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return &response, nil
}

func CopyrightHandler() bool {
	response, err := GetDataFromAPI()
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return false
	}

	username, _ := helper.ReadJsonFile()

	hashed := helper.GenerateHash(username)

	for _, userData := range response.Data {
		if hashed == userData.CopyrightAuthorization {
			fmt.Println("Copyright Authorized by", userData.Username)
			return true
		}
	}

	return false // If user not found or authorization failed
}
