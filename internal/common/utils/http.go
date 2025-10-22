package utils

import (
	"fmt"
	"io"
	"net/http"
)

func ExecuteHTTPRequest(req *http.Request) []byte {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("failed to send request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("failed to read response")
	}

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("API error: status code %d, body: %s", resp.StatusCode, string(body))
		panic(errMsg)
	}

	return body
}

func SetHTTPHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
