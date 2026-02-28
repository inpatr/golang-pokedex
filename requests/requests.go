package requests

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io"
)

type locationAreaResponse struct {
	count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []locationAreaResult `json:"results"`
}

type locationAreaResult struct {
	Name string `json:"name"`
	url string `json:"url"`
}

func GetLocationAreas(url string) (locationAreaResponse, error) {
	var reqBody locationAreaResponse
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error on getting location areas: %v", err)
		return reqBody, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error on reading stream: %v", err)
		return reqBody, err
	}

	err = json.Unmarshal(data, &reqBody)
	if err != nil {
		fmt.Printf("Error on Unmarshalling: %v", err)
		return reqBody, err
	}
	
	return reqBody, nil
}
