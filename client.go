package main

import (
	b64 "encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type ScheduleClientConfig struct {
	User     string
	Password string
	Date     string
	BaseURL  string
}

func RequestXML(config ScheduleClientConfig) ([]byte, error) {
	requestURL := config.BaseURL + config.Date + ".xml"
	httpClient := http.Client{}
	request, _ := http.NewRequest("GET", requestURL, nil)
	request.Header.Add("authorization", encodeAsBasicAuth(config.User, config.Password))
	log.Println("executing request to " + requestURL)
	resp, respErr := httpClient.Do(request)
	if respErr != nil {
		log.Println(respErr.Error())
		return nil, errors.New("HTTP Error: " + respErr.Error())
	}
	data, parseBodyErr := ioutil.ReadAll(resp.Body)
	if parseBodyErr == nil {
		return data, nil
	} else {
		return nil, errors.New("Parse Response Body Error: " + parseBodyErr.Error())
	}
}

func encodeAsBasicAuth(user string, password string) string {
	return "Basic " + b64.URLEncoding.EncodeToString([]byte(user+":"+password))
}
