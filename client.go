package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	requestUrlTemplate = "http://api.openweathermap.org/data/2.5/find?q=%s&units=metric"
)

type Response struct {
	List []struct {
		Main struct {
			Temp float32 `json:"temp"`
		} `json:"main"`
	} `json:"list"`
}

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c Client) Request(city string) (*Weather, error) {
	requestUrl := getUrl(city)
	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var responseObject Response
	json.Unmarshal(body, &responseObject)
	return &Weather{Temperature: responseObject.List[0].Main.Temp}, nil
}

func getUrl(city string) string {
	return fmt.Sprintf(requestUrlTemplate, city)
}
