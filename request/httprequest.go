package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpRequest struct{}

func (hr *HttpRequest) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, fmt.Errorf("Http Get error, status code: %d", resp.StatusCode)

}

func (hr *HttpRequest) Post(uri, header, param string) ([]byte, error) {

	resp, err := http.Post(uri, "application/json", strings.NewReader(param))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
