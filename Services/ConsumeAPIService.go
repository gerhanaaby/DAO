package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ConsumeAPIService(url []string, JsonData []byte) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url[0], bytes.NewBuffer(JsonData))
	if err != nil {
		return nil, err
	}

	fmt.Print(url[1])
	req.Header.Set(`Content-Type`, `application/json`)
	req.Header.Set(`X-Gateway-APIKey`, url[1])
	req.Header.Set(`Authorization`, url[2])

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, _ := ioutil.ReadAll(res.Body)

	return respBody, nil
}
