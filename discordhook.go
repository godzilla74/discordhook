package discordhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func SendMessage(url string, message Message) error {
	payload := new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(message)
	if err != nil {
		return err
	}

	response, err := http.Post(url, "application/json", payload)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
			}
		}(response.Body)

		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf(string(responseBody))
	}

	return nil
}
