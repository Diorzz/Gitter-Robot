package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpPost(uri, header, param string) (string, error) {

	resp, err := http.Post(uri, "application/json", strings.NewReader(param))
	if err != nil {
		fmt.Printf("Post error, %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func httpGet(uri string) (string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Printf("Post error, %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
